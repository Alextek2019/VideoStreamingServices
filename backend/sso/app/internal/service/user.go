package service

import (
	"context"
	domain "vss/sso/internal/domain/user"
)

type User interface {
	Register(context.Context, domain.RegisterUserArgs) (domain.User, error)
	Update(context.Context, domain.UpdateUserArgs) error
}
