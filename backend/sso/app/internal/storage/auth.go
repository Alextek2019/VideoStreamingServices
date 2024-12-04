package storage

import (
	"context"
	"time"
	domain "vss/sso/internal/domain/auth"
)

type AuthRepo interface {
	ValidateUser(context.Context, Credentials) (bool, User, error)
}

type AuthCache interface {
	UpdateSessionToken(context.Context, SessionToken) error
}

type Credentials struct {
	Login          string `db:"login"`
	HashedPassword string `db:"password"`
}

type SessionToken struct {
	UserID      string
	AccessToken string
	TTL         time.Duration
}

func ValidateUserDTO(in domain.SignInRequest) Credentials {
	return Credentials{
		Login:          in.Login,
		HashedPassword: in.Password,
	}
}
