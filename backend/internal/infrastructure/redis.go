package infrastructure

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

type redisConfig struct {
	RDBHost     string `mapstructure:"RDB_HOST"`
	RDBPassword string `mapstructure:"RDB_PASSWORD"`
	RDBName     int    `mapstructure:"RDB_NAME"`
}

func connectToRedis() (*redis.Client, error) {
	config := &redisConfig{}
	_ = viper.Unmarshal(&config)

	rdb := redis.NewClient(&redis.Options{
		Addr:     config.RDBHost,
		Password: config.RDBPassword,
		DB:       config.RDBName,
	})

	err := rdb.Ping(context.Background()).Err()
	if err != nil {
		return nil, err
	}
	return rdb, nil
}
