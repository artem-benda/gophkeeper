package main

import (
	"log/slog"

	"github.com/artem-benda/gophermart/client/internal/application"
)

func main() {
	slog.Info("Starting client application", slog.String("versionInfo", VersionString()))
	config := mustReadConfig()
	db := mustCreateDB(config.DatabaseFilePath)
	app := application.NewApp()

	app.Run(db)
}
