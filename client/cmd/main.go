package main

import (
	"github.com/artem-benda/gphkeeper/client/internal/application/command"
	"log/slog"
)

func main() {
	slog.Info("Starting client application", slog.String("versionInfo", VersionString()))
	var opts = new(Options)
	parser := mustParseOptions(opts)
	err := command.HandleRegisterCommand(parser, opts)
	if err != nil {
		slog.Error("Failed to execute register  ", err)
	}
}
