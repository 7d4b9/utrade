package mongo

import (
	"github.com/7d4b9/utrade/dev/docker"
	"github.com/spf13/viper"
)

var Config = viper.New()

func init() {
	Config.AutomaticEnv()
	Config.SetEnvPrefix("mongo")
}

type Client struct {
	*docker.Client
}
