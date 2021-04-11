package config

import (
	"github.com/gopherd/doge/config"
)

// GopherdConfig represent config of gopherd
type GopherdConfig struct {
	config.Config
}

func NewGopherdConfig() *GopherdConfig {
	return &GopherdConfig{}
}
