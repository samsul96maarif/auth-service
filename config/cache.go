/*
 * Author: Samsul Ma'arif <samsulma828@gmail.com>
 * Copyright (c) 2023.
 */

package config

import (
	"os"
	"time"

	"github.com/samsul96maarif/auth-service/lib/cache"

	"github.com/go-redis/redis"
)

func NewCacheUtil(duration time.Duration) cache.CacheUtil {
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("CACHE_HOST"),
		Password: os.Getenv("CACHE_PASSWORD"),
		DB:       0,
	})
	return cache.NewRedisCache(client, duration)
}
