package command

import (
	"github.com/artem-benda/gphkeeper/client/internal/application"
	"github.com/jessevdk/go-flags"
)

type RegisterCommand struct {
	login    string `long:"login" description:"username, non empty string" required:"true"`
	password string `long:"password" description:"password, non empty string" required:"true"`
}

func IsRegisterCommand(parser *flags.Parser) bool {
	return parser.Command.Find("register") != nil
}

func HandleRegisterCommand(deps *application.AppDependencies, parser *flags.Parser, cmd *RegisterCommand) error {
	return deps.US.Register(cmd.login, cmd.password)
}
