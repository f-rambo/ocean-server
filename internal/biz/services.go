package biz

import (
	"context"
	"errors"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type Service struct {
	ID           int    `json:"id" gorm:"column:id;primaryKey;AUTO_INCREMENT"`
	Name         string `json:"name,omitempty" gorm:"column:name; default:''; NOT NULL"`
	Repo         string `json:"repo,omitempty" gorm:"column:repo; default:''; NOT NULL"`
	Registry     string `json:"registry" gorm:"column:registry; default:''; NOT NULL"`
	RegistryUser string `json:"registry_user" gorm:"column:registry_user; default:''; NOT NULL"`
	RegistryPwd  string `json:"registry_pwd" gorm:"column:registry_pwd; default:''; NOT NULL"`
	Workflow     string `json:"workflow" gorm:"column:workflow; default:''; NOT NULL"`
	CIItems      []*CI  `json:"ci_items,omitempty" gorm:"-"`
	gorm.Model
}

type CI struct {
	ID           int               `json:"id" gorm:"column:id;primaryKey;AUTO_INCREMENT"`
	NameSpace    string            `json:"namespace,omitempty" gorm:"column:namespace; default:''; NOT NULL"`
	Version      string            `json:"version,omitempty" gorm:"column:version; default:''; NOT NULL"`
	Branch       string            `json:"branch,omitempty" gorm:"column:branch; default:''; NOT NULL"`
	Tag          string            `json:"tag,omitempty" gorm:"column:tag; default:''; NOT NULL"`
	Args         map[string]string `json:"args,omitempty" gorm:"column:args; type:json"`
	Description  string            `json:"description,omitempty" gorm:"column:description; default:''; NOT NULL"`
	WorkflowName string            `json:"workflow_name,omitempty" gorm:"column:workflow_name; default:''; NOT NULL"`
	ServiceID    int               `json:"service_id,omitempty" gorm:"column:service_id; default:0; NOT NULL"`
	gorm.Model
}

func (c *CI) SetServiceID(id int) {
	c.ServiceID = id
}

func (c *CI) SetWorkflowName(name string) {
	c.WorkflowName = name
}

type ServicesRepo interface {
	Save(context.Context, *Service) error
	Get(context.Context, int) (*Service, error)
	GetServices(context.Context) ([]*Service, error)
	Delete(context.Context, *Service) error
	SaveCI(context.Context, *CI) error
	GetCI(context.Context, int) (*CI, error)
	GetCIs(context.Context, int) ([]*CI, error)
	DeleteCI(ctx context.Context, ci *CI) error
	Deploy(context.Context, *Service, *CI) error
	GetOceanService(context.Context) (*Service, error)
}

type ServicesUseCase struct {
	repo ServicesRepo
	log  *log.Helper
}

func NewServicesUseCase(repo ServicesRepo, logger log.Logger) *ServicesUseCase {
	return &ServicesUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (s *ServicesUseCase) SaveService(ctx context.Context, service *Service) error {
	// 只保存service，忽略ci
	return s.repo.Save(ctx, service)
}

func (s *ServicesUseCase) GetService(ctx context.Context, id int) (*Service, error) {
	return s.repo.Get(ctx, id)
}

func (s *ServicesUseCase) GetServices(ctx context.Context) ([]*Service, error) {
	// 不包含ci，只有service列表
	return s.repo.GetServices(ctx)
}

func (s *ServicesUseCase) DeleteService(ctx context.Context, id int) error {
	service, err := s.GetService(ctx, id)
	if err != nil || service == nil {
		return err
	}
	// 需要把相关ci删除掉
	return s.repo.Delete(ctx, service)
}

func (s *ServicesUseCase) SaveCI(ctx context.Context, ci *CI) error {
	if ci.ServiceID == 0 {
		return errors.New("service id is empty")
	}
	service, err := s.GetService(ctx, ci.ServiceID)
	if err != nil {
		return err
	}
	if service == nil {
		return errors.New("service not found")
	}
	return s.repo.SaveCI(ctx, ci)
}

func (s *ServicesUseCase) GetCI(ctx context.Context, id int) (*CI, error) {
	return s.repo.GetCI(ctx, id)
}

func (s *ServicesUseCase) GetCIs(ctx context.Context, serviceID int) ([]*CI, error) {
	return s.repo.GetCIs(ctx, serviceID)
}

func (s *ServicesUseCase) DeleteCI(ctx context.Context, CIID int) error {
	ci, err := s.GetCI(ctx, CIID)
	if err != nil {
		return err
	}
	return s.repo.DeleteCI(ctx, ci)
}

func (s *ServicesUseCase) Deploy(ctx context.Context, CIID int) error {
	// 新建app，有Operator维护和创建资源
	ci, err := s.repo.GetCI(ctx, CIID)
	if err != nil {
		return err
	}
	service, err := s.GetService(ctx, ci.ServiceID)
	if err != nil {
		return err
	}
	return s.repo.Deploy(ctx, service, ci)
}

func (s *ServicesUseCase) GetOceanService(ctx context.Context) (*Service, error) {
	return s.repo.GetOceanService(ctx)
}
