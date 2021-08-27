package cache

import (
	"github.com/mohammadne/bookman/auth/internal/models"
	"github.com/mohammadne/go-pkgs/failures"
)

type Cache interface {
	// IsHealthy checks correctness of service
	IsHealthy() failures.Failure

	// SetToken sets body into cahce
	SetJwt(id uint64, body *models.Jwt) failures.Failure

	// GetToken gets id-value and put it into body
	GetToken(id uint64) (*models.Jwt, failures.Failure)
}
