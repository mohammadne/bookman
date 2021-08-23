package cache

import "github.com/mohammadne/bookman/core/failures"

type Cache interface {
	// Initialize is cache-service setup
	Initialize()

	// IsHealthy checks correctness of service
	IsHealthy() failures.Failure

	// Get gets id-value and put it into body
	Get(id string) (string, failures.Failure)

	// Save saves body into cahce
	Set(id string, body string) failures.Failure
}
