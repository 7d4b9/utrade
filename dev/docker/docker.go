package docker

import (
	"github.com/docker/docker/client"
	"github.com/spf13/viper"
)

var Config = viper.New()

func init() {
	Config.AutomaticEnv()
	Config.SetEnvPrefix("mongo")
}

type Client struct {
	*client.Client
}
