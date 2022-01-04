package configuration

import "github.com/go-redis/redis"

func ConfigCache(addressUrl string, password string) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:         addressUrl,
		Password:     password,
		DB:           0,
		PoolSize:     50,
		MinIdleConns: 10,
	})
	return rdb
}
