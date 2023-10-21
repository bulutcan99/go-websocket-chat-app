package repository

//
// import (
// 	"context"
// 	"github.com/bulutcan99/go-websocket/model"
// 	"github.com/bulutcan99/go-websocket/pkg/config"
// 	custom_error "github.com/bulutcan99/go-websocket/pkg/error"
// 	"github.com/bulutcan99/go-websocket/pkg/utility"
// 	"github.com/google/uuid"
// 	"github.com/jmoiron/sqlx"
// )
//
// type UserInterface interface {
// 	GetShownedUserByEmail(email string) (model.UserShown, error)
// 	ChangePassword(id uuid.UUID, oldPassword string, newPassword string) error
// }
//
// type UserRepo struct {
// 	db      *sqlx.DB
// 	context context.Context
// }
//
// func NewUserRepo(psql *config.PostgreSQL) *UserRepo {
// 	return &UserRepo{
// 		db:      psql.DB,
// 		context: psql.Context,
// 	}
// }
//
// func (r *UserRepo) GetShownedUserByEmail(email string) (model.UserShown, error) {
// 	var user model.UserShown
// 	query := `SELECT * FROM users WHERE email = $1`
// 	err := r.Get(&user, query, email)
// 	if err != nil {
// 		return model.UserShown{}, custom_error.DatabaseError()
// 	}
//
// 	return user, nil
// }
//
// func (r *UserRepo) ChangePassword(id uuid.UUID, oldPassword string, newPassword string) error {
// 	var user model.User
// 	query := `SELECT * FROM users WHERE id = $1`
// 	err := r.Get(&user, query, id)
// 	if err != nil {
// 		return custom_error.DatabaseError()
// 	}
//
// 	if !utility.ComparePasswords(user.PasswordHash, oldPassword) {
// 		return custom_error.ValidationError()
// 	}
//
// 	updateQuery := `UPDATE users SET password_hash = $1 WHERE id = $2`
// 	hashedPassword := utility.GeneratePassword(newPassword)
// 	_, err = r.Exec(updateQuery, hashedPassword, id)
// 	if err != nil {
// 		return custom_error.DatabaseError()
// 	}
//
// 	return nil
// }
