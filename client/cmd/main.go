package main

import (
	"log/slog"

	"github.com/artem-benda/gophkeeper/client/internal/application/command"
	"github.com/artem-benda/gophkeeper/client/internal/infrastructure/grpc"
)

// main function is the entry point for the application
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
		} else {
			slog.Info("Successfully executed register command")
		}
		// series of add commands
	case command.IsAddSecretLoginPasswordCommand(parser):
		slog.Info("Executing add secret login and password command...")
		guid, err := command.HandleAddSecretLoginPasswordCommand(deps, parser, &opts.AddSecretLoginPassword)
		if err != nil {
			slog.Error("Failed to execute add secret login and password  ", slog.Any("error", err))
		} else {
			slog.Info("Successfully executed add secret login and password command", slog.String("guid", guid))
		}
	case command.IsAddSecretBankingCardCommand(parser):
		slog.Info("Executing add secret banking card command...")
		guid, err := command.HandleAddSecretBankingCardCommand(deps, parser, &opts.AddSecretBankingCard)
		if err != nil {
			slog.Error("Failed to execute add secret banking card  ", slog.Any("error", err))
		} else {
			slog.Info("Successfully executed add secret login and password command", slog.String("guid", guid))
		}
	case command.IsAddSecretTextCommand(parser):
		slog.Info("Executing add secret text command...")
		guid, err := command.HandleAddSecretTextCommand(deps, parser, &opts.AddSecretText)
		if err != nil {
			slog.Error("Failed to execute add secret text  ", slog.Any("error", err))
		} else {
			slog.Info("Successfully executed add secret text command", slog.String("guid", guid))
		}
	case command.IsAddSecretBinaryCommand(parser):
		slog.Info("Executing add secret binary from file command...")
		guid, err := command.HandleAddSecretBinaryCommand(deps, parser, &opts.AddSecretBinary)
		if err != nil {
			slog.Error("Failed to execute add secret binary from file  ", slog.Any("error", err))
		} else {
			slog.Info("Successfully executed add secret binary from file command", slog.String("guid", guid))
		}
		// series of replace commands
	case command.IsReplaceSecretLoginPasswordCommand(parser):
		slog.Info("Executing replace secret login and password command...")
		err := command.HandleReplaceSecretLoginPasswordCommand(deps, parser, &opts.ReplaceSecretLoginPassword)
		if err != nil {
			slog.Error("Failed to execute replace secret login and password  ", slog.Any("error", err))
		} else {
			slog.Info("Successfully executed replace secret login and password command")
		}
	case command.IsReplaceSecretBankingCardCommand(parser):
		slog.Info("Executing replace secret banking card command...")
		err := command.HandleReplaceSecretBankingCardCommand(deps, parser, &opts.ReplaceSecretBankingCard)
		if err != nil {
			slog.Error("Failed to execute replace secret banking card  ", slog.Any("error", err))
		} else {
			slog.Info("Successfully executed replace secret login and password command")
		}
	case command.IsReplaceSecretTextCommand(parser):
		slog.Info("Executing replace secret text command...")
		err := command.HandleReplaceSecretTextCommand(deps, parser, &opts.ReplaceSecretText)
		if err != nil {
			slog.Error("Failed to execute replace secret text  ", slog.Any("error", err))
		} else {
			slog.Info("Successfully executed replace secret text command")
		}
	case command.IsReplaceSecretBinaryCommand(parser):
		slog.Info("Executing replace secret binary from file command...")
		err := command.HandleReplaceSecretBinaryCommand(deps, parser, &opts.ReplaceSecretBinary)
		if err != nil {
			slog.Error("Failed to execute replace secret binary from file  ", slog.Any("error", err))
		} else {
			slog.Info("Successfully executed replace secret binary from file command")
		}
	case command.IsRemoveSecretCommand(parser):
		slog.Info("Executing remove secret command...")
		err := command.HandleRemoveSecretCommand(deps, parser, &opts.RemoveSecret)
		if err != nil {
			slog.Error("Failed to execute remove secret  ", slog.Any("error", err))
		} else {
			slog.Info("Successfully executed remove secret command")
		}
	case command.IsGetSecretCommand(parser):
		slog.Info("Executing get secret command...")
		err := command.HandleGetSecretCommand(deps, parser, &opts.GetSecret)
		if err != nil {
			slog.Error("Failed to execute get secret  ", slog.Any("error", err))
		} else {
			slog.Info("Successfully executed get secret command")
		}
	case command.IsGetAllSecretsCommand(parser):
		slog.Info("Executing get all secret command...")
		err := command.HandleGetAllSecretsCommand(deps, parser, &opts.GetAllSecrets)
		if err != nil {
			slog.Error("Failed to execute get all secrets  ", slog.Any("error", err))
		} else {
			slog.Info("Successfully executed get all secrets command")
		}
	default:
		slog.Error("Unknown command")
	}
}
