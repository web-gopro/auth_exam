package handlers

import (
	"github.com/saidamir98/udevs_pkg/logger"
	"github.com/web-gopro/auth_exam/redis"
	"github.com/web-gopro/auth_exam/storage"
)

type Handler struct {
	storage  storage.StorageI
	log   logger.LoggerI
	cache redis.RedisRepoI
}

func NewHandlers( cache redis.RedisRepoI,storage storage.StorageI) Handler {

	return Handler{storage:storage , cache: cache}
}
