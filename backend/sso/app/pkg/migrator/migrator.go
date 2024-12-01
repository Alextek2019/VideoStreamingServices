package migrator

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/pressly/goose"
	pgConnector "vss/sso/pkg/connectors/pg_connector"
)

const (
	migrationsDirPath = "migrations"
)

type Migratior struct {
	db *sql.DB
}

func New(cfg pgConnector.Postgres) (*Migratior, error) {
	db, err := goose.OpenDBWithDriver("postgres", cfg.BuildConnectionUrl())
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetConnMaxLifetime(cfg.ConnMaxLifetime)
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetConnMaxIdleTime(cfg.ConnMaxIdleTime)

	if err = db.Ping(); err != nil {
		return nil, errors.Wrapf(err,
			"pkg.migrator.New %s",
			"cannot ping postgres on connecting")
	}

	return &Migratior{
		db: db,
	}, nil
}

func (m *Migratior) Up() error {
	if err := goose.Up(m.db, migrationsDirPath); err != nil {
		return err
	}

	return nil
}

func (m *Migratior) Down() error {
	if err := goose.Down(m.db, migrationsDirPath); err != nil {
		return err
	}

	return nil
}

func (m *Migratior) UpByOne() error {
	if err := goose.UpByOne(m.db, migrationsDirPath); err != nil {
		return err
	}

	return nil
}

func (m *Migratior) Version() (int64, error) {
	ver, err := goose.GetDBVersion(m.db)
	if err != nil {
		return 0, err
	}

	return ver, nil
}

func (m *Migratior) ExecQuery(query string) error {
	if _, err := m.db.Exec(query); err != nil {
		return err
	}

	return nil
}
