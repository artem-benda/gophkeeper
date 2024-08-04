package db

import (
	"context"
	"errors"
	gophkeeper "github.com/artem-benda/gophkeeper/server"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/jackc/pgx/v5/pgxpool"
)

func MustRunDBMigrations(dbURL string) {
	d, err := iofs.New(gophkeeper.FS, "db/migrations")
	if err != nil {
		panic(err)
	}
	m, err := migrate.NewWithSourceInstance("iofs", d, dbURL)
	if err != nil {
		panic(err)
	}
	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		panic(err)
	}
}

func MustCreateConnectionPool(databaseDSN string) *pgxpool.Pool {
	dbPool, err := pgxpool.New(context.Background(), databaseDSN)
	if err != nil {
		panic(err)
	}
	return dbPool
}
