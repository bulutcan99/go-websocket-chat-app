package repository

import (
	"github.com/bulutcan99/go-websocket/app/model"
	custom_error "github.com/bulutcan99/go-websocket/pkg/error"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type AuthRepo struct {
	*sqlx.DB
}

type AuthInterface interface {
	CreateUser(u model.User) error
	GetUserById(id uuid.UUID) (model.User, error)
	GetUserRoleById(id uuid.UUID) (string, error)
}

func (r *AuthRepo) CreateUser(u model.User) error {
	query := `
        INSERT INTO users (
            id,
            created_at,
            updated_at,
            email,
            name_surname,
            password_hash,
            user_status,
            user_role
        ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err := r.Exec(
		query,
		u.ID, u.CreatedAt, u.UpdatedAt, u.Email, u.NameSurname, u.PasswordHash, u.Status, u.UserRole,
	)
	if err != nil {
		return custom_error.DatabaseError()
	}

	return nil
}

func (r *AuthRepo) GetUserById(id uuid.UUID) (model.User, error) {
	var user model.User
	query := `SELECT * FROM users WHERE id = $1`
	err := r.Get(&user, query, id)
	if err != nil {
		return model.User{}, custom_error.DatabaseError()
	}

	return user, nil
}

func (r *AuthRepo) GetUserRoleById(id uuid.UUID) (string, error) {
	var user model.User
	query := `SELECT email FROM users WHERE id = $1`
	err := r.Get(&user, query, id)
	if err != nil {
		return "", custom_error.DatabaseError()
	}

	return user.UserRole, nil
}
