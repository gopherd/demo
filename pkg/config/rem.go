package config

import (
	"github.com/gopherd/doge/config"
)

// RemConfig represent config of rem service
type RemConfig struct {
	config.Config
}

func NewRemConfig() *RemConfig {
	return &RemConfig{}
}
