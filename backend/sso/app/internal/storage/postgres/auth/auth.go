package auth

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"vss/sso/internal/config"
	"vss/sso/internal/storage"
	pgconnector "vss/sso/pkg/connectors/pg_connector"
	logger "vss/sso/pkg/logger/handlers/slogpretty"
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

func (r *PGRepo) ValidateUser(ctx context.Context, args storage.Credentials) (bool, storage.User, error) {
	var user storage.User
	err := r.db.GetContext(
		ctx,
		&user,
		queryValidateUser,
		args.Login,
		args.HashedPassword,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, storage.User{}, nil
		}
		return false, storage.User{}, err
	}

	return true, user, nil
}
