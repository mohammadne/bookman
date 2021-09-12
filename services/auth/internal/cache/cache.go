package cache

import (
	"github.com/mohammadne/bookman/auth/internal/models"
	"github.com/mohammadne/bookman/auth/pkg/failures"
)

type Cache interface {
	// IsHealthy checks correctness of service
	IsHealthy() failures.Failure

	// SetToken sets body into cahce
	SetJwt(id uint64, body *models.Jwt) failures.Failure

	// GetUserId gets user-id
	GetUserId(*models.AccessDetails) (uint64, failures.Failure)

	RevokeJwt(uuid string) (int64, failures.Failure)
}
