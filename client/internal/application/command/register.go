package command

import (
	"github.com/artem-benda/gphkeeper/client/internal/domain/contract"
	"github.com/jessevdk/go-flags"
)

type RegisterCommand struct {
	login    string `long:"login" description:"username, non empty string" required:"true"`
	password string `long:"password" description:"password, non empty string" required:"true"`
}

type RegisterCommandDeps struct {
	US contract.UserService
}

func (d *RegisterCommandDeps) HandleRegisterCommand(parser *flags.Parser, cmd *RegisterCommand) error {
	if parser.Command.Find("register") == nil {
		return d.US.Register(cmd.login, cmd.password)
	}
	return nil
}
