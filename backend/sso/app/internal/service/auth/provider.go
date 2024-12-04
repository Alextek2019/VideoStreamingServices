package auth

import (
	"context"
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
	"time"

	domain "vss/sso/internal/domain/auth"
	"vss/sso/internal/storage"
	pgrepo "vss/sso/internal/storage/postgres/auth"
	domainErrors "vss/sso/pkg/errors"
)

type Provider struct {
	repo storage.Auth
}

func NewProvider(ctx context.Context) (*Provider, error) {
	repo, err := pgrepo.New(ctx)
	if err != nil {
		return nil, errors.Wrapf(err, "service.auth.NewProvider could not create user repository")
	}

	return &Provider{
		repo: repo,
	}, nil
}

func (h *Provider) SignIn(ctx context.Context, args domain.SignInRequest) (domain.SignInResponse, error) {
	isValid, user, err := h.repo.ValidateUser(ctx, storage.ValidateUserDTO(args))
	if err != nil {
		return domain.SignInResponse{}, errors.Wrapf(err, "service.auth.SignIn could not validate user credentials")
	} else if !isValid {
		return domain.SignInResponse{}, domainErrors.ErrInvalidAuth
	}

	userID, _ := uuid.FromString(user.ID)
	accessToken, _ := uuid.DefaultGenerator.NewV4()

	return domain.SignInResponse{
		UserID:      userID,
		TTL:         time.Now().Add(time.Hour),
		AccessToken: accessToken,
	}, nil
}

func (h *Provider) SignOut(context.Context, domain.SignOutRequest) error {

	return nil
}
