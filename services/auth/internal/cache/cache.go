package cache

import (
	"context"

	"github.com/mohammadne/bookman/auth/internal/models"
	"github.com/mohammadne/bookman/auth/pkg/failures"
)

type Cache interface {
	// IsHealthy checks correctness of service
	IsHealthy(context.Context) failures.Failure

	// SetToken sets body into cahce
	SetJwt(ctx context.Context, id uint64, body *models.Jwt) failures.Failure

	// GetUserId gets user-id
	GetUserId(context.Context, *models.AccessDetails) (uint64, failures.Failure)

	// RevokeJwt
	RevokeJwt(ctx context.Context, uuid string) (int64, failures.Failure)
}
