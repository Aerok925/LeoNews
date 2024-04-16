package configs

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/sethvargo/go-envconfig"
)

type AppConfig struct {
	Database DB `env:",prefix=DB_"`
}

func Init() (*AppConfig, error) {
	_ = godotenv.Overload()

	var config AppConfig
	err := envconfig.Process(context.Background(), &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
