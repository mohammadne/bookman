package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/mohammadne/bookman/library/pkg/database"
	"github.com/mohammadne/bookman/library/pkg/logger"
	"github.com/mohammadne/bookman/library/pkg/rest"
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
	cfg.Database = &database.Config{}
	cfg.Rest = &rest.Config{}

	// process
	envconfig.MustProcess("library", cfg)
	envconfig.MustProcess("library_logger", cfg.Logger)
	envconfig.MustProcess("library_database", cfg.Database)
	envconfig.MustProcess("library_rest", cfg.Rest)

	return cfg
}
