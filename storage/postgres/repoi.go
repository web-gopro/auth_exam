package postgres

import (
	"context"

	"github.com/web-gopro/auth_exam/models"
)

type UserRepoI interface {
	CreateUser(ctx context.Context, req models.User) (*models.UserCreateResp, error)
	GetUserById(ctx context.Context, req models.GetById) (*models.User, error)
	IsExists(ctx context.Context, req models.Common) (*models.CommonResp, error)
	UserLogin(ctx context.Context, req models.UserLogin) (*models.Claims, error)
}
