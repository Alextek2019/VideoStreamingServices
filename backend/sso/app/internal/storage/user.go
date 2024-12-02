package storage

import (
	"context"
)

type UserRepo interface {
	CreateUser(context.Context, CreateUser) (User, error)
	UpdateUser(context.Context, UpdateUser) error
}
