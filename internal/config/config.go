package config

import (
	"os"
	"strings"

	"github.com/go-playground/validator"
	_ "github.com/joho/godotenv/autoload"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/v2"
	"github.com/rs/zerolog"
)

type Config struct {
	Logging LogConfig
	Env     string `koanf:"primary_env"`
	Cache   CacheConfig
}

type LogConfig struct {
	Level  string `koanf:"log_level" validate:"required"`
	Format string `koanf:"log_format" validate:"required"`
}

type CacheConfig struct {
	MaxSize         int    `koanf:"cache_max_size" validate:"required"`
	CleanupInterval string `koanf:"cache_cleanup_interval" validate:"required"`
	EvictionPolicy  string `koanf:"cache_eviction_policy" validate:"required"`
}

func LoadConfig() (*Config, error) {
	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout}).With().Timestamp().Logger()
	k := koanf.New(".")
	err := k.Load(env.Provider("IN_MEMORY_", ".", func(s string) string {
		return strings.ToLower(strings.TrimPrefix(s, "IN_MEMORY_"))
	}), nil)

	if err != nil {
		logger.Fatal().Err(err).Msg("could not load initial env variables")
	}
	mainConfig := &Config{}

	err = k.Unmarshal("", mainConfig)
	if err != nil {
		logger.Fatal().Err(err).Msg("could not unmarshal main config")
	}

	validate := validator.New()
	err = validate.Struct(mainConfig)
	if err != nil {
		logger.Fatal().Err(err).Msg("config validation failed")
	}

	return mainConfig, nil

}
