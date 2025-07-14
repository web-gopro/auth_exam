package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/web-gopro/auth_exam/models"
)

type UserRepo struct {
	db  *pgx.Conn
}

func NewUserRepo(db *pgx.Conn) UserRepoI {

	return &UserRepo{db: db}
}

func (u *UserRepo) CreateUser(ctx context.Context, req models.User) (*models.UserCreateResp, error) {

	query := `
		INSERT INTO
			users (
				status,
				name,
				email,
				password,
				created_by
			)VALUES(
				$1,$2,$3,$4,$5
			)
			`

	_, err := u.db.Exec(
		ctx,
		query,
		req.Status,
		req.Name,
		req.Email,
		req.Password,
		req.CreatedBy,
	)
	if err != nil {

		fmt.Println("err on db CreateUser",err.Error())
		return nil, err
	}

	return &models.UserCreateResp{Status: "successfully inserted "}, nil

}

func (u *UserRepo) GetUserById(ctx context.Context, req models.GetById) (*models.User, error) {

	var resp models.User
	qury := `
		SELECT 
				id,
				status,
				name,
				email, 
				password,
				created_at,
				created_by
		FROM 
			users 
		WHERE
			id= $1
	`

	err := u.db.QueryRow(
		ctx,
		qury,
		req.Id,
	).Scan(
		&resp.ID,
		&resp.Status,
		&resp.Name,
		&resp.Email,
		&resp.Password,
		&resp.CreatedAt,
		&resp.CreatedBy,
	)

	if err != nil {

		fmt.Println("err on db GetUserById", err.Error())
		return nil, err
	}

	return &resp, nil

}

func (u *UserRepo) IsExists(ctx context.Context, req models.Common) (*models.CommonResp, error) {
	var isExists bool

	query := fmt.Sprintf("SELECT EXISTS (SELECT 1 FROM %s WHERE %s = '%s')", req.Table_name, req.Column_name, req.Expvalue)

	err := u.db.QueryRow(ctx, query).Scan(&isExists)

	if err != nil {
		fmt.Println("error on CheckExists", err.Error())
		return &models.CommonResp{IsExists: false}, nil
	}

	return &models.CommonResp{IsExists: isExists}, nil


}

func (u *UserRepo) UserLogin(ctx context.Context, req models.UserLogin) (*models.Claims, error) {

	// var viwerId, gmail, hashPassword, userRole string

	// query := `
	// 	SELECT
	// 		user_id,
	// 		email,
	// 		password,
	// 		user_role
	// 	FROM
	// 		users
	// 	WHERE
	// 		username =$1
	// `

	// err := u.db.QueryRow(ctx, query, req.Username).Scan(&viwerId, &gmail, &hashPassword, &userRole)

	// if err != nil {
	// 	return nil, err
	// }

	// if !helpers.CompareHashPassword(hashPassword, req.Password) {
	// 	return nil, errors.New("password is incorrect")
	// }

	// return &book_shop.Clamis{UserId: viwerId, UserRole: userRole}, nil

	return nil, nil
}
