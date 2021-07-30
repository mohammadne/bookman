package internal

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/mohammadne/bookman/user/internal/entities"
)

type IUserRepository interface {
	Create(ctx context.Context, user *entities.User) (*entities.User, error)
	GetByID(ctx context.Context, id uuid.UUID) (*entities.User, error)
}
