package main

import (
	"context"
	"github.com/gcx-17211270/cx/cxctl/internal/create/conf"
	pb "github.com/gcx-17211270/cx/cxctl/internal/create/pkg"
	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

type server struct {
	pb.UnimplementedCreateServer
}

func main() {

	var cfg *conf.Configuration
	app := &cli.App{
		Name:        "cxctl",
		Usage:       "用来对资源进行增删改查",
		UsageText:   "cxctl command [command options] [arguments...]",
		Description: "a clumsy imitation to kubectl",
		Before: func(c *cli.Context) error {
			cfg = conf.ResolveConfiguration(c)
			return nil
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "unix-socket-addr",
				Usage: "unix socket addr to cxd",
				Value: "/root/code/cx/cx.ipc",
			},
		},
		Commands: []*cli.Command{
			{
				Name: "start",
				Action: func(context *cli.Context) error {
					lis, err := net.Listen("unix", cfg.UnixAddr)
					if err != nil {
						log.Fatalf("failed to listen: %v", err)
					}
					s := grpc.NewServer()
					pb.RegisterCreateServer(s, &server{})
					if err = s.Serve(lis); err != nil {
						log.Fatalf("failed to serve: %v", err)
					}
					return nil
				},
			},
		},
	}
	app.Run(os.Args)
}

func (s *server) Create(c context.Context, in *pb.CreateRequest) (*pb.CreateReply, error) {
	log.Printf("Received: %v", in)
	return &pb.CreateReply{Ok: true}, nil
}
