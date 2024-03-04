// Пакет для получения экземпляра конфигурации из yaml и переменных окружения
package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	App     `yaml:"app"`
	HTTP    `yaml:"http"`
	Logger  `yaml:"logger"`
	Storage `yaml:"storage"`
}

type App struct {
	Name    string `env-required:"true" yaml:"name" env:"APP_NAME"`
	Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
}

type HTTP struct {
	Port         string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
	WaitingClose int    `yaml:"waiting_timeout_close" env:"WAIT_CLOSE"`
}

type Logger struct {
	Level string `env-required:"true" yaml:"log_env" env:"LOG_ENV"`
}

type Storage struct {
	Addr string `env-required:"true" yaml:"addr" env:"STOR_ADDR"`
	Port string `env-required:"true" yaml:"port" env:"STOR_PORT"`
	Name string `env-required:"true" yaml:"name" env:"STOR_NAME"`
	User string `env-required:"true" yaml:"user" env:"STOR_USER"`
	Pass string `env-required:"true" yaml:"pass" env:"STOR_PASS"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./config/config.yaml", cfg)
	if err != nil {
		return nil, err
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, fmt.Errorf("config error %w", err)
	}

	return cfg, nil
}
