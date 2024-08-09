package command

import (
	"context"

	"github.com/artem-benda/gophkeeper/client/internal/application"
	"github.com/jessevdk/go-flags"
)

type AddLoginPasswordSecretCommand struct {
	Login    string `short:"l" long:"login" description:"username, non empty string" required:"true"`
	Password string `short:"p" long:"password" description:"password, non empty string" required:"true"`
	SecretLogin string `short:"sl" long:"secret-login" description:"secret login, non empty string"`
	SecretPassword string `short:"sp" long:"secret-password" description:"secret password, non empty string"`
}

func IsAddLoginPasswordSecretCommand(parser *flags.Parser) bool {
	return parser.Command.Find("add-secret-login-password") != nil
}

func HandleLoginPasswordSecretCommand(deps *application.AppDependencies, parser *flags.Parser, cmd *RegisterCommand) error {
	ctx := context.Background()
	deps.US.Login(ctx, cmd.Login, cmd.Password)
	return deps.SS.
}
