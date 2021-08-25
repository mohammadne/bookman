package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/mohammadne/bookman/auth/internal/cache"
	"github.com/mohammadne/bookman/auth/internal/jwt"
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

	// process
	envconfig.MustProcess("", cfg)
	envconfig.MustProcess("logger", cfg.Logger)
	envconfig.MustProcess("jwt", cfg.Jwt)
	envconfig.MustProcess("cache", cfg.Cache)
	envconfig.MustProcess("rest", cfg.Rest)

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
	}
}
