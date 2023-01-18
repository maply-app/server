package cache

import (
	"fmt"
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
	"maply/config"
)

var Redis *redis.Client

func InitRedis(cfg config.RedisConfig) {
	connectionString := fmt.Sprintf("%s:%s",
		cfg.Host,
		cfg.Port,
	)
	Redis = redis.NewClient(&redis.Options{
		Addr: connectionString,
		DB:   0,
	})

	_, err := Redis.Ping().Result()
	if err != nil {
		log.Fatalf("Failed to initialize Redis: %s", err.Error())
	}
}
