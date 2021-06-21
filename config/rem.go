package config

import (
	"github.com/gopherd/doge/config"
)

// RemConfig represent config of rem service
type RemConfig struct {
	config.BaseConfig
}

func NewRemConfig() *RemConfig {
	cfg := &RemConfig{}
	cfg.Core.ID = 2
	return cfg
}
