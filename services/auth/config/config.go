package config

import (
	"github.com/mohammadne/bookman/auth/internal/cache"
	"github.com/mohammadne/bookman/auth/internal/jwt"
	"github.com/mohammadne/bookman/auth/internal/network/grpc"
	"github.com/mohammadne/bookman/auth/internal/network/rest_api"
	"github.com/mohammadne/bookman/auth/pkg/logger"
)

type Config struct {
	// Environment string `default:"dev"`
	Logger     *logger.Config
	Jwt        *jwt.Config
	Cache      *cache.Config
	Rest       *rest_api.Config
	GrpcServer *grpc.Config
	GrpcUser   *grpc.Config
}
