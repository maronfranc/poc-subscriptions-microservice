package config

import (
	"fmt"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

var _once sync.Once

type Config struct {
	Server struct {
		Port uint `yaml:"host"`
	}
	Database struct {
		Host         string `yaml:"host"`
		Port         string `yaml:"port"`
		DatabaseName string `yaml:"name"`
		Username     string `yaml:"user"`
		Password     string `yaml:"password"`
	} `yaml:"database"`
}

func GetConfig() *Config {
	var (
		configPath = "config.yml"
		cfg        Config
	)

	_once.Do(func() {
		err := cleanenv.ReadConfig(configPath, &cfg)
		if err != nil {
			fmt.Println("Failed to load config.", err)
			panic(err)
		}
	})

	return &cfg
}
