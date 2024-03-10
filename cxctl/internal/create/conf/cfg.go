package conf

import (
	"github.com/urfave/cli/v2"
)

type Configuration struct {
	UnixAddr string
}

func ResolveConfiguration(c *cli.Context) *Configuration {
	cfg := &Configuration{}
	cfg.UnixAddr = c.String("unix-socket-addr")
	return cfg
}
