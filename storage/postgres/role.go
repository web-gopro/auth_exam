package postgres

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/web-gopro/auth_exam/models"
)

type RoleRepo struct {
	db *pgx.Conn
}

func NewRoleRepo(db *pgx.Conn) RoleRepoI {

	return &RoleRepo{db: db}

}

func (r *RoleRepo) Create(ctx context.Context, req *models.CreateRoleRequest, createrBy string) (*models.Role, error) {

	id := uuid.New()
	query := `
		INSERT INTO
			roles(
				id,
				status,
				name,
				created_by
			)VALUES(
				$1,$2,$3,$4
			)

	`

	_, err := r.db.Exec(
		ctx,
		query,
		id,
		req.Status,
		req.Name,
		createrBy,
	)

	if err != nil {

		fmt.Println("err on db CreateRole", err.Error())
		return nil, err
	}

	resp, err := r.GetByID(context.Background(), models.GetById{Id: id.String()})

	if err != nil {

		fmt.Println("err on db getRole", err.Error())
		return nil, err
	}

	return resp, nil
}
func (r *RoleRepo) GetByID(ctx context.Context, req models.GetById) (*models.Role, error) {

	resp := models.Role{}

	query := `
		SELECT 	
			id,
			status,
			name,
			created_at,
			created_by
		FROM
			roles
		WHERE 
			id=$1
	`

	err := r.db.QueryRow(
		ctx,
		query,
		req.Id,
	).Scan(

		&resp.ID,
		&resp.Status,
		&resp.Name,
		&resp.CreatedAt,
		&resp.CreatedBy,
	)

	if err != nil {

		fmt.Println("err on db GetRole", err.Error())
		return nil, err
	}
	return &resp, nil
}
func (r *RoleRepo) GetByName(ctx context.Context, req string) (*models.Role, error) {

	return nil, nil
}
func (r *RoleRepo) Update(ctx context.Context, updates *models.UpdateRoleRequest) (*models.Role, error) {

	query := `
        UPDATE 
			roles
        SET
            status = $1,
            name = $2,
        WHERE 
			id = $3
        RETURNING
            id,
            status,
            name,
            created_at,
            created_by
    `

	var resp models.Role

	err := r.db.QueryRow(
		ctx,
		query,
		updates.Status,
		updates.Name,
		updates.ID,
	).Scan(
		&resp.ID,
		&resp.Status,
		&resp.Name,
		&resp.CreatedAt,
		&resp.CreatedBy,
	)

	if err != nil {
		fmt.Println("err on db UpdateRole:", err.Error())
		return nil, err
	}

	return &resp, nil
}
func (r *RoleRepo) List(ctx context.Context, req *models.GetList) ([]*models.Role, error) {

	return nil, nil
}
