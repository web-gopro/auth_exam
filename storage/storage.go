package storage

import (
	"github.com/jackc/pgx/v5"
	"github.com/web-gopro/auth_exam/storage/postgres"
)

type StorageI interface {
	UserRepo() postgres.UserRepoI
	SysUserRepo() postgres.SysUserRepoI
}

type storage struct {
	userRepo    postgres.UserRepoI
	sysUserRepo postgres.SysUserRepoI
}

func NewStorage(db *pgx.Conn) StorageI {

	return &storage{
		userRepo: postgres.NewUserRepo(db),
		sysUserRepo: postgres.NewSysUserRepo(db),
	}
}

func (s *storage) UserRepo() postgres.UserRepoI {
	return s.userRepo
}

func (s *storage) SysUserRepo() postgres.SysUserRepoI{

	return s.sysUserRepo
}
