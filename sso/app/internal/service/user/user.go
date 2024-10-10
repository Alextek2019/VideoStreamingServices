package user

import (
	"context"
	"fmt"
	"vss/sso/internal/storage"

	domain "vss/sso/internal/domain/user"
	"vss/sso/internal/service"
	"vss/sso/internal/storage/postgres/user"

	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
)

type Service struct {
	repo storage.UserRepo
}

func New(ctx context.Context) (service.User, error) {
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

func (u *Service) Register(ctx context.Context, args domain.RegisterUserArgs) error {
	if !args.Validate() {
		return fmt.Errorf("invalid login or password")
	}

	_, err := u.repo.CreateUser(ctx, storage.CreateUserDTO(args))
	if err != nil {
		return errors.Wrapf(err,
			"Service.User.RegisterUserArgs %s",
			"could not create user in database")
	}

	return nil
}

func (u *Service) Get(ctx context.Context, userID uuid.UUID) (domain.User, error) {
	userEntity, err := u.repo.GetUser(ctx, userID.String())
	if err != nil {
		return domain.User{}, errors.Wrapf(err, "could not find user with uuid: %s", userID.String())
	}

	return domain.User{
		UserID: userID,
		Login:  userEntity.Login,
	}, nil
}

func (u *Service) Update(ctx context.Context, args domain.UpdateUserArgs) error {
	err := u.repo.UpdateUser(ctx, storage.UpdateUserDTO(args))
	if err != nil {
		return errors.Wrapf(err, "could not update user with uuid %s", args.UserID)
	}

	return nil
}

func (u *Service) Delete(ctx context.Context, userID uuid.UUID) error {
	err := u.repo.DeleteUser(ctx, userID.String())
	if err != nil {
		return errors.Wrapf(err, "could not delete user with uuid %s", userID.String())
	}

	return nil
}
