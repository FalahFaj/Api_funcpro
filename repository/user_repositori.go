package repository

import (
	"context"
	"database/sql"
	"projek_funcpro_kel12/model"
)

type UserRepository interface {
	Buat(ctx context.Context,user *model.User) (int64, error)
	GetUserById(ctx context.Context, id int64) (*model.User, error)
	GetUserByEmail(ctx context.Context, mail string) (*model.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) Buat(ctx context.Context,user *model.User) (int64, error) {
	var id int64
	query := `INSERT INTO users (nama, email, password, role) VALUES ($1, $2, $3, $4) RETURNING id`

	err := r.db.QueryRow(query, user.Nama, user.Email, user.Password, user.Role).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *userRepository) GetUserById(ctx context.Context, id int64) (*model.User, error) {
	var user model.User

	query := `SELECT id, nama, email, password, role, created_at FROM users WHERE id = $1`

	err := r.db.QueryRow(query, id).Scan(&user.Id, &user.Nama, &user.Email, &user.Password, &user.Role, &user.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) GetUserByEmailctx(ctx context.Context,email string) (*model.User, error) {
	var user model.User

	query := `SELECT id, nama, email, password, role, created_at FROM users WHERE email = $1`

	err := r.db.QueryRow(query, email).Scan(&user.Id, &user.Nama, &user.Email, &user.Password, &user.Role, &user.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
