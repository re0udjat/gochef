package config

import (
	"log"
	"os"
	"time"
)

type Config struct {
	Env        string `yaml:"env" env:"ENV" default:"dev"`
	Database   string `yaml:"db" required:"true"`
	HTTPServer `yaml:"http_server"`
}

type HTTPServer struct {
	Host        string        `yaml:"host" default:"localhost"`
	Port        uint16        `yaml:"port" default:"8080"`
	Timeout     time.Duration `yaml:"timeout" default:"3s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" default:"60s"`
}

func New() {
	configPath := os.Getenv("GOCHEF_CONFIG_PATH")
	if configPath == "" {
		log.Fatal("GOCHEF_CONFIG_PATH is not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("Config file does not exist: %s", configPath)
	}
}
