package repository

import (
	"context"
	"database/sql"
	"grpc-app/model"
)

type UserRepository interface {
	Create(ctx context.Context, user *model.User) (*model.User, error)
	GetByID(ctx context.Context, id int64) (*model.User, error)
	GetByEmail(ctx context.Context, email string) (*model.User, error)
	Update(ctx context.Context, user *model.User) (*model.User, error)
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context, limit, offset int) ([]*model.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, user *model.User) (*model.User, error) {
	query := `
		INSERT INTO users (name, email, created_at, updated_at) 
		VALUES ($1, $2, NOW(), NOW()) 
		RETURNING id, created_at, updated_at`

	err := r.db.QueryRowContext(ctx, query, user.Name, user.Email).Scan(
		&user.ID, &user.CreatedAt, &user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) GetByID(ctx context.Context, id int64) (*model.User, error) {
	user := &model.User{}
	query := `SELECT id, name, email, created_at, updated_at FROM users WHERE id = $1`

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	user := &model.User{}
	query := `SELECT id, name, email, created_at, updated_at FROM users WHERE email = $1`

	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) Update(ctx context.Context, user *model.User) (*model.User, error) {
	query := `
		UPDATE users 
		SET name = $1, email = $2, updated_at = NOW() 
		WHERE id = $3 
		RETURNING updated_at`

	err := r.db.QueryRowContext(ctx, query, user.Name, user.Email, user.ID).Scan(&user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

func (r *userRepository) List(ctx context.Context, limit, offset int) ([]*model.User, error) {
	query := `
		SELECT id, name, email, created_at, updated_at 
		FROM users 
		ORDER BY created_at DESC 
		LIMIT $1 OFFSET $2`

	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*model.User
	for rows.Next() {
		user := &model.User{}
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
