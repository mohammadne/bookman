package repositories

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/mohammadne/bookman/user/internal"
	"github.com/mohammadne/bookman/user/internal/entities"
)

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) internal.IUserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, user *entities.User) (*entities.User, error) {
	return nil, nil
}

func (r *userRepository) GetByID(ctx context.Context, id uuid.UUID) (*entities.User, error) {
	return nil, nil
}
