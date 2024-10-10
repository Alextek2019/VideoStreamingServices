package storage

import (
	"context"
	"github.com/gofrs/uuid"
)

type UserRepo interface {
	CreateUser(context.Context, CreateUser) (uuid.UUID, error)
	GetUser(context.Context, string) (User, error)
	UpdateUser(context.Context, UpdateUser) error
	DeleteUser(context.Context, string) error
}
