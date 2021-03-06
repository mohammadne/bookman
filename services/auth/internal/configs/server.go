package configs

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/mohammadne/bookman/auth/internal/cache"
	"github.com/mohammadne/bookman/auth/internal/jwt"
	"github.com/mohammadne/bookman/auth/internal/network/grpc"
	"github.com/mohammadne/bookman/auth/internal/network/rest_api"
	"github.com/mohammadne/bookman/auth/pkg/logger"
	"github.com/mohammadne/bookman/auth/pkg/tracer"
)

type server struct {
	Logger   *logger.Config
	Tracer   *tracer.Config
	Jwt      *jwt.Config
	Cache    *cache.Config
	RestApi  *rest_api.Config
	AuthGrpc *grpc.Config
	UserGrpc *grpc.Config
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
	config.Jwt = &jwt.Config{}
	config.Cache = &cache.Config{}
	config.RestApi = &rest_api.Config{}
	config.AuthGrpc = &grpc.Config{}
	config.UserGrpc = &grpc.Config{}

	// process
	envconfig.MustProcess("library", config)
	envconfig.MustProcess("library_logger", config.Logger)
	envconfig.MustProcess("library_tracer", config.Tracer)
	envconfig.MustProcess("library_jwt", config.Jwt)
	envconfig.MustProcess("library_cache", config.Cache)
	envconfig.MustProcess("library_rest_api", config.RestApi)
	envconfig.MustProcess("auth_grpc", config.AuthGrpc)
	envconfig.MustProcess("user_grpc", config.UserGrpc)
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
		Subsystem:  "library",
	}

	config.Jwt = &jwt.Config{
		AccessSecret:   "4ykllbAMpImzZlE",
		AccessExpires:  1,
		RefreshSecret:  "sezXJL0Jl5kO0Du",
		RefreshExpires: 168,
	}

	config.Cache = &cache.Config{
		URL: "localhost:6379",
	}

	config.RestApi = &rest_api.Config{
		Host: "localhost",
		Port: "8080",
	}

	config.AuthGrpc = &grpc.Config{
		Host: "localhost",
		Port: "4040",
	}

	config.UserGrpc = &grpc.Config{
		Host: "localhost",
		Port: "4041",
	}
}
