package main

import (
	"context"
	"fmt"

	"github.com/saidamir98/udevs_pkg/logger"
	"github.com/web-gopro/auth_exam/api"
	"github.com/web-gopro/auth_exam/config"
	"github.com/web-gopro/auth_exam/pkg"
	"github.com/web-gopro/auth_exam/pkg/db"
	"github.com/web-gopro/auth_exam/redis"
	"github.com/web-gopro/auth_exam/storage"
)

func main() {

	// log := logger.NewLogger("", logger.LevelDebug)

	cfg := config.Load()

	pkgCoon, err := pkg.ConnectDB(cfg.PgConfig)

	if err != nil {

		fmt.Println("err on pkg", err.Error())
		return
	}

	fmt.Println(pkgCoon)
	redisCli, err := db.ConnRedis(context.Background(), cfg.RedisConfig)

	if err != nil {
		fmt.Println("err on redis", err.Error())
		logger.Error(err)
		return
	}

	fmt.Println(redisCli)

	cache := redis.NewRedisRepo(redisCli)

	str := storage.NewStorage(pkgCoon)

	engine := api.Api(api.Options{Cache: cache, Storage: str})

	engine.Run(":8080")
}
