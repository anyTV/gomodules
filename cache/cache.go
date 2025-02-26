package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/anyTV/gomodules/ftime"
	logger "github.com/anyTV/gomodules/logging"
	"github.com/redis/go-redis/v9"
)

var cacheCtx = context.Background()

var log = logger.New("cache.go")

// ----------------------------------------------------------------------------

type Cache struct {
	redis *redis.Client
}

func (c *Cache) Delete(key string) {
	delRes, err := c.redis.Del(cacheCtx, key).Result()

	if err != nil {
		// We'll ignore the error but log it
		log.Warnf(
			"Failed to delete key: %s. Code: %d. Error: %s",
			key, delRes, err.Error(),
		)
	}
}

// Primitive caching
// To add additional commands, go to https://github.com/redis/go-redis/blob/e63669e1706936ac794277340c51a51c5facca70/command.go#L840
// and create the corresponding methods

func (c *Cache) GetVal(key string) (string, error) {
	return c.redis.Get(cacheCtx, key).Result()
}

func (c *Cache) GetBytes(key string) ([]byte, error) {
	return c.redis.Get(cacheCtx, key).Bytes()
}

func (c *Cache) GetUint64(key string) (uint64, error) {
	return c.redis.Get(cacheCtx, key).Uint64()
}

func (c *Cache) Set(key string, val any) error {
	return c.SetTtl(key, val, ftime.Zero)
}

func (c *Cache) SetTtl(key string, value any, duration time.Duration) error {
	err := c.redis.Set(cacheCtx, key, value, duration).Err()

	if err != nil {
		return fmt.Errorf("failed to cache `%s` with error: %s", key, err.Error())
	}

	return nil
}

// GetStringMapString ----------------------------------------------------------------------------
// Map string string
func (c *Cache) GetStringMapString(key string) (map[string]string, error) {
	res, err := c.redis.HGetAll(cacheCtx, key).Result()

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Cache) SetStringMapString(key string, m map[string]string) error {
	return c.SetStringMapStringTtl(key, m, ftime.Zero)
}

func (c *Cache) SetStringMapStringTtl(key string, m map[string]string, duration time.Duration) error {
	if err := c.redis.HSet(cacheCtx, key, m).Err(); err != nil {
		return err
	}

	if err := c.redis.Expire(cacheCtx, key, duration).Err(); err != nil {
		return err
	}

	return nil
}

type CacheOptions struct {
	host string
	port string
	db   int
	pass string
}

func New(o *CacheOptions) *Cache {
	log.Infof("Initializing caching connection...")

	redisOpts := redis.Options{
		Addr:     fmt.Sprintf("%s:%s", o.host, o.port),
		Password: o.pass,
		DB:       o.db,
	}

	log.Infof("Connection made to %s:%s at DB: %d", o.host, o.port, o.db)

	return &Cache{redis.NewClient(&redisOpts)}
}
