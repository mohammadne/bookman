package configs

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/mohammadne/bookman/user/pkg/database"
	"github.com/mohammadne/bookman/user/pkg/logger"
)

type migrate struct {
	Logger   *logger.Config
	Database *database.Config
}

func Migrate(env string) *migrate {
	config := &migrate{}

	switch env {
	case "prod":
		config.loadProd()
	default:
		config.loadDev()
	}

	return config
}

func (config *migrate) loadProd() {
	config.Logger = &logger.Config{}
	config.Database = &database.Config{}

	// process
	envconfig.MustProcess("library", config)
	envconfig.MustProcess("library_logger", config.Logger)
	envconfig.MustProcess("library_database", config.Database)
}

func (config *migrate) loadDev() {
	config.Logger = &logger.Config{
		Development:      true,
		EnableCaller:     true,
		EnableStacktrace: false,
		Encoding:         "console",
		Level:            "warn",
	}

	config.Database = &database.Config{
		Username: "root",
		Password: "password",
		Host:     "localhost",
		Port:     "3306",
		Schema:   "bookman",
	}
}
