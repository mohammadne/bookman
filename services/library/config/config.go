package config

import (
	"github.com/mohammadne/bookman/library/pkg/database"
	"github.com/mohammadne/bookman/library/pkg/logger"
	"github.com/mohammadne/bookman/library/pkg/web/rest"
)

type Config struct {
	Logger   *logger.Config
	Database *database.Config
	Rest     *rest.Config
}
