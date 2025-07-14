package db

import (
	"context"
	"fmt"

	"strconv"

	"github.com/go-redis/redis/v8"
	"github.com/web-gopro/auth_exam/config"
)

func RedisAdr(host string, port int) string {

	return host + ":" + strconv.Itoa(port)
}

func ConnRedis(ctx context.Context, cfg config.RedisConfig) (*redis.Client, error) {

	redisCli := redis.NewClient(&redis.Options{
		Addr: RedisAdr(cfg.Host, cfg.Port),
		DB:   cfg.DBIndex,
	})

	_, err := redisCli.Ping(ctx).Result()
	if err != nil {

		fmt.Println("err on con redis" ,err)
		return nil, err
	}

	fmt.Println("sucesfully conected with redis")
	return redisCli, nil
}
