package cache

import (
	"context"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/mohammadne/bookman/auth/internal/models"
	"github.com/mohammadne/go-pkgs/failures"
	"github.com/mohammadne/go-pkgs/logger"
)

const (
	single  = 0
	cluster = 1
)

var (
	failureNotHealthy = failures.Database{}.NewInternalServer("database is not healthy")
	failureNotFound   = failures.Database{}.NewNotFound("no matching record found in database")
	failureSet        = failures.Database{}.NewInternalServer("error setting value into database")
	failureGet        = failures.Database{}.NewInternalServer("error getting value from database")
)

func NewRedis(cfg *Config, l logger.Logger) Cache {
	rc := &redisCache{config: cfg, logger: l}

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

	return rc
}

type redisCache struct {
	// passed dependencies
	config *Config
	logger logger.Logger

	// internal dependencies
	context  context.Context
	instance redis.Cmdable
}

func (rc *redisCache) IsHealthy() failures.Failure {
	err := rc.instance.Ping(rc.context).Err()
	if err != nil {
		rc.logger.Error("redis is not health", logger.Error(err))
		return failureNotHealthy
	}

	return nil
}

func (rc *redisCache) SetTokenDetail(userId uint64, td *models.TokenDetails) failures.Failure {
	if errAccess := rc.setToken(userId, td.AccessToken); errAccess != nil {
		return failureSet
	}

	if errRefresh := rc.setToken(userId, td.RefreshToken); errRefresh != nil {
		return failureSet
	}

	return nil
}

func (rc *redisCache) setToken(userId uint64, token *models.Token) failures.Failure {
	value := strconv.Itoa(int(userId))
	expire := time.Unix(token.Expires, 0).Sub(time.Now())

	err := rc.instance.Set(rc.context, token.UUID, value, expire).Err()
	if err != nil {
		rc.logger.Error("error setting token into redis", logger.Error(err))
		return failureSet
	}

	return nil
}

func (rc *redisCache) GetToken(id uint64) (*models.TokenDetails, failures.Failure) {
	// value, err := rc.instance.Get(rc.context, id).Result()
	// if err != nil {
	// 	if err == redis.Nil {
	// 		return "", failureNotFound
	// 	}

	// 	rc.logger.Error("error getting from redis", logger.Error(err))
	// 	return "", failureGet
	// }

	// if len(value) == 0 {
	// 	return "", failureNotFound
	// }

	// return value, nil
	return nil, nil
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
