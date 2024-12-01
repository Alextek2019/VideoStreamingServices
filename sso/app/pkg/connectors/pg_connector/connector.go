package pgconnector

import (
	"context"
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"go.opentelemetry.io/otel"
)

type Log interface {
	Info(string)
	Error(string)
}

func Connect(ctx context.Context, cfg Postgres, log Log) (*sqlx.DB, error) {
	conn, err := InitPsqlDB(context.Background(), cfg)
	if err != nil {
		return nil, err
	}

	go func() {
		<-ctx.Done()
		log.Info("Closing database connection")
		if err := conn.Close(); err != nil {
			log.Error(err.Error())
		} else {
			log.Info("Database connection closed properly")
		}
	}()

	return conn, nil
}

func InitPsqlDB(ctx context.Context, cfg Postgres) (*sqlx.DB, error) {
	_, span := otel.Tracer("").Start(ctx, "storage.InitPsqlDB")
	defer span.End()

	connectionURL := cfg.BuildConnectionUrl()

	database, err := sqlx.Open("pgx", connectionURL)
	if err != nil {
		return nil, errors.Wrapf(err,
			"pkg.connectors.pg_connector.InitPsqlDB %s",
			fmt.Sprintf("could not connect to postgres, connectionURL: [%s]", connectionURL))
	}

	database.SetMaxOpenConns(cfg.MaxOpenConns)
	database.SetConnMaxLifetime(cfg.ConnMaxLifetime)
	database.SetMaxIdleConns(cfg.MaxIdleConns)
	database.SetConnMaxIdleTime(cfg.ConnMaxIdleTime)

	if err = database.Ping(); err != nil {
		return nil, errors.Wrapf(err,
			"pkg.connectors.pg_connector.InitPsqlDB %s",
			"cannot ping postgres on connecting")
	}

	return database, nil
}
