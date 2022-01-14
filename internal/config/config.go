package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"sync"
	"webApi/pkg/logging"
)

type Config struct {
	IsDebug *bool `yaml:"is_debug"`
	Listen  struct {
		Type   string `yaml:"type"`
		BindIp string `yaml:"bind_ip"`
		Port   string `yaml:"port"`
	} `yaml:"listen"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		logger := logging.GetLogger()
		logger.Info("read configuration")
		instance = &Config{}
		if err := cleanenv.ReadConfig("config.yaml", instance); err != nil {
			description, err := cleanenv.GetDescription(instance, nil)
			logger.Warn(description)
			logger.Fatal(err)
		}
	})

	return instance
}
