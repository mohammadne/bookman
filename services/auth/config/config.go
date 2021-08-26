package config

import (
	"github.com/mohammadne/bookman/auth/internal/cache"
	"github.com/mohammadne/bookman/auth/internal/jwt"
	grpc_server "github.com/mohammadne/bookman/auth/internal/network/grpc/server"
	"github.com/mohammadne/bookman/auth/internal/network/rest"
	"github.com/mohammadne/bookman/auth/pkg/logger"
)

type Config struct {
	// Environment string `default:"dev"`
	Logger *logger.Config
	Jwt    *jwt.Config
	Cache  *cache.Config
	Rest   *rest.Config
	Grpc   *grpc_server.Config
}
