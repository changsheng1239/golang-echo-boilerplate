package config

import (
	"errors"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type AppConfig struct {
	Env    string `env:"ENV" env-default:"development"`
	Server server
}

type server struct {
	Host string `env:"SERVER_HOST" env-default:"localhost"`
	Port string `env:"SERVER_PORT" env-default:"8080"`
}

func NewConfig() *AppConfig {
	var (
		cfg AppConfig
		err error
	)
	// load .env file into env variables
	err = godotenv.Load()
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		panic(err)
	}
	// read env variables into config struct
	err = cleanenv.ReadEnv(&cfg)
	if err != nil {
		panic(err)
	}

	return &cfg
}
