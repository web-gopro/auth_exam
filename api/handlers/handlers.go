package handlers

import (
	"github.com/web-gopro/auth_exam/redis"
	"github.com/web-gopro/auth_exam/storage"
)

type Handler struct {
	storage storage.StorageI
	cache   redis.RedisRepoI
}

func NewHandlers(cache redis.RedisRepoI, storage storage.StorageI) Handler {

	return Handler{storage: storage, cache: cache}
}
