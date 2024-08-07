package main

import (
	"errors"
	"log/slog"
	"os"

	"github.com/artem-benda/gophkeeper/client/internal/application/command"
	"github.com/jessevdk/go-flags"
)

// Options - go-flags options
type Options struct {
	Endpoint                   string                                    `short:"e" long:"endpoint" default:"localhost:8080" description:"address and port of server to connect to" required:"true"`
	PassPhrase                 string                                    `short:"k" long:"pass-phrase" default:"supersecretkey123" description:"pass phrase used to encrypt secret information" required:"true"`
	Register                   command.RegisterCommand                   `command:"register" alias:"r" description:"Register user"`
	AddSecretLoginPassword     command.AddSecretLoginPasswordCommand     `command:"add-secret-login-password" alias:"alp" description:"Add secret login password"`
	AddSecretBankingCard       command.AddSecretBankingCardCommand       `command:"add-secret-banking-card" alias:"abc" description:"Add secret banking card"`
	AddSecretText              command.AddSecretTextCommand              `command:"add-secret-text" alias:"at" description:"Add secret text"`
	AddSecretBinary            command.AddSecretBinaryCommand            `command:"add-secret-binary" alias:"ab" description:"Add secret binary from file"`
	ReplaceSecretLoginPassword command.ReplaceSecretLoginPasswordCommand `command:"replace-secret-login-password" alias:"rlp" description:"Replace secret with login password"`
	ReplaceSecretBankingCard   command.ReplaceSecretBankingCardCommand   `command:"replace-secret-banking-card" alias:"rbc" description:"Replace secret with banking card"`
	ReplaceSecretText          command.ReplaceSecretTextCommand          `command:"replace-secret-text" alias:"rt" description:"Replace secret with text"`
	ReplaceSecretBinary        command.ReplaceSecretBinaryCommand        `command:"replace-secret-binary" alias:"rb" description:"Replace secret with binary from file"`
	RemoveSecret               command.RemoveSecretCommand               `command:"remove-secret" alias:"rm" description:"Remove secret"`
	GetSecret                  command.GetSecretCommand                  `command:"get-secret" alias:"g" description:"Get secret"`
	GetAllSecrets              command.GetAllSecretsCommand              `command:"get-all-secrets" alias:"ga" description:"Get all secrets"`
}

// mustParseOptions - parse options from command line args
func mustParseOptions(opts *Options) *flags.Parser {
	p := flags.NewParser(opts, flags.Default)
	_, err := p.Parse()
	if err != nil {
		var flagsErr flags.ErrorType
		switch {
		case errors.As(err, &flagsErr):
			if errors.Is(flagsErr, flags.ErrHelp) {
				os.Exit(0)
			}
		default:
			os.Exit(1)
		}
		slog.Error("Failed to parse options: %s", slog.Any("error", err))
		os.Exit(1)
	}

	if p.Active == nil {
		panic("Expected command, run with --help option to see commands list")
	}
	return p
}
