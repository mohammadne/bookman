package config

import (
	"github.com/mohammadne/bookman/user/internal/database"
	"github.com/mohammadne/bookman/user/internal/web/rest"
	"github.com/mohammadne/bookman/user/pkg/logger"
)

type Config struct {
	Logger   *logger.Config
	Database *database.Config
	Rest     *rest.Config
}
