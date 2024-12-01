package service

import (
	"context"
	domain "vss/sso/internal/domain/user"

	"github.com/gofrs/uuid"
)

type User interface {
	Register(context.Context, domain.RegisterUserArgs) (domain.User, error)
	Get(context.Context, uuid.UUID) (domain.User, error)
	Update(context.Context, domain.UpdateUserArgs) error
	Delete(context.Context, uuid.UUID) error
}
