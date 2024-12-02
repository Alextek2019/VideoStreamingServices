package user

import (
	"context"
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

func (u *PGRepo) CreateUser(ctx context.Context, args storage.CreateUser) (storage.User, error) {
	var user storage.User
	err := u.db.GetContext(
		ctx,
		&user,
		queryCreateUser,
		args.Login,
		args.HashedPassword,
	)
	if err != nil {
		return user, errors.Wrapf(err, "cannot create user")
	}

	return user, nil
}

func (u *PGRepo) UpdateUser(ctx context.Context, args storage.UpdateUser) error {

	return nil
}
