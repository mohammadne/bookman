package config

import (
	"github.com/mohammadne/bookman/user/internal/database"
	grpc_server "github.com/mohammadne/bookman/user/internal/network/grpc/server"
	"github.com/mohammadne/bookman/user/internal/network/rest"
	"github.com/mohammadne/bookman/user/pkg/logger"
)

type Config struct {
	Logger   *logger.Config
	Database *database.Config
	Rest     *rest.Config
	Grpc     *grpc_server.Config
}
