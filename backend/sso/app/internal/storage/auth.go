package storage

import (
	"context"
	domain "vss/sso/internal/domain/auth"
)

type Auth interface {
	ValidateUser(context.Context, Credentials) (bool, User, error)
}

type Credentials struct {
	Login          string `db:"login"`
	HashedPassword string `db:"password"`
}

func ValidateUserDTO(in domain.SignInRequest) Credentials {
	return Credentials{
		Login:          in.Login,
		HashedPassword: in.Password,
	}
}
