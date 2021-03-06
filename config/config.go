package config

import (
	"github.com/lecex/core/config"
	"github.com/lecex/core/env"
)

// Conf 配置
var Conf config.Config = config.Config{
	Name:    env.Getenv("MICRO_API_NAMESPACE", "go.micro.api.") + "device",
	Version: "v1.0.3",
}
