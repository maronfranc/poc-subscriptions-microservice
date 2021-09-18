package config

import (
	"log"
	"os"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Server struct {
		Port uint `yaml:"host" env:"SERVER_PORT" env-default:4000`
	} `yaml:"server"`
	Database struct {
		Host         string `yaml:"host" env:"MONGO_DB_HOST" env-default:"localhost"`
		Port         string `yaml:"port" env:"MONGO_DB_PORT" env-default:"27017"`
		DatabaseName string `yaml:"name" env:"MONGO_DB_NAME" env-default:"mongo-database"`
		Username     string `yaml:"user" env:"MONGO_DB_USER" env-default:"user"`
		Password     string `yaml:"password" env:"MONGO_DB_PASSWORD" env-default:"password"`
	} `yaml:"database"`
	Rabbitmq struct {
		Host     string `yaml:"host" env:"AMQP_HOST" env-default:"localhost"`
		Port     string `yaml:"port" env:"AMQP_PORT" env-default:"5672"`
		Username string `yaml:"user" env:"AMQP_USER" env-default:"guest"`
		Password string `yaml:"password" env:"AMQP_PASSWORD" env-default:"guest"`
	} `yaml:"message_broker"`
}

var _once sync.Once
var Cfg *Config

// GetConfig read configuration file
func GetConfig() *Config {
	path, exists := os.LookupEnv("CONFIG_PATH")

	var configPath = "config.yml"
	if exists {
		configPath = path
	}

	var cfg Config

	_once.Do(func() {
		err := cleanenv.ReadConfig(configPath, &cfg)
		if err != nil {
			log.Println("Failed to load config.", err)
			panic(err)
		}
	})
	return &cfg
}
