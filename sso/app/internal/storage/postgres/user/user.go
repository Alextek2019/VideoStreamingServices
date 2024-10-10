package user

import (
	"context"
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
	"vss/sso/internal/config"
	"vss/sso/internal/storage"
	logger "vss/sso/pkg/logger/handlers/slogpretty"

	pgconnector "vss/sso/pkg/connectors/pg_connector"

	"github.com/jmoiron/sqlx"
)

type PGRepo struct {
	db *sqlx.DB
}

func New(ctx context.Context) (*PGRepo, error) {
	db, err := pgconnector.Connect(ctx, config.Get().Postgres, logger.Log)
	if err != nil {
		return nil, errors.Wrapf(err,
			"Storage.Postgres.User.New %s",
			"cannot create connection to postgres")
	}

	return &PGRepo{
		db: db,
	}, nil
}

func (u *PGRepo) CreateUser(ctx context.Context, args storage.CreateUser) (uuid.UUID, error) {

	return uuid.UUID{}, nil
}

func (u *PGRepo) GetUser(ctx context.Context, uuid string) (storage.User, error) {

	return storage.User{}, nil
}

func (u *PGRepo) UpdateUser(ctx context.Context, args storage.UpdateUser) error {

	return nil
}

func (u *PGRepo) DeleteUser(ctx context.Context, uuid string) error {

	return nil
}
