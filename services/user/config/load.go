package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/mohammadne/bookman/user/internal/database"
	"github.com/mohammadne/bookman/user/internal/network/rest"
	"github.com/mohammadne/bookman/user/pkg/logger"
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

	// process
	envconfig.MustProcess("bookman_user", cfg)
	envconfig.MustProcess("bookman_user_logger", cfg.Logger)
	envconfig.MustProcess("bookman_user_database", cfg.Logger)
	envconfig.MustProcess("bookman_user_web", cfg.Logger)

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
		Database: &database.Config{
			Username: "root",
			Password: "lenovoe590",
			Host:     "127.0.0.1:3306",
			Schema:   "bookman",
		},
		Rest: &rest.Config{
			URL: "localhost:8080",
		},
	}
}
