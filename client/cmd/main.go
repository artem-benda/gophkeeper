package main

import (
	"log/slog"

	"github.com/artem-benda/gophkeeper/client/internal/application/command"
	"github.com/artem-benda/gophkeeper/client/internal/infrastructure/grpc"
)

func main() {
	slog.Info("Starting client application", slog.String("versionInfo", VersionString()))
	var opts = new(Options)
	parser := mustParseOptions(opts)
	client, conn := grpc.MustCreateGRPCClient(opts.Endpoint)
	defer conn.Close()
	deps := mustCreateAppDependencies(client)
	// Handle Register command or else continue
	switch {
	case command.IsRegisterCommand(parser):
		err := command.HandleRegisterCommand(deps, parser, &opts.Register)
		if err != nil {
			slog.Error("Failed to execute register  ", slog.Any("error", err))
		}
	}
}
