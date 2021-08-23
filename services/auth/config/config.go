package config

import (
	"github.com/mohammadne/bookman/auth/internal/cache"
	"github.com/mohammadne/bookman/auth/internal/jwt"
	"github.com/mohammadne/bookman/auth/internal/web/rest"
	"github.com/mohammadne/bookman/auth/pkg/logger"
)

type Config struct {
	// Environment string `default:"dev"`
	Logger *logger.Config
	Jwt    *jwt.Config
	Cache  *cache.Config
	Rest   *rest.Config
}
