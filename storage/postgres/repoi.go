package postgres

import (
	"context"

	"github.com/web-gopro/auth_exam/models"
)

type UserRepoI interface {
	CreateUser(ctx context.Context, req models.UserCreReq) (*models.UserCreateResp, error)
	GetUserById(ctx context.Context, req models.GetById) (*models.User, error)
	IsExists(ctx context.Context, req models.Common) (*models.CommonResp, error)
	UserLogin(ctx context.Context, req models.LoginReq) (*models.Claims, error)
}

type SysUserRepoI interface {
	CreateSysUser(ctx context.Context, req models.SysUserCretReq, createdBy string) (*models.SysUserCreateResp, error)
	GetSysUser(ctx context.Context, req models.GetById) (*models.SysUserGetResp, error)
	SysUserLogin(ctx context.Context, req models.LoginReq) (*models.Claims, error)
}

type RoleRepoI interface {
	Create(ctx context.Context, role *models.CreateRoleRequest, createrBy string) (*models.Role, error)
	GetByID(ctx context.Context, req models.GetById) (*models.Role, error)
	GetByName(ctx context.Context, name string) (*models.Role, error)
	Update(ctx context.Context, updates *models.UpdateRoleRequest) (*models.Role, error)
	List(ctx context.Context, filter *models.GetList) ([]*models.Role, error)
}
