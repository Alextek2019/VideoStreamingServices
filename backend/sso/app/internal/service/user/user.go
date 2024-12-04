package user

import (
	"context"
	"fmt"
	"vss/sso/internal/storage"

	domain "vss/sso/internal/domain/user"
	"vss/sso/internal/service"
	"vss/sso/internal/storage/postgres/user"

	"github.com/pkg/errors"
)

type Service struct {
	repo storage.UserRepo
}

func New(ctx context.Context) (service.User, error) {
	repo, err := user.New(ctx)
	if err != nil {
		return nil, errors.Wrapf(err, "could not create user repository")
	}

	return &Service{
		repo: repo,
	}, nil
}

func (u *Service) Register(ctx context.Context, args domain.RegisterUserArgs) (domain.User, error) {
	if !args.Validate() {
		return domain.User{}, fmt.Errorf("invalid login or password")
	}

	response, err := u.repo.CreateUser(ctx, storage.CreateUserDTO(args))
	if err != nil {
		return domain.User{}, errors.Wrapf(err, "could not create user in database")
	}

	return UserDTO(response), nil
}

func (u *Service) Update(ctx context.Context, args domain.UpdateUserArgs) error {
	err := u.repo.UpdateUser(ctx, storage.UpdateUserDTO(args))
	if err != nil {
		return errors.Wrapf(err, "could not update user with uuid %s", args.UserID)
	}

	return nil
}
