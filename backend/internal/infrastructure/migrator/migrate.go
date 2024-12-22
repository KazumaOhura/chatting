package migrator

import (
	"log/slog"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func OnStart(l *slog.Logger, m *migrate.Migrate) {
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		l.Error("Failed to migration up")
	}
}

func OnStop(l *slog.Logger, m *migrate.Migrate) {
	if err := m.Down(); err != nil {
		l.Error("Failed to migration down")
	}
}

func NewMigrate(l *slog.Logger) *migrate.Migrate {

	// TODO コンフィグから読み込む
	m, err := migrate.New(
		"file://migrations",
		"postgres://postgres:postgres@database:5432/chatting?sslmode=disable",
	)

	if err != nil {
		l.Error("Failed to initialize migrate object")
		return nil
	}

	return m
}
