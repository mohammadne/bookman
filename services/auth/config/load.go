package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/mohammadne/bookman/auth/internal/cache"
	"github.com/mohammadne/bookman/auth/internal/jwt"
	grpc_server "github.com/mohammadne/bookman/auth/internal/network/grpc/server"
	"github.com/mohammadne/bookman/auth/internal/network/rest"
	"github.com/mohammadne/bookman/auth/pkg/logger"
)

const (
	errLoadEnv = "Error loading .env file"
)

func Load(env Environment) *Config {
	if env == Development && godotenv.Load() != nil {
		panic(map[string]interface{}{"err": errLoadEnv})
	}

	// initialize
	cfg := new(Config)
	cfg.Logger = &logger.Config{}
	cfg.Jwt = &jwt.Config{}
	cfg.Cache = &cache.Config{}
	cfg.Rest = &rest.Config{}
	cfg.Grpc = &grpc_server.Config{}

	// process
	envconfig.MustProcess("auth", cfg)
	envconfig.MustProcess("auth_logger", cfg.Logger)
	envconfig.MustProcess("auth_jwt", cfg.Jwt)
	envconfig.MustProcess("auth_cache", cfg.Cache)
	envconfig.MustProcess("auth_rest", cfg.Rest)
	envconfig.MustProcess("auth_grpc", cfg.Grpc)

	return cfg
}
