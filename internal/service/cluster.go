package service

import (
	"context"
	"fmt"

	v1 "github.com/f-rambo/ocean/api/cluster/v1"
	"github.com/f-rambo/ocean/internal/biz"
	"github.com/f-rambo/ocean/internal/conf"

	"google.golang.org/protobuf/types/known/emptypb"
)

type ClusterService struct {
	v1.UnimplementedClusterServiceServer
	uc *biz.ClusterUsecase
	c  *conf.Data
}

func NewClusterService(uc *biz.ClusterUsecase, c *conf.Data) *ClusterService {
	return &ClusterService{uc: uc, c: c}
}

func (c *ClusterService) Get(ctx context.Context, _ *emptypb.Empty) (*v1.Clusters, error) {
	clusters, err := c.uc.Get(ctx)
	if err != nil {
		return nil, err
	}
	data := &v1.Clusters{}
	for _, cluster := range clusters {
		v := &v1.Cluster{
			Id:          int32(cluster.ID),
			ClusterName: cluster.Name,
			Config:      string(cluster.Config),
			Addons:      string(cluster.Addons),
			Applyed:     cluster.Deployed,
		}
		for _, node := range cluster.Nodes {
			v.Nodes = append(v.Nodes, &v1.Node{
				Id:           int32(node.ID),
				Name:         node.Name,
				Host:         node.Host,
				Role:         node.Role,
				User:         node.User,
				Password:     node.Password,
				SudoPassword: node.SudoPassword,
			})
		}
		data.Clusters = append(data.Clusters, v)
	}
	return data, nil
}

func (c *ClusterService) Save(ctx context.Context, cluster *v1.Cluster) (*v1.Msg, error) {
	bizCluster := &biz.Cluster{
		ID:   int(cluster.Id),
		Name: cluster.ClusterName,
	}
	bizCluster.Config = []byte(cluster.Config)
	bizCluster.Addons = []byte(cluster.Addons)
	for _, node := range cluster.Nodes {
		bizCluster.Nodes = append(bizCluster.Nodes, biz.Node{
			ID:       int(node.Id),
			Name:     node.Name,
			Host:     node.Host,
			Role:     node.Role,
			User:     node.User,
			Password: node.SudoPassword,
		})
	}
	err := c.uc.Save(ctx, bizCluster)
	if err != nil {
		return nil, err
	}
	// 提示信息
	admin := c.c.Semaphore.GetAdmin()
	adminPassword := c.c.Semaphore.GetAdminPassword()
	host := c.c.Semaphore.GetHost()
	port := c.c.Semaphore.GetPort()
	msg := &v1.Msg{
		Message: fmt.Sprintf(`
		登录Semaphore查看任务执行进度:
		url : http://%s:%d
		admin: %s
		password: %s
		`, host, port, admin, adminPassword),
		Reason: v1.ErrorReason_SUCCEED,
	}
	return msg, nil
}

// Delete
func (c *ClusterService) Delete(ctx context.Context, clusterID *v1.ClusterID) (*v1.Msg, error) {
	err := c.uc.Delete(ctx, int(clusterID.Id))
	if err != nil {
		return nil, err
	}
	msg := &v1.Msg{
		Message: "删除成功",
		Reason:  v1.ErrorReason_SUCCEED,
	}
	return msg, nil
}

// apply
func (c *ClusterService) Apply(ctx context.Context, clusterName *v1.ClusterName) (*v1.Msg, error) {
	err := c.uc.Apply(ctx, clusterName.Name)
	if err != nil {
		return nil, err
	}
	return &v1.Msg{Message: "apply success"}, nil
}
