package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/web-gopro/auth_exam/models"
	"github.com/web-gopro/auth_exam/pkg/helpers"
)

type SysUserRepo struct {
	db *pgx.Conn
}

func NewSysUserRepo(db *pgx.Conn) SysUserRepoI {

	return &SysUserRepo{
		db: db,
	}
}

func (p *SysUserRepo) CreateSysUser(ctx context.Context, req models.SysUserCretReq, createdBy string) (*models.SysUserCreateResp, error) {

	id := uuid.New()
	var roleId string
	query := `
		INSERT INTO
			sysusers (
				id,
				status,
				email,
				name,
				password,
				created_by
			)VALUES(
				$1,$2,$3,$4,$5,$6
			)
			`

	_, err := p.db.Exec(
		ctx,
		query,
		id,
		req.Status,
		req.Email,
		req.Name,
		req.Password,
		createdBy,
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
func (p *SysUserRepo) GetSysUser(ctx context.Context, req models.GetById) (*models.SysUserGetResp, error) {

	resp := models.SysUserGetResp{}
	query := `
		SELECT
   			su.id,
   			su.name,
    		su.email,
   			su.status,
   			su.created_at,
  			su.created_by,
    		r.name AS role_name
		FROM
    		sysusers su
		LEFT JOIN
   			sysuser_roles sr ON su.id = sr.sysuser_id
		LEFT JOIN
   			roles r ON sr.role_id = r.id
		WHERE
   			 su.id = $1
   		AND su.status = 'active'
    	AND (r.status = 'active' OR r.status IS NULL);
	`

	err := p.db.QueryRow(
		ctx, query,
		req.Id,
	).Scan(
		&resp.ID,
		&resp.Name,
		&resp.Email,
		&resp.Status,
		&resp.CreatedAt,
		&resp.CreatedBy,
		&resp.Role,
	)

	if err != nil {

		fmt.Println("err on db get sysuser ", err.Error())
		return nil, err
	}
	return &resp, nil
}
func (p *SysUserRepo) SysUserLogin(ctx context.Context, req models.LoginReq) (*models.Claims, error) {

	var hashPassword string
	var resp models.Claims
	query := `
		SELECT
			su.id,
			su.password,
   		 	r.name AS role_name
		FROM
    		sysusers su
		JOIN
   			 sysuser_roles sr ON su.id = sr.sysuser_id
		JOIN
    		roles r ON sr.role_id = r.id
		WHERE
    	LOWER(su.email) = LOWER($1)
   		AND su.status = 'active'
    	AND r.status = 'active';
`

	err := p.db.QueryRow(
		ctx,
		query,
		req.Email,
	).Scan(
		&resp.User_id,
		&hashPassword,
		&resp.User_role,
	)

	if err != nil {

		fmt.Println("err on db sysuser login ", err.Error())
		return nil, err
	}

	if !helpers.CompareHashPassword(hashPassword, req.User_password) {
		return nil, errors.New("password is incorrect")
	}

	return &resp, nil
}
