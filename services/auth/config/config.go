package config

import (
	"github.com/mohammadne/bookman/auth/internal/cache"
	"github.com/mohammadne/bookman/auth/internal/web/rest"
	"github.com/mohammadne/bookman/auth/pkg/logger"
)

type Config struct {
	Logger *logger.Config
	Cache  *cache.Config
	Rest   *rest.Config
}
