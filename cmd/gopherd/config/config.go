package config

import (
	"github.com/gopherd/doge/config"
)

// Config represent config of gopherd
type Config struct {
	config.Config
}

func NewConfig() *Config {
	return &Config{}
}
