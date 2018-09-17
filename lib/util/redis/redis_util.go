package redis

import (
	"github.com/go-redis/redis"
	"time"
)

type Config struct {
	Host string
	Port string
}

type redis_client struct {
	conf Config
	redisdb *redis.Client
}

var Rediscli redis_client

func New(conf Config) *redis_client {
	if Rediscli.conf.Host != conf.Host || Rediscli.conf.Port != conf.Port {
		Rediscli = redis_client{conf, nil}
		Rediscli.redisdb = redis.NewClient(&redis.Options{
			Addr:        conf.Host + ":" + conf.Port,
			DialTimeout: 3 * time.Second,
			ReadTimeout: 3 * time.Second,
		})
	}
	return &Rediscli
}

func (rediscli redis_client) GetHashVal(key string) map[string]string {

	group := "repo:"
	data_map := rediscli.redisdb.HGetAll( group + key).Val()


	return data_map
}

