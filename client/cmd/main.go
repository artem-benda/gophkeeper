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
	deps := mustCreateAppDependencies(client, opts.PassPhrase)
	switch {
	case command.IsRegisterCommand(parser):
		slog.Info("Executing register command...")
		err := command.HandleRegisterCommand(deps, parser, &opts.Register)
		if err != nil {
			slog.Error("Failed to execute register  ", slog.Any("error", err))
		}
		slog.Info("Successfully executed register command")
	case command.IsAddSecretLoginPasswordCommand(parser):
		slog.Info("Executing add secret login and password command...")
		guid, err := command.HandleAddSecretLoginPasswordCommand(deps, parser, &opts.AddSecretLoginPassword)
		if err != nil {
			slog.Error("Failed to execute add secret login and password  ", slog.Any("error", err))
		}
		slog.Info("Successfully executed add secret login and password command", slog.String("guid", guid))
	case command.IsAddSecretBankingCardCommand(parser):
		slog.Info("Executing add secret login and password command...")
		guid, err := command.HandleAddSecretBankingCardCommand(deps, parser, &opts.AssSecretBankingCardd)
		if err != nil {
			slog.Error("Failed to execute add secret login and password  ", slog.Any("error", err))
		}
		slog.Info("Successfully executed add secret login and password command", slog.String("guid", guid))
	default:
		slog.Error("Unknown command")
	}
}
