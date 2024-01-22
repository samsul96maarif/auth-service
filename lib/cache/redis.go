/**
 * @author [Samsul Ma'arif]
 * @email [samsulma828@gmail.com]
 * @create date 2023-11-02 20:41:39
 * @modify date 2023-11-02 20:41:39
 * @desc [description]
 */

package cache

import (
	"context"
	"time"

	"github.com/samsul96maarif/auth-service/lib/logger"

	"github.com/go-redis/redis"
)

type redisCache struct {
	*redis.Client
	defaultExpires time.Duration
}

func NewRedisCache(client *redis.Client, expires time.Duration) *redisCache {
	if expires < 0 {
		expires = 24 * time.Hour
	}
	return &redisCache{client, expires}
}

func (u *redisCache) SetToCache(key string, data interface{}, expires time.Duration) error {
	if expires < 0 {
		expires = u.defaultExpires
	}
	err := u.Set(key, data, expires).Err()
	if err != nil {
		logger.Error(context.Background(), "set to cache err: "+err.Error(), nil)
	}
	return err
}

func (u *redisCache) GetFromCache(key string) (res string, err error) {
	if res, err = u.Get(key).Result(); err != nil {
		logger.Error(context.Background(), "get from cache err: "+err.Error(), nil)
	}
	return
}
