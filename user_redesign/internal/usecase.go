package internal

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/mohammadne/bookman/user/internal/entities"
	"github.com/mohammadne/bookman/user/pkg/errors"
)

type IUserUsecase interface {
	Create(ctx context.Context, user *entities.User) (*entities.User, *errors.RestError)
	GetByID(ctx context.Context, id uuid.UUID) (*entities.User, *errors.RestError)
}
