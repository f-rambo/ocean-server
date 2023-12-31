package data

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	wfv1 "github.com/argoproj/argo-workflows/v3/pkg/apis/workflow/v1alpha1"
	"github.com/f-rambo/ocean/internal/biz"
	"github.com/f-rambo/ocean/pkg/argoworkflows"
	operatoroceaniov1alpha1 "github.com/f-rambo/operatorapp/api/v1alpha1"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/spf13/cast"
	"gorm.io/gorm"
	k8serr "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type servicesRepo struct {
	data *Data
	log  *log.Helper
}

func NewServicesRepo(data *Data, logger log.Logger) biz.ServicesRepo {
	return &servicesRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (s *servicesRepo) k8s() error {
	if s.data.k8sClient != nil {
		return nil
	}
	return s.data.newKubernetes()
}

func (s *servicesRepo) Save(ctx context.Context, svc *biz.Service) error {
	if svc.ID == 0 {
		var Count int64
		err := s.data.db.Model(&biz.Service{}).Where("name = ?", svc.Name).Count(&Count).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			return err
		}
		if Count > 0 {
			return errors.New("service name already exists")
		}
	}
	return s.data.db.Omit("created_at").Save(svc).Error
}

func (s *servicesRepo) Get(ctx context.Context, id int) (*biz.Service, error) {
	service := &biz.Service{}
	err := s.data.db.Where("id = ?", id).First(service).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	cis := make([]*biz.CI, 0)
	err = s.data.db.Where("service_id = ?", id).Find(&cis).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	service.CIItems = cis
	return service, nil
}

func (s *servicesRepo) GetServices(ctx context.Context) ([]*biz.Service, error) {
	services := make([]*biz.Service, 0)
	err := s.data.db.Find(&services).Error
	if err != nil {
		return nil, err
	}
	return services, nil
}

func (s *servicesRepo) Delete(ctx context.Context, svc *biz.Service) error {
	// gorm 事务
	var err error
	tx := s.data.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
		if err != nil {
			tx.Rollback()
		}
	}()
	err = tx.Where("id = ?", svc.ID).Delete(&biz.CI{}).Error
	if err != nil {
		return err
	}
	err = tx.Where("service_id = ?", svc.ID).Delete(&biz.CI{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *servicesRepo) SaveCI(ctx context.Context, ci *biz.CI) error {
	if ci.ID == 0 {
		// 先创建获取主键ID
		err := s.data.db.Save(ci).Error
		if err != nil {
			return err
		}
	}
	// kubernetes clientset
	if err := s.k8s(); err != nil {
		return err
	}
	svc, err := s.Get(ctx, ci.ServiceID)
	if err != nil {
		return err
	}
	wf, err := argoworkflows.UnmarshalWorkflow(svc.Workflow, true)
	if err != nil {
		return err
	}
	s.wfAssign(ctx, svc, ci, &wf)
	wf.Labels = map[string]string{
		"service":    svc.Name,
		"service_id": cast.ToString(svc.ID),
		"ci_id":      cast.ToString(ci.ID),
	}
	resWf, err := s.data.workflowClient.Workflows(wf.Namespace).Create(ctx, &wf)
	if err != nil {
		return err
	}
	ci.SetWorkflowName(resWf.Name)
	err = s.data.db.Omit("created_at").Save(ci).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *servicesRepo) GetCI(ctx context.Context, id int) (*biz.CI, error) {
	ci := &biz.CI{}
	err := s.data.db.Where("id = ?", id).First(ci).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return ci, nil
}

func (s *servicesRepo) GetCIs(ctx context.Context, serviceID int) ([]*biz.CI, error) {
	cis := make([]*biz.CI, 0)
	err := s.data.db.Where("service_id = ?", serviceID).Find(&cis).Error
	if err != nil {
		return nil, err
	}
	return cis, nil
}

func (s *servicesRepo) DeleteCI(ctx context.Context, ci *biz.CI) error {
	if ci == nil || ci.ID == 0 {
		return errors.New("ci is nil")
	}
	err := s.k8s()
	if err != nil {
		return err
	}
	service, err := s.Get(ctx, ci.ServiceID)
	if err != nil {
		return err
	}
	err = s.data.workflowClient.Workflows(service.NameSpace).Delete(ctx, ci.WorkflowName)
	if err != nil && !k8serr.IsNotFound(err) {
		return err
	}
	err = s.data.db.Where("id = ?", ci.ID).Delete(&biz.CI{}).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	return nil
}

func (s *servicesRepo) Deploy(ctx context.Context, svc *biz.Service, ci *biz.CI) error {
	app := &operatoroceaniov1alpha1.App{}
	app.Name = svc.Name
	app.Namespace = svc.NameSpace
	app.Labels = map[string]string{"app": app.Name}
	app.Spec.Service = operatoroceaniov1alpha1.Service{
		Enable:        true,
		EnableIngress: true,
		EnableService: true,
		Replicas:      svc.Replicas,
		Image:         fmt.Sprintf("%s/%s:%s", svc.Registry, svc.Name, ci.Version),
		CPU:           svc.CPU,
		LimitCPU:      svc.LimitCpu,
		Memory:        svc.Memory,
		LimitMemory:   svc.LimitMemory,
		Config:        svc.Config,
		Secret:        svc.Secret,
	}
	for _, port := range svc.Ports {
		app.Spec.Service.Ports = append(app.Spec.Service.Ports, operatoroceaniov1alpha1.Port{
			IngressPath:   port.IngressPath,
			ContainerPort: port.ContainerPort,
		})
	}
	err := s.k8s()
	if err != nil {
		return err
	}
	resAppData, err := s.data.operatorappClient.Apps(app.Namespace).Get(ctx, app.Name, metav1.GetOptions{})
	if err != nil && !k8serr.IsNotFound(err) {
		return err
	}
	if resAppData == nil || k8serr.IsNotFound(err) {
		_, err = s.data.operatorappClient.Apps(app.Namespace).Create(ctx, app)
		if err != nil {
			return err
		}
		return nil
	}
	app.ResourceVersion = resAppData.ResourceVersion
	_, err = s.data.operatorappClient.Apps(app.Namespace).Update(ctx, app)
	if err != nil {
		return err
	}
	return nil
}

func (s *servicesRepo) UnDeploy(ctx context.Context, svc *biz.Service) error {
	err := s.k8s()
	if err != nil {
		return err
	}
	err = s.data.operatorappClient.Apps(svc.NameSpace).Delete(ctx, svc.Name)
	if err != nil && !k8serr.IsNotFound(err) {
		return err
	}
	return nil
}

func (s *servicesRepo) GetOceanService(ctx context.Context) (*biz.Service, error) {
	wf, err := argoworkflows.GetDefaultWorkflowStr()
	if err != nil {
		return nil, err
	}
	service := &biz.Service{
		Name:         "ocean",
		Repo:         "https://github.com/f-rambo/ocean.git",
		Registry:     "https://docker.io/frambos",
		RegistryUser: "",
		RegistryPwd:  "",
		Workflow:     wf,
		NameSpace:    "default",
	}
	ci := &biz.CI{
		Version:     "0.0.1",
		Branch:      "main",
		Tag:         "",
		Args:        "", // json string
		Description: "example",
		ServiceID:   service.ID,
	}
	service.CIItems = append(service.CIItems, ci)
	return service, nil
}

func (s *servicesRepo) wfAssign(ctx context.Context, svc *biz.Service, ci *biz.CI, wf *wfv1.Workflow) {
	args := make(map[string]string)
	if ci != nil && ci.Args != "" {
		json.Unmarshal([]byte(ci.Args), &args)
	}
	wf.Namespace = svc.NameSpace
	for i, param := range wf.Spec.Arguments.Parameters {
		switch param.Name {
		case "name":
			wf.Spec.Arguments.Parameters[i].Value = wfv1.AnyStringPtr(svc.Name)
		case "repo":
			wf.Spec.Arguments.Parameters[i].Value = wfv1.AnyStringPtr(svc.Repo)
		case "registry":
			wf.Spec.Arguments.Parameters[i].Value = wfv1.AnyStringPtr(svc.Registry)
		case "registry_user":
			wf.Spec.Arguments.Parameters[i].Value = wfv1.AnyStringPtr(svc.RegistryUser)
		case "registry_pwd":
			wf.Spec.Arguments.Parameters[i].Value = wfv1.AnyStringPtr(svc.RegistryPwd)
		case "version":
			wf.Spec.Arguments.Parameters[i].Value = wfv1.AnyStringPtr(ci.Version)
		case "branch":
			wf.Spec.Arguments.Parameters[i].Value = wfv1.AnyStringPtr(ci.Branch)
		case "tag":
			wf.Spec.Arguments.Parameters[i].Value = wfv1.AnyStringPtr(ci.Tag)
		default:
			if val, ok := args[param.Name]; ok {
				wf.Spec.Arguments.Parameters[i].Value = wfv1.AnyStringPtr(val)
			}
		}
	}
}
