package redis

import (
	"ads.cost.com/config"
	"ads.cost.com/logger"
	"github.com/go-redis/redis"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func InitRedis() error {
	conf := config.GetConfig()
	RedisClient := redis.NewClient(&redis.Options{
		Addr:		conf.RedisConfig.Addr,
		Password: 	conf.RedisConfig.Password,
		DB: 		conf.RedisConfig.DB,
	})

	pong, err := RedisClient.Ping().Result()
	if err != nil {
		err = errors.Wrap(err, "InitRedis")
		return err
	}
	logger.GetLogger().Info("Redis ping:", zap.String("ping", pong))
	return nil
}