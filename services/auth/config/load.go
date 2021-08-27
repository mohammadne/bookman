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

func Load(e Environment) *Config {
	// loads environment variables from .env into apllication
	if err := godotenv.Load(); err != nil {
		panic(map[string]interface{}{"err": err, "msg": errLoadEnv})
	}

	switch e {
	case Production:
		return loadProd()
	case Development:
		return loadDev()
	}

	return nil
}

func loadProd() *Config {
	// initialize
	cfg := new(Config)
	cfg.Logger = &logger.Config{}
	cfg.Jwt = &jwt.Config{}
	cfg.Cache = &cache.Config{}
	cfg.Rest = &rest.Config{}
	cfg.Grpc = &grpc_server.Config{}

	// process
	envconfig.MustProcess("bookman_auth", cfg)
	envconfig.MustProcess("bookman_auth_logger", cfg.Logger)
	envconfig.MustProcess("bookman_auth_jwt", cfg.Jwt)
	envconfig.MustProcess("bookman_auth_cache", cfg.Cache)
	envconfig.MustProcess("bookman_auth_rest", cfg.Rest)
	envconfig.MustProcess("bookman_auth_grpc", cfg.Grpc)

	return cfg
}

func loadDev() *Config {
	return &Config{
		Logger: &logger.Config{
			Development:      true,
			EnableCaller:     false,
			EnableStacktrace: false,
			Encoding:         "console",
			Level:            "warn",
		},
		Jwt: &jwt.Config{
			AccessSecret:   "access_secret",
			AccessExpires:  1,
			RefreshSecret:  "refresh_secret",
			RefreshExpires: 24 * 7,
		},
		Cache: &cache.Config{
			URL: "localhost:6379",
		},
		Rest: &rest.Config{
			URL: "localhost:8080",
		},
		Grpc: &grpc_server.Config{
			URL: "localhost:4040",
		},
	}
}
