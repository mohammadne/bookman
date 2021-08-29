package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/mohammadne/bookman/user/internal/database"
	"github.com/mohammadne/bookman/user/internal/network/grpc"
	"github.com/mohammadne/bookman/user/internal/network/rest"
	"github.com/mohammadne/bookman/user/pkg/logger"
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
	cfg.GrpcServer = &grpc.Config{}
	cfg.GrpcAuth = &grpc.Config{}

	// process
	envconfig.MustProcess("user", cfg)
	envconfig.MustProcess("user_logger", cfg.Logger)
	envconfig.MustProcess("user_database", cfg.Database)
	envconfig.MustProcess("user_rest", cfg.Rest)
	envconfig.MustProcess("user_grpc", cfg.GrpcServer)
	envconfig.MustProcess("auth_grpc", cfg.GrpcAuth)

	return cfg
}
