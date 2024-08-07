package command

import (
	"context"

	"github.com/artem-benda/gophkeeper/client/internal/application"
	"github.com/artem-benda/gophkeeper/client/internal/infrastructure/ctx"
	"github.com/jessevdk/go-flags"
)

// AddSecretLoginPasswordCommand adds new secret login password
type AddSecretLoginPasswordCommand struct {
	Login          string `short:"l" long:"login" description:"login, non empty string" required:"true"`
	Password       string `short:"p" long:"password" description:"password, non empty string" required:"true"`
	Name           string `long:"name" description:"name, non empty string" required:"true"`
	SecretLogin    string `long:"secret-login" description:"secret login"`
	SecretPassword string `long:"secret-password" description:"secret password"`
	Metadata       string `long:"metadata" description:"metadata"`
}

// IsAddSecretLoginPasswordCommand returns true if command is AddSecretLoginPasswordCommand
func IsAddSecretLoginPasswordCommand(parser *flags.Parser) bool {
	return parser.Active.Name == "add-secret-login-password"
}

// HandleAddSecretLoginPasswordCommand handles AddSecretLoginPasswordCommand
func HandleAddSecretLoginPasswordCommand(deps *application.AppDependencies, parser *flags.Parser, cmd *AddSecretLoginPasswordCommand) (string, error) {
	loginCtx := context.Background()
	token, err := deps.US.Login(loginCtx, cmd.Login, cmd.Password)
	if err != nil {
		return "", err
	}
	// Новый контекст для выполнения запроса, требующего авторизации
	ctxWithAuth := ctx.WithAuthToken(context.Background(), token)
	return deps.SS.AddLoginPassword(ctxWithAuth, cmd.Name, cmd.SecretLogin, cmd.SecretPassword, cmd.Metadata)
}
