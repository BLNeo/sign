package redis

import "github.com/redis/go-redis/v9"

func NewIRedis() IRedis {
	return &Redis{
		rdb: GetRdb(),
	}
}

type IRedis interface {
}

type Redis struct {
	rdb *redis.Client
}
