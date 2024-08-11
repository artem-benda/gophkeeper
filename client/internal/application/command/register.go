package command

import (
	"context"

	"github.com/artem-benda/gophkeeper/client/internal/application"
	"github.com/jessevdk/go-flags"
)

// RegisterCommand represents command to register new user
type RegisterCommand struct {
	Login    string `short:"l" long:"login" description:"username, non empty string" required:"true"`
	Password string `short:"p" long:"password" description:"password, non empty string" required:"true"`
}

// IsRegisterCommand returns true if command is "register"
func IsRegisterCommand(parser *flags.Parser) bool {
	return parser.Active.Name == "register"
}

// HandleRegisterCommand handles "register" command
func HandleRegisterCommand(deps *application.AppDependencies, parser *flags.Parser, cmd *RegisterCommand) error {
	ctx := context.Background()
	return deps.US.Register(ctx, cmd.Login, cmd.Password)
}
