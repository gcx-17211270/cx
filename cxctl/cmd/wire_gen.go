// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/gcx-17211270/cx/cxctl/internal/create/conf"
	"github.com/gcx-17211270/cx/cxctl/internal/create/infra"
	"github.com/gcx-17211270/cx/cxctl/internal/create/service"
	service2 "github.com/gcx-17211270/cx/cxctl/internal/service"
	"github.com/urfave/cli/v2"
)

// Injectors from wire.go:

func Initialize(ctx *cli.Context, cfg *conf.Configuration) *service2.Server {
	createClient := infra.NewGrpcClient(cfg, ctx)
	create := service.NewCreate(ctx, cfg, createClient)
	client := infra.NewClient()
	server := service2.NewService(create, client)
	return server
}
