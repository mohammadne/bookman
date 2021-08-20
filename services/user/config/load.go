package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/mohammadne/bookman/user/internal/logger"
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
	}
}
