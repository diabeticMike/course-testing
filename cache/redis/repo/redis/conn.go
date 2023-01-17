package redis

import (
	"github.com/go-redis/redis"
)

const (
	redisAddr   = "localhost:6379"
	redisPasswd = ""
	redisDB     = 0
)

// NewConn connects to redis and return redis client
func NewConn() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPasswd,
		DB:       redisDB,
	})
	if _, err := client.Ping().Result(); err != nil {
		return nil, err
	}
	return client, nil
}
