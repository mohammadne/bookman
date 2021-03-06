package configs

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/mohammadne/bookman/user/internal/network/grpc"
	"github.com/mohammadne/bookman/user/internal/network/rest_api"
	"github.com/mohammadne/bookman/user/pkg/database"
	"github.com/mohammadne/bookman/user/pkg/logger"
	"github.com/mohammadne/bookman/user/pkg/tracer"
)

type server struct {
	Logger   *logger.Config
	Tracer   *tracer.Config
	Database *database.Config
	RestApi  *rest_api.Config
	UserGrpc *grpc.Config
	AuthGrpc *grpc.Config
}

func Server(env string) *server {
	config := &server{}

	switch env {
	case "prod":
		config.loadProd()
	default:
		config.loadDev()
	}

	return config
}

func (config *server) loadProd() {
	config.Logger = &logger.Config{}
	config.Tracer = &tracer.Config{}
	config.Database = &database.Config{}
	config.RestApi = &rest_api.Config{}
	config.UserGrpc = &grpc.Config{}
	config.AuthGrpc = &grpc.Config{}

	// process
	envconfig.MustProcess("library", config)
	envconfig.MustProcess("library_logger", config.Logger)
	envconfig.MustProcess("library_tracer", config.Tracer)
	envconfig.MustProcess("library_database", config.Database)
	envconfig.MustProcess("library_rest_api", config.RestApi)
	envconfig.MustProcess("user_grpc", config.AuthGrpc)
	envconfig.MustProcess("auth_grpc", config.AuthGrpc)
}

func (config *server) loadDev() {
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
		Subsystem:  "user",
	}

	config.Database = &database.Config{
		Username: "root",
		Password: "password",
		Host:     "localhost",
		Port:     "3306",
		Schema:   "bookman",
	}

	config.RestApi = &rest_api.Config{
		Host: "localhost",
		Port: "8081",
	}

	config.UserGrpc = &grpc.Config{
		Host: "localhost",
		Port: "4041",
	}

	config.AuthGrpc = &grpc.Config{
		Host: "localhost",
		Port: "4040",
	}
}
