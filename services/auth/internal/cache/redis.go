package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/mohammadne/bookman/core/failures"
	"github.com/mohammadne/bookman/core/logger"
)

const (
	single  = 0
	cluster = 1
)

var (
	failureNotHealthy = failures.Database{}.NewInternalServer("database is not healthy")
	failureNotFound   = failures.Database{}.NewNotFound("no matching record found in database")
	failureGet        = failures.Database{}.NewInternalServer("error getting value from database")
	failureSet        = failures.Database{}.NewInternalServer("error setting value into database")
)

func New(cfg *Config, l logger.Logger) Cache {
	return &redisCache{config: cfg, logger: l}
}

type redisCache struct {
	// passed dependencies
	config *Config
	logger logger.Logger

	// internal dependencies
	context  context.Context
	instance redis.Cmdable
}

func (rc *redisCache) Initialize() {
	rc.context = context.TODO()

	if rc.config.Mode == cluster {
		rc.instance = rc.newClusterRedis()
	} else {
		rc.instance = rc.newSingleRedis()
	}

	if err := rc.instance.Ping(rc.context).Err(); err != nil {
		rc.logger.Fatal(
			"error to ping redis in initialization",
			logger.String("err:", err.Error()),
		)
	}
}

func (rc *redisCache) IsHealthy() failures.Failure {
	err := rc.instance.Ping(rc.context).Err()
	if err != nil {
		rc.logger.Error("redis is not health", logger.Error(err))
		return failureNotHealthy
	}

	return nil
}

func (rc *redisCache) Get(id string) (string, failures.Failure) {
	value, err := rc.instance.Get(rc.context, id).Result()
	if err != nil {
		if err == redis.Nil {
			return "", failureNotFound
		}

		rc.logger.Error("error getting from redis", logger.Error(err))
		return "", failureGet
	}

	if len(value) == 0 {
		return "", failureNotFound
	}

	return value, nil
}

func (rc *redisCache) Set(id string, body string) failures.Failure {
	expirationTime := time.Second * time.Duration(rc.config.ExpirationTime)
	err := rc.instance.Set(rc.context, id, body, expirationTime).Err()
	if err != nil {
		rc.logger.Error("error setting into redis", logger.Error(err))
		return failureSet
	}

	return nil
}

// newSingleRedis returns a new `RedisHandler` with a single Redis client.
func (rc *redisCache) newSingleRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:               rc.config.URL,
		Password:           rc.config.Password,
		MaxRetries:         rc.config.MaxRetries,
		MinRetryBackoff:    rc.config.MinRetryBackoff,
		MaxRetryBackoff:    rc.config.MaxRetryBackoff,
		ReadTimeout:        rc.config.ReadTimeout,
		PoolSize:           rc.config.PoolSize,
		PoolTimeout:        rc.config.PoolTimeout,
		IdleTimeout:        rc.config.IdleTimeout,
		IdleCheckFrequency: rc.config.IdleCheckFrequency,
	})
}

// newClusterRedis returns a new `RedisHandler` with a clustered Redis client.
func (rc *redisCache) newClusterRedis() *redis.ClusterClient {
	return redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:              []string{rc.config.MasterURL, rc.config.SlaveURL},
		Password:           rc.config.Password,
		MaxRetries:         rc.config.MaxRetries,
		MinRetryBackoff:    rc.config.MinRetryBackoff,
		MaxRetryBackoff:    rc.config.MaxRetryBackoff,
		ReadTimeout:        rc.config.ReadTimeout,
		PoolSize:           rc.config.PoolSize,
		PoolTimeout:        rc.config.PoolTimeout,
		IdleTimeout:        rc.config.IdleTimeout,
		IdleCheckFrequency: rc.config.IdleCheckFrequency,
	})
}
