package service

import (
	"github.com/urfave/cli/v2"
)

type Create interface {
	// Create 直接创建资源
	Create(ctx *cli.Context, s string) error

	// CreateF 从文件里创建资源
	CreateF(ctx *cli.Context, str string) error
}
