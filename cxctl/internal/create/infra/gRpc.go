package infra

import (
	"github.com/gcx-17211270/cx/cxctl/internal/create/conf"
	pb "github.com/gcx-17211270/cx/cxctl/internal/create/pkg"
	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func NewGrpcClient(cfg *conf.Configuration, c *cli.Context) pb.CreateClient {
	conn, err := grpc.Dial(cfg.UnixAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v\n", err)
	}
	go func() {
		for {
			select {
			case <-c.Done():
				conn.Close()
			}
		}
	}()
	return pb.NewCreateClient(conn)
}
