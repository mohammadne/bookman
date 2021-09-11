package configs

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/mohammadne/bookman/library/internal/database"
	"github.com/mohammadne/bookman/library/internal/network/rest_api"
	"github.com/mohammadne/bookman/library/pkg/logger"
)

type server struct {
	Logger   *logger.Config
	Database *database.Config
	Rest     *rest_api.Config
}

func Server(env string) *server {
	switch env {
	case "prod":
		return prodServer()
	default:
		return devServer()
	}
}

func prodServer() *server {
	{
		// TODO: temp passing config
		if err := godotenv.Load(); err != nil {
			panic(err)
		}
	}

	config := new(server)

	config.Logger = &logger.Config{}
	config.Database = &database.Config{}
	config.Rest = &rest_api.Config{}

	// process
	envconfig.MustProcess("library", config)
	envconfig.MustProcess("library_logger", config.Logger)
	envconfig.MustProcess("library_database", config.Database)
	envconfig.MustProcess("library_rest_api", config.Rest)

	return config
}

func devServer() *server {
	return &server{}
}
