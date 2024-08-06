package main

import (
	"github.com/artem-benda/gphkeeper/client/internal/application/command"
	"log/slog"
)

func main() {
	slog.Info("Starting client application", slog.String("versionInfo", VersionString()))
	var opts = new(Options)
	parser := mustParseOptions(opts)
	deps := mustCreateAppDependencies()
	// Handle Register command or else continue
	switch {
	case command.IsRegisterCommand(parser):
		err := command.HandleRegisterCommand(deps, parser, &opts.Register)
		if err != nil {
			slog.Error("Failed to execute register  ", err)
		}
	}
}
