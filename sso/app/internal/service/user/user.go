package user

import (
	"context"
	"github.com/pkg/errors"
	"vss/sso/internal/storage/postgres"
	"vss/sso/internal/storage/postgres/user"
)

type Service struct {
	repo postgres.UserRepo
}

func New(ctx context.Context) (*Service, error) {
	repo, err := user.New(ctx)
	if err != nil {
		return nil, errors.Wrapf(err,
			"service.user.New %s",
			"could not create user repository",
		)
	}

	return &Service{
		repo: repo,
	}, nil
}

func (u *Service) RegisterUser() {

}
