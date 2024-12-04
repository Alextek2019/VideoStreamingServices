package auth

import (
	"context"
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
	"time"

	domain "vss/sso/internal/domain/auth"
	"vss/sso/internal/storage"
	pgrepo "vss/sso/internal/storage/postgres/auth"
	rdrepo "vss/sso/internal/storage/redis/auth"
	domainErrors "vss/sso/pkg/errors"
)

type Provider struct {
	repo  storage.AuthRepo
	cache storage.AuthCache
}

func NewProvider(ctx context.Context) (*Provider, error) {
	repo, err := pgrepo.New(ctx)
	if err != nil {
		return nil, errors.Wrapf(err, "service.auth.NewProvider could not create auth postgres repository")
	}

	cache, err := rdrepo.New(ctx)
	if err != nil {
		return nil, errors.Wrapf(err, "service.auth.NewProvider could not create auth redis repository")
	}

	return &Provider{
		repo:  repo,
		cache: cache,
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

	err = h.cache.UpdateSessionToken(ctx, storage.SessionToken{
		UserID:      userID.String(),
		AccessToken: accessToken.String()},
	)
	if err != nil {
		return domain.SignInResponse{}, errors.Wrapf(err, "service.auth.SignIn could not update session token")
	}

	return domain.SignInResponse{
		UserID:      userID,
		TTL:         time.Now().Add(time.Hour),
		AccessToken: accessToken,
	}, nil
}

func (h *Provider) SignOut(context.Context, domain.SignOutRequest) error {

	return nil
}
