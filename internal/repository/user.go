package repository

import (
	"context"

	"weather_api/internal/models"

	"github.com/jmoiron/sqlx"
)

type UserRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r *UserRepo) GetUsers(ctx context.Context) ([]models.User, error) {
	query := `
		SELECT id, name, email, created_at, updated_at, deleted_at
		FROM users
		WHERE deleted_at IS NULL
		ORDER BY id;
	`

	var users []models.User
	if err := r.db.SelectContext(ctx, &users, query); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepo) GetUserByID(ctx context.Context, id int) (models.User, error) {
	query := `
		SELECT id, name, email, created_at, updated_at, deleted_at
		FROM users
		WHERE id = $1 AND deleted_at IS NULL;
	`

	var user models.User
	if err := r.db.GetContext(ctx, &user, query, id); err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (r *UserRepo) CreateUser(ctx context.Context, user models.User) (models.User, error) {
	query := `
		INSERT INTO users (name, email)
		VALUES ($1, $2)
		RETURNING id, name, email, created_at, updated_at, deleted_at;
	`

	var created models.User
	if err := r.db.GetContext(ctx, &created, query, user.Name, user.Email); err != nil {
		return models.User{}, err
	}

	return created, nil
}

func (r *UserRepo) UpdateUser(ctx context.Context, id int, user models.User) (models.User, error) {
	query := `
		UPDATE users
		SET name = $1,
		    email = $2,
		    updated_at = now()
		WHERE id = $3 AND deleted_at IS NULL
		RETURNING id, name, email, created_at, updated_at, deleted_at;
	`

	var updated models.User
	if err := r.db.GetContext(ctx, &updated, query, user.Name, user.Email, id); err != nil {
		return models.User{}, err
	}

	return updated, nil
}

func (r *UserRepo) DeleteUser(ctx context.Context, id int) error {
	query := `
		UPDATE users
		SET deleted_at = now(),
		    updated_at = now()
		WHERE id = $1 AND deleted_at IS NULL;
	`

	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
