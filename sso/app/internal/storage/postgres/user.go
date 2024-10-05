package postgres

import "context"

type UserRepo interface {
	CreateUser(context.Context, CreateUser) (int, error)
	GetUser(context.Context, string) (User, error)
	UpdateUser(context.Context, UpdateUser) error
	DeleteUser(context.Context, string) error
}
