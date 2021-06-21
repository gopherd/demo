package config

import (
	"github.com/gopherd/doge/config"
)

// RamConfig represent config of ram service
type RamConfig struct {
	config.BaseConfig
}

func NewRamConfig() *RamConfig {
	cfg := &RamConfig{}
	cfg.Core.ID = 1
	return cfg
}
