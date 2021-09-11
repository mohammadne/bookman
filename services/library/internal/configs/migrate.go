package configs

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/mohammadne/bookman/library/internal/database"
	"github.com/mohammadne/bookman/library/pkg/logger"
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
	{
		// TODO: temp passing config
		if err := godotenv.Load(); err != nil {
			panic(err)
		}
	}

	config.Logger = &logger.Config{}
	config.Database = &database.Config{}

	// process
	envconfig.MustProcess("library", config)
	envconfig.MustProcess("library_logger", config.Logger)
	envconfig.MustProcess("library_database", config.Database)
}

func (config *migrate) loadDev() {
	config.Logger = &logger.Config{}

	config.Database = &database.Config{}
}
