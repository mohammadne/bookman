package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/mohammadne/bookman/auth/internal/cache"
	"github.com/mohammadne/bookman/auth/internal/jwt"
	"github.com/mohammadne/bookman/auth/internal/web/rest"
	"github.com/mohammadne/bookman/auth/pkg/logger"
)

func Load(e Environment) *Config {
	switch e {
	case Production:
		return loadProd()
	case Development:
		return loadDev()
	}

	return nil
}

func loadProd() (cfg *Config) {
	// initialize
	cfg.Logger = &logger.Config{}
	cfg.Jwt = &jwt.Config{}
	cfg.Cache = &cache.Config{}
	cfg.Rest = &rest.Config{}

	// process
	envconfig.MustProcess("bookman_user", cfg)
	envconfig.MustProcess("bookman_user_logger", cfg.Logger)
	envconfig.MustProcess("bookman_user_jwt", cfg.Jwt)
	envconfig.MustProcess("bookman_user_cache", cfg.Cache)
	envconfig.MustProcess("bookman_user_rest", cfg.Rest)

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
