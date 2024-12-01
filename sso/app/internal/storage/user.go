package storage

import (
	"context"
)

type UserRepo interface {
	CreateUser(context.Context, CreateUser) (User, error)
	GetUser(context.Context, string) (User, error)
	UpdateUser(context.Context, UpdateUser) error
	DeleteUser(context.Context, string) error
}
