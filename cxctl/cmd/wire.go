//go:build wireinject
// +build wireinject

package main

import (
	"github.com/gcx-17211270/cx/cxctl/internal/create/conf"
	"github.com/gcx-17211270/cx/cxctl/internal/create/infra"
	"github.com/gcx-17211270/cx/cxctl/internal/create/service"
	service2 "github.com/gcx-17211270/cx/cxctl/internal/service"
	"github.com/google/wire"
	"github.com/urfave/cli/v2"
)

func Initialize(ctx *cli.Context, cfg *conf.Configuration) *service2.Server {
	wire.Build(service2.NewService, service.NewCreate, infra.NewClient, infra.NewGrpcClient)
	return &service2.Server{}
}
