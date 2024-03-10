package main

import (
	"fmt"
	"github.com/gcx-17211270/cx/cxctl/internal/create/conf"
	"github.com/gcx-17211270/cx/cxctl/internal/service"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {
	var server *service.Server
	app := &cli.App{
		Name:        "cxctl",
		Usage:       "用来对资源进行增删改查",
		UsageText:   "cxctl command [command options] [arguments...]",
		Description: "a clumsy imitation to kubectl",
		Before: func(c *cli.Context) error {
			cfg := conf.ResolveConfiguration(c)
			server = Initialize(c, cfg)
			return nil
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "unix-socket-addr",
				Usage: "unix socket addr to cxd",
				Value: "unix:/root/code/cx/cx.ipc",
			},
		},
	}
	app.Commands = append(app.Commands, &cli.Command{
		Name:      "create",
		Usage:     "cxctl create [command options] [arguments]",
		UsageText: "创建资源，输入可以从两个地方选择，-f本地文件，-h网络文件",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "f",
				Usage: "使用来创建资源（本地/网络）",
				Action: func(context *cli.Context, s string) error {
					return server.CreateF(context, s)
				},
			},
		},
		Action: func(context *cli.Context) error { return nil },
	})
	app.Commands = append(app.Commands, &cli.Command{
		Name:      "apply",
		Usage:     "cxctl apply [command options] [arguments]",
		UsageText: "创建或更新资源，输入可以从两个地方选择，-f本地文件，-h网络文件",
		Action: func(context *cli.Context) error {
			fmt.Println("cxctl apply finished")
			return nil
		},
	})
	app.Commands = append(app.Commands, &cli.Command{
		Name:      "delete",
		Usage:     "cxctl delete [command options] [arguments]",
		UsageText: "删除资源，输入可以是资源以及资源名称；或者是从文件获取，-f本地文件，-h网络文件",
		Action: func(context *cli.Context) error {
			fmt.Println("cxctl delete finished")
			return nil
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name: "f",
				Action: func(context *cli.Context, s string) error {
					fmt.Println("cxctl delete from local file finished-", s)
					return nil
				},
			},
			&cli.StringFlag{
				Name: "-h1",
				Action: func(context *cli.Context, s string) error {
					fmt.Println("cxctl delete from network file finished-", s)
					return nil
				},
			},
		},
	})
	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
}
