package postgres

import (
	"context"

	"github.com/web-gopro/auth_exam/models"
)

type UserRepoI interface {
	CreateUser(ctx context.Context, req models.UserCreReq) (*models.UserCreateResp, error)
	GetUserById(ctx context.Context, req models.GetById) (*models.User, error)
	IsExists(ctx context.Context, req models.Common) (*models.CommonResp, error)
	UserLogin(ctx context.Context, req models.UserLogin) (*models.Claims, error)
}


type SysUserRepoI interface{

	CreateSysUser(ctx context.Context, req models.SysUserCretReq) (*models.SysUserCreateResp, error)
	GetSysUserById(ctx context.Context, req models.GetById) (*models.SysUser, error)
	SysUserLogin(ctx context.Context, req models.UserLogin) (*models.Claims, error)

}