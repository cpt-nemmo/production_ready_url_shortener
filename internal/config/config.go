package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env          string `yaml:"env" env-required:"true"`
	Storage_path string `yaml:"storage_path" env-required:"true"`
	HTTPServer   `yaml:"http_server"`
}

type HTTPServer struct {
	Address      string        `yaml:"addres" env-required:"true" env-default:"localhost:8080"`
	Timeout      time.Duration `yaml:"timeout" env-default:"4s" env-required:"true"`
	Idle_timeout time.Duration `yaml:"idle_timeout" env-default:"60s" env-required:"true"`
}

func MustLoadConf() *Config {
	configPath := "./config/local.yml"

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("can't read config: %s", err)
	}

	return &cfg
}
