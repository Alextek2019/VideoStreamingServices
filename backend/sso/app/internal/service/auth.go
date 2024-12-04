package service

import (
	"context"
	domain "vss/sso/internal/domain/auth"
)

type Provider interface {
	SignIn(context.Context, domain.SignInRequest) (domain.SignInResponse, error)
	SignOut(context.Context, domain.SignOutRequest) error
}

type Validator interface {
	VerifyToken(context.Context)
}
