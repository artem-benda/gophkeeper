package command

import (
	"context"

	"github.com/artem-benda/gophkeeper/client/internal/application"
	"github.com/artem-benda/gophkeeper/client/internal/infrastructure/ctx"
	"github.com/jessevdk/go-flags"
)

// RemoveSecretCommand - remove secret command
type RemoveSecretCommand struct {
	Login    string `short:"l" long:"login" description:"login, non empty string" required:"true"`
	Password string `short:"p" long:"password" description:"password, non empty string" required:"true"`
	GUID     string `short:"g" long:"guid" description:"secret guid, non empty string" required:"true"`
}

// IsRemoveSecretCommand - return true if command is remove secret command
func IsRemoveSecretCommand(parser *flags.Parser) bool {
	return parser.Active.Name == "remove-secret"
}

// HandleRemoveSecretCommand - handle remove secret command
func HandleRemoveSecretCommand(deps *application.AppDependencies, parser *flags.Parser, cmd *RemoveSecretCommand) error {
	loginCtx := context.Background()
	token, err := deps.US.Login(loginCtx, cmd.Login, cmd.Password)
	if err != nil {
		return err
	}
	// Новый контекст для выполнения запроса, требующего авторизации
	ctxWithAuth := ctx.WithAuthToken(context.Background(), token)
	return deps.SS.Remove(ctxWithAuth, cmd.GUID)
}
