package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/mohammadne/bookman/user/internal/database"
	grpc_server "github.com/mohammadne/bookman/user/internal/network/grpc/server"
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
	envconfig.MustProcess("user", cfg)
	envconfig.MustProcess("user_logger", cfg.Logger)
	envconfig.MustProcess("user_database", cfg.Database)
	envconfig.MustProcess("user_rest", cfg.Rest)
	envconfig.MustProcess("user_grpc", cfg.Grpc)

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
			Password: "password",
			Host:     "127.0.0.1:3306",
			Schema:   "bookman",
		},
		Rest: &rest.Config{
			URL: "localhost:8081",
		},
		Grpc: &grpc_server.Config{
			URL: "localhost:4041",
		},
	}
}
