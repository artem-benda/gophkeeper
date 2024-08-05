package main

import (
	"errors"
	"github.com/artem-benda/gphkeeper/client/internal/application/command"
	"github.com/jessevdk/go-flags"
	"os"
)

type Options struct {
	Endpoint  string                  `short:"e" long:"endpoint" description:"address and port of server to connect to" required:"true"`
	CryptoKey string                  `short:"k" long:"crypto-key" description:"crypto key in base64std format to store your data securely" required:"true"`
	Register  command.RegisterCommand `command:"register"`
}

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
			os.Exit(1)
		default:
			os.Exit(1)
		}
	}

	if p.Active == nil {
		panic("Expected command, run with --help option to see commands list")
	}
	return p
}
