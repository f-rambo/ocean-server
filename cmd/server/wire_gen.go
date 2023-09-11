// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/f-rambo/ocean/internal/biz"
	"github.com/f-rambo/ocean/internal/conf"
	"github.com/f-rambo/ocean/internal/data"
	"github.com/f-rambo/ocean/internal/server"
	"github.com/f-rambo/ocean/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	greeterRepo := data.NewGreeterRepo(dataData, logger)
	greeterUsecase := biz.NewGreeterUsecase(greeterRepo, logger)
	greeterService := service.NewGreeterService(greeterUsecase)
	clusterRepo := data.NewClusterRepo(dataData, logger)
	clusterUsecase := biz.NewClusterUseCase(clusterRepo, logger)
	clusterService := service.NewClusterService(clusterUsecase, confData)
	appRepo := data.NewAppRepo(dataData, logger)
	appUsecase := biz.NewAppUsecase(appRepo, logger)
	appService := service.NewAppService(appUsecase)
	grpcServer := server.NewGRPCServer(confServer, greeterService, clusterService, appService, logger)
	httpServer := server.NewHTTPServer(confServer, greeterService, clusterService, appService, logger)
	app := newApp(logger, grpcServer, httpServer)
	return app, func() {
		cleanup()
	}, nil
}
