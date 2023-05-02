package storage

import (
	"context"

	"github.com/mohammadne/bookman/user/internal/models"
	"github.com/mohammadne/bookman/user/pkg/database"
	"github.com/mohammadne/bookman/user/pkg/failures"
	"github.com/mohammadne/bookman/user/pkg/logger"
	"go.opentelemetry.io/otel/trace"
)

type Storage interface {
	CreateUser(context.Context, *models.User) failures.Failure
	FindUserById(context.Context, uint64) (*models.User, failures.Failure)
	FindUserByEmailAndPassword(context.Context, string, string) (*models.User, failures.Failure)
	FindUserByEmail(context.Context, string) (*models.User, failures.Failure)
	UpdateUser(context.Context, *models.User) failures.Failure
	DeleteUser(context.Context, *models.User) failures.Failure
}

type storage struct {
	logger   logger.Logger
	tracer   trace.Tracer
	database database.Database
}

func NewStorage(lg logger.Logger, tr trace.Tracer, db database.Database) Storage {
	return &storage{logger: lg, tracer: tr, database: db}
}
