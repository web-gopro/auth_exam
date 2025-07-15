package postgres

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/web-gopro/auth_exam/models"
)

type SysUserRepo struct {
	db *pgx.Conn
}

func NewSysUserRepo(db *pgx.Conn) SysUserRepoI {

	return &SysUserRepo{
		db: db,
	}
}

func (p *SysUserRepo) CreateSysUser(ctx context.Context, req models.SysUserCretReq) (*models.SysUserCreateResp, error) {

	id := uuid.New()
	var roleId string
	query := `
		INSERT INTO
			sysusers (
				id,
				status,
				name,
				password,
				created_by
			)VALUES(
				$1,$2,$3,$4,$5
			)
			`

	_, err := p.db.Exec(
		ctx,
		query,
		id,
		req.Status,
		req.Name,
		req.Password,
		req.CreatedBy,
	)

	if err != nil {

		fmt.Println("err on db CreateSysUser", err.Error())
		return nil, err
	}

	roleQuery := `
		SELECT 
			id
		FROM 
			roles
		WHERE LOWER(name) = LOWER($1) AND status = 'active';


	`
	err = p.db.QueryRow(
		ctx,
		roleQuery,
		req.Role,
	).Scan(&roleId)

	if err != nil {

		fmt.Println("err on db get role id", err.Error())
		return nil, err
	}

	sysRoleQuery := `
		INSERT INTO 
			sysuser_roles (
				sysuser_id, 
				role_id
				)VALUES(
				$1,$2
				)`

	_, err = p.db.Exec(ctx, sysRoleQuery, id, roleId)

	if err != nil {

		fmt.Println("err on db create SySrole id", err.Error())
		return nil, err
	}

	return &models.SysUserCreateResp{Id: id.String(), Role: req.Role}, nil
}
func (p *SysUserRepo) GetSysUserById(ctx context.Context, req models.GetById) (*models.SysUser, error) {

	return nil, nil
}
func (p *SysUserRepo) SysUserLogin(ctx context.Context, req models.UserLogin) (*models.Claims, error) {

	return nil, nil
}
