package storage

import (
	"context"

	"github.com/mohammadne/bookman/user/internal/models"
	"github.com/mohammadne/bookman/user/pkg/database"
	"github.com/mohammadne/bookman/user/pkg/failures"
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
	database database.Database
}

func NewStorage(db database.Database) Storage {
	return &storage{database: db}
}
