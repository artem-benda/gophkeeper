package command

import (
	"context"

	"github.com/artem-benda/gophkeeper/client/internal/application"
	"github.com/artem-benda/gophkeeper/client/internal/infrastructure/ctx"
	"github.com/jessevdk/go-flags"
)

// ReplaceSecretLoginPasswordCommand represents command to replace secret login and password
type ReplaceSecretLoginPasswordCommand struct {
	Login          string `short:"l" long:"login" description:"login, non empty string" required:"true"`
	Password       string `short:"p" long:"password" description:"password, non empty string" required:"true"`
	GUID           string `short:"g" long:"guid" description:"secret guid, non empty string" required:"true"`
	Name           string `long:"name" description:"name, non empty string" required:"true"`
	SecretLogin    string `long:"secret-login" description:"secret login"`
	SecretPassword string `long:"secret-password" description:"secret password"`
	Metadata       string `long:"metadata" description:"metadata"`
}

// IsReplaceSecretLoginPasswordCommand returns true if command is ReplaceSecretLoginPasswordCommand
func IsReplaceSecretLoginPasswordCommand(parser *flags.Parser) bool {
	return parser.Active.Name == "replace-secret-login-password"
}

// HandleReplaceSecretLoginPasswordCommand handles ReplaceSecretLoginPasswordCommand
func HandleReplaceSecretLoginPasswordCommand(deps *application.AppDependencies, parser *flags.Parser, cmd *ReplaceSecretLoginPasswordCommand) error {
	loginCtx := context.Background()
	token, err := deps.US.Login(loginCtx, cmd.Login, cmd.Password)
	if err != nil {
		return err
	}
	// Новый контекст для выполнения запроса, требующего авторизации
	ctxWithAuth := ctx.WithAuthToken(context.Background(), token)
	return deps.SS.ReplaceWithLoginPassword(ctxWithAuth, cmd.GUID, cmd.Name, cmd.SecretLogin, cmd.SecretPassword, cmd.Metadata)
}
