package service

import (
	"errors"
	"fmt"
	"github.com/gcx-17211270/cx/cxctl/internal/create/conf"
	pb "github.com/gcx-17211270/cx/cxctl/internal/create/pkg"
	"github.com/gcx-17211270/cx/cxctl/internal/service"
	"github.com/urfave/cli/v2"
	"io"
	"os"
	"strings"
)

type Create struct {
	ctx          *cli.Context
	cfg          *conf.Configuration
	createClient pb.CreateClient
}

var _ service.Create = &Create{}

func NewCreate(ctx *cli.Context, cfg *conf.Configuration, createClient pb.CreateClient) service.Create {
	return &Create{ctx: ctx, cfg: cfg, createClient: createClient}
}

// Create 直接创建资源
func (*Create) Create(ctx *cli.Context, s string) error {
	return fmt.Errorf("直接创建资源暂时不支持")
}

// CreateF 从文件里创建资源
func (s *Create) CreateF(ctx *cli.Context, str string) error {
	// 分本地文件和网络文件
	if strings.HasPrefix(str, "http") {
		return fmt.Errorf("从网络上获取部署描述文件创建资源暂时不支持")

	} else {
		file, err := os.Open(str)
		if err != nil {
			return errors.Join(err, fmt.Errorf("cannot open file. maybe it is from network but have wrong scheme."))
		}
		all, err := io.ReadAll(file)
		if err != nil {
			return errors.Join(err, fmt.Errorf("read file content error occur"))
		}
		return s.create(ctx, string(all))
	}
	return nil
}

func (c *Create) create(ctx *cli.Context, content string) error {
	create, err := c.createClient.Create(ctx.Context, &pb.CreateRequest{RequestId: 1})
	fmt.Println(create)
	return err
}
