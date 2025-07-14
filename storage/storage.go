package storage

import (
	"github.com/jackc/pgx/v5"
	"github.com/web-gopro/auth_exam/storage/postgres"
)

type StorageI interface {
	UserRepo() postgres.UserRepoI
}

type storage struct {
	userRepo postgres.UserRepoI
}

func NewStorage(db *pgx.Conn) StorageI {

	return &storage{
		userRepo: postgres.NewUserRepo(db),
	}
}

func (s *storage) UserRepo() postgres.UserRepoI {
	return s.userRepo
}
