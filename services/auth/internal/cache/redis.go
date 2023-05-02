package cache

import (
	"context"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/mohammadne/bookman/auth/internal/models"
	"github.com/mohammadne/bookman/auth/pkg/failures"
	"github.com/mohammadne/bookman/auth/pkg/logger"
	"go.opentelemetry.io/otel/trace"
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
	failureRevoke     = failures.Database{}.NewInternalServer("error revoking value from database")
)

func NewRedis(cfg *Config, lg logger.Logger, tr trace.Tracer) Cache {
	rc := &redisCache{logger: lg, tracer: tr}

	if cfg.Mode == cluster {
		rc.cmd = newClusterRedis(cfg)
	} else {
		rc.cmd = newSingleRedis(cfg)
	}

	if err := rc.cmd.Ping(context.TODO()).Err(); err != nil {
		rc.logger.Fatal(
			"error to ping redis in initialization",
			logger.String("err:", err.Error()),
		)
	}

	return rc
}

type redisCache struct {
	logger logger.Logger
	tracer trace.Tracer

	// internal dependencies
	cmd redis.Cmdable
}

func (rc *redisCache) IsHealthy(ctx context.Context) failures.Failure {
	err := rc.cmd.Ping(ctx).Err()
	if err != nil {
		rc.logger.Error("redis is not health", logger.Error(err))
		return failureNotHealthy
	}

	return nil
}

// failureUnprocessableEntity
func (rc *redisCache) SetJwt(ctx context.Context, userId uint64, td *models.Jwt) failures.Failure {
	if errAccess := rc.setToken(ctx, userId, td.AccessToken); errAccess != nil {
		return failureSet
	}

	if errRefresh := rc.setToken(ctx, userId, td.RefreshToken); errRefresh != nil {
		return failureSet
	}

	return nil
}

func (rc *redisCache) setToken(ctx context.Context, userId uint64, token *models.Token) failures.Failure {
	value := strconv.Itoa(int(userId))
	expire := time.Unix(token.Expires, 0).Sub(time.Now())

	err := rc.cmd.Set(ctx, token.UUID, value, expire).Err()
	if err != nil {
		rc.logger.Error("error setting token into redis", logger.Error(err))
		return failureSet
	}

	return nil
}

func (rc *redisCache) GetUserId(ctx context.Context, ad *models.AccessDetails) (uint64, failures.Failure) {
	userid, err := rc.cmd.Get(ctx, ad.TokenUuid).Result()
	if err != nil {
		if err == redis.Nil {
			return 0, failureNotFound
		}

		rc.logger.Error("error getting from redis", logger.Error(err))
		return 0, failureGet
	}

	userID, _ := strconv.ParseUint(userid, 10, 64)
	return userID, nil
}

// failureUnautorized
// failureUnautorized
func (rc *redisCache) RevokeJwt(ctx context.Context, uuid string) (int64, failures.Failure) {
	deleted, err := rc.cmd.Del(ctx, uuid).Result()
	if err != nil {
		rc.logger.Error("error revoking from redis", logger.Error(err))
		return 0, failureRevoke
	}

	return deleted, nil
}

// newSingleRedis returns a new `RedisHandler` with a single Redis client.
func newSingleRedis(cfg *Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:               cfg.URL,
		Password:           cfg.Password,
		MaxRetries:         cfg.MaxRetries,
		MinRetryBackoff:    cfg.MinRetryBackoff,
		MaxRetryBackoff:    cfg.MaxRetryBackoff,
		ReadTimeout:        cfg.ReadTimeout,
		PoolSize:           cfg.PoolSize,
		PoolTimeout:        cfg.PoolTimeout,
		IdleTimeout:        cfg.IdleTimeout,
		IdleCheckFrequency: cfg.IdleCheckFrequency,
	})
}

// newClusterRedis returns a new `RedisHandler` with a clustered Redis client.
func newClusterRedis(cfg *Config) *redis.ClusterClient {
	return redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:              []string{cfg.MasterURL, cfg.SlaveURL},
		Password:           cfg.Password,
		MaxRetries:         cfg.MaxRetries,
		MinRetryBackoff:    cfg.MinRetryBackoff,
		MaxRetryBackoff:    cfg.MaxRetryBackoff,
		ReadTimeout:        cfg.ReadTimeout,
		PoolSize:           cfg.PoolSize,
		PoolTimeout:        cfg.PoolTimeout,
		IdleTimeout:        cfg.IdleTimeout,
		IdleCheckFrequency: cfg.IdleCheckFrequency,
	})
}
