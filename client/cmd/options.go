package main

import (
	"errors"
	"os"

	"github.com/artem-benda/gophkeeper/client/internal/application/command"
	"github.com/jessevdk/go-flags"
)

type Options struct {
	Endpoint   string                  `short:"e" long:"endpoint" default:"localhost:8080" description:"address and port of server to connect to" required:"true"`
	PassPhrase string                  `short:"k" long:"pass-phrase" default:"supersecretkey123" description:"pass phrase used to encrypt secret information" required:"true"`
	Register   command.RegisterCommand `command:"register" alias:"r" description:"Register user"`
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
