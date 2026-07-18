package config

import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
	AppPort string `env:"APP_PORT" env-default:"8080"`

	DatabaseHost     string `env:"DATABASE_HOST" env-required:"true"`
	DatabasePort     int    `env:"DATABASE_PORT" env-default:"5432"`
	DatabaseUser     string `env:"DATABASE_USER" env-required:"true"`
	DatabasePassword string `env:"DATABASE_PASSWORD" env-required:"true"`
	DatabaseName     string `env:"DATABASE_NAME" env-required:"true"`
	DatabaseSSLMode  string `env:"DATABASE_SSL_MODE" env-default:"disable"`
}

func LoadConfig() (*Config, error) {
	cfg := &Config{}

	_ = cleanenv.ReadConfig(".env", cfg)

	if err := cleanenv.ReadEnv(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
