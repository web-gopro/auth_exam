package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisRepoI interface {
	Exist(ctx context.Context, key string) (bool, error)
	Set(ctx context.Context, key, value string, exp int) error
	Get(ctx context.Context, key string) (string, error)
	Del(ctx context.Context, key string) (any, error)
	GetDell(ctx context.Context, key string) (string, error)
}
type redisRepo struct {
	cli *redis.Client
}

func NewRedisRepo(cli *redis.Client) RedisRepoI {

	return &redisRepo{cli}
}

func (r *redisRepo) Exist(ctx context.Context, key string) (bool, error) {

	isExists, err := r.cli.Do(ctx, "EXISTS", key).Result()

	if err != nil {
		fmt.Println("error on check exists", err.Error())
		return false, err
	}

	value, _ := isExists.(int)

	return value == 1, nil
}

func (r *redisRepo) Set(ctx context.Context, key, value string, exp int) error {

	_, err := r.cli.SetEX(ctx, key, value, time.Second*time.Duration(exp)).Result()

	if err != nil {
		fmt.Println("erro on setting to cache ", err.Error())
		return err
	}

	return nil
}

func (r *redisRepo) Get(ctx context.Context, key string) (string, error) {

	return "", nil
}
func (r *redisRepo) Del(ctx context.Context, key string) (any, error) {

	return nil, nil
}

func (r *redisRepo) GetDell(ctx context.Context, key string) (string, error) {

	anyData, err := r.cli.GetDel(ctx, key).Result()
	if err != nil {
		fmt.Println("erro on GetDel to cache ", err.Error())
		return "", err
	}
	return anyData, nil
}
