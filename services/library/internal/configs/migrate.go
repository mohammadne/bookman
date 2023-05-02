package configs

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/mohammadne/bookman/library/internal/database"
	"github.com/mohammadne/bookman/library/pkg/logger"
	"github.com/mohammadne/bookman/library/pkg/tracer"
)

type migrate struct {
	Logger   *logger.Config
	Tracer   *tracer.Config
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
	config.Tracer = &tracer.Config{}
	config.Database = &database.Config{}

	// process
	envconfig.MustProcess("library", config)
	envconfig.MustProcess("library_logger", config.Logger)
	envconfig.MustProcess("library_tracer", config.Tracer)
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

	config.Tracer = &tracer.Config{
		Enabled:    false,
		Host:       "localhost",
		Port:       "",
		SampleRate: 0,
		Namespace:  "bookman",
		Subsystem:  "library",
	}

	config.Database = &database.Config{
		Driver:       "mysql",
		Host:         "localhost",
		Port:         "3306",
		User:         "root",
		Password:     "password",
		DatabaseName: "bookman",
		SSLMode:      "",
	}
}
