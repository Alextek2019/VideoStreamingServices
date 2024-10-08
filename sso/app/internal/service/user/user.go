package user

import (
	"context"
	"fmt"

	domain "vss/sso/internal/domain/user"
	"vss/sso/internal/storage/postgres"
	"vss/sso/internal/storage/postgres/user"

	"github.com/pkg/errors"
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

func (u *Service) RegisterUser(ctx context.Context, args domain.RegisterUser) error {
	if !args.Validate() {
		return fmt.Errorf("invalid login or password")
	}

	_, err := u.repo.CreateUser(ctx, postgres.CreateUserDTO(args))
	if err != nil {
		return errors.Wrapf(err,
			"Service.User.RegisterUser %s",
			"could not create user in database")
	}

	return nil
}
