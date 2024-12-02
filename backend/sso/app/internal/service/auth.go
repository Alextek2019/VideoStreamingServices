package service

import "context"

type Auth interface {
	SignIn(context.Context)
	SignOut(context.Context)
	VerifyToken(context.Context)
}
