// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.6.3
// - protoc             v4.25.0
// source: api/cluster/v1alpha1/cluster.proto

package v1alpha1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationClusterServiceApply = "/cluster.v1alpha1.ClusterService/Apply"
const OperationClusterServiceDelete = "/cluster.v1alpha1.ClusterService/Delete"
const OperationClusterServiceGet = "/cluster.v1alpha1.ClusterService/Get"
const OperationClusterServiceSave = "/cluster.v1alpha1.ClusterService/Save"

type ClusterServiceHTTPServer interface {
	Apply(context.Context, *ClusterName) (*Msg, error)
	Delete(context.Context, *ClusterID) (*Msg, error)
	Get(context.Context, *emptypb.Empty) (*Clusters, error)
	Save(context.Context, *Cluster) (*Msg, error)
}

func RegisterClusterServiceHTTPServer(s *http.Server, srv ClusterServiceHTTPServer) {
	r := s.Route("/")
	r.GET("/cluster/v1alpha1/get", _ClusterService_Get0_HTTP_Handler(srv))
	r.POST("/cluster/v1alpha1/save", _ClusterService_Save0_HTTP_Handler(srv))
	r.DELETE("/cluster/v1alpha1/delete/{id}", _ClusterService_Delete0_HTTP_Handler(srv))
	r.POST("/cluster/v1alpha1/apply", _ClusterService_Apply0_HTTP_Handler(srv))
}

func _ClusterService_Get0_HTTP_Handler(srv ClusterServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationClusterServiceGet)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Get(ctx, req.(*emptypb.Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Clusters)
		return ctx.Result(200, reply)
	}
}

func _ClusterService_Save0_HTTP_Handler(srv ClusterServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in Cluster
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationClusterServiceSave)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Save(ctx, req.(*Cluster))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Msg)
		return ctx.Result(200, reply)
	}
}

func _ClusterService_Delete0_HTTP_Handler(srv ClusterServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ClusterID
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationClusterServiceDelete)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Delete(ctx, req.(*ClusterID))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Msg)
		return ctx.Result(200, reply)
	}
}

func _ClusterService_Apply0_HTTP_Handler(srv ClusterServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ClusterName
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationClusterServiceApply)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Apply(ctx, req.(*ClusterName))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Msg)
		return ctx.Result(200, reply)
	}
}

type ClusterServiceHTTPClient interface {
	Apply(ctx context.Context, req *ClusterName, opts ...http.CallOption) (rsp *Msg, err error)
	Delete(ctx context.Context, req *ClusterID, opts ...http.CallOption) (rsp *Msg, err error)
	Get(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *Clusters, err error)
	Save(ctx context.Context, req *Cluster, opts ...http.CallOption) (rsp *Msg, err error)
}

type ClusterServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewClusterServiceHTTPClient(client *http.Client) ClusterServiceHTTPClient {
	return &ClusterServiceHTTPClientImpl{client}
}

func (c *ClusterServiceHTTPClientImpl) Apply(ctx context.Context, in *ClusterName, opts ...http.CallOption) (*Msg, error) {
	var out Msg
	pattern := "/cluster/v1alpha1/apply"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationClusterServiceApply))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ClusterServiceHTTPClientImpl) Delete(ctx context.Context, in *ClusterID, opts ...http.CallOption) (*Msg, error) {
	var out Msg
	pattern := "/cluster/v1alpha1/delete/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationClusterServiceDelete))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ClusterServiceHTTPClientImpl) Get(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*Clusters, error) {
	var out Clusters
	pattern := "/cluster/v1alpha1/get"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationClusterServiceGet))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ClusterServiceHTTPClientImpl) Save(ctx context.Context, in *Cluster, opts ...http.CallOption) (*Msg, error) {
	var out Msg
	pattern := "/cluster/v1alpha1/save"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationClusterServiceSave))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
