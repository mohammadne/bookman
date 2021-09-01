package config

import (
	"github.com/mohammadne/bookman/library/pkg/database"
	"github.com/mohammadne/bookman/library/pkg/logger"
)

type Config struct {
	Logger   *logger.Config
	Database *database.Config
}
