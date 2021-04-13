package config

import (
	"github.com/gopherd/doge/config"
)

// RamConfig represent config of ram service
type RamConfig struct {
	config.Config
}

func NewRamConfig() *RamConfig {
	return &RamConfig{}
}
