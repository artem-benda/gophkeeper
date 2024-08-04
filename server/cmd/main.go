// Package main - точка входа и вспомогательные методы запуска приложения
package main

import (
	"github.com/artem-benda/gophkeeper/server/internal/infrastructure/db"
	"log/slog"
)

// main - точка входа приложения
func main() {
	slog.Info("Starting client application", slog.String("versionInfo", VersionString()))
	config := mustReadConfig()
	db.MustRunDBMigrations(config.DatabaseDSN)
	dbPool := db.MustCreateConnectionPool(config.DatabaseDSN)
	deps := mustCreateAppDependencies(dbPool, config)

	mustRunGrpcServer(deps.userService, deps.secretService, config)
}
