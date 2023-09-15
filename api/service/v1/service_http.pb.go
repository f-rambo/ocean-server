// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.6.3
// - protoc             v3.19.4
// source: api/service/v1/service.proto

package v1

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

const OperationServiceServiceGetOceanService = "/service.v1.ServiceService/GetOceanService"

type ServiceServiceHTTPServer interface {
	GetOceanService(context.Context, *emptypb.Empty) (*Service, error)
}

func RegisterServiceServiceHTTPServer(s *http.Server, srv ServiceServiceHTTPServer) {
	r := s.Route("/")
	r.GET("/service/v1/get/ocean", _ServiceService_GetOceanService0_HTTP_Handler(srv))
}

func _ServiceService_GetOceanService0_HTTP_Handler(srv ServiceServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationServiceServiceGetOceanService)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetOceanService(ctx, req.(*emptypb.Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Service)
		return ctx.Result(200, reply)
	}
}

type ServiceServiceHTTPClient interface {
	GetOceanService(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *Service, err error)
}

type ServiceServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewServiceServiceHTTPClient(client *http.Client) ServiceServiceHTTPClient {
	return &ServiceServiceHTTPClientImpl{client}
}

func (c *ServiceServiceHTTPClientImpl) GetOceanService(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*Service, error) {
	var out Service
	pattern := "/service/v1/get/ocean"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationServiceServiceGetOceanService))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
