package command

import (
	"context"

	"github.com/artem-benda/gophkeeper/client/internal/application"
	"github.com/artem-benda/gophkeeper/client/internal/infrastructure/ctx"
	"github.com/jessevdk/go-flags"
)

// AddSecretTextCommand is a command for adding secret text
type AddSecretTextCommand struct {
	Login      string `short:"l" long:"login" description:"login, non empty string" required:"true"`
	Password   string `short:"p" long:"password" description:"password, non empty string" required:"true"`
	Name       string `long:"name" description:"name, non empty string" required:"true"`
	SecretText string `long:"secret-text" description:"secret text"`
	Metadata   string `long:"metadata" description:"metadata"`
}

// IsAddSecretTextCommand returns true if command is AddSecretTextCommand
func IsAddSecretTextCommand(parser *flags.Parser) bool {
	return parser.Active.Name == "add-secret-text"
}

// HandleAddSecretTextCommand handles AddSecretTextCommand
func HandleAddSecretTextCommand(deps *application.AppDependencies, parser *flags.Parser, cmd *AddSecretTextCommand) (string, error) {
	loginCtx := context.Background()
	token, err := deps.US.Login(loginCtx, cmd.Login, cmd.Password)
	if err != nil {
		return "", err
	}
	// Новый контекст для выполнения запроса, требующего авторизации
	ctxWithAuth := ctx.WithAuthToken(context.Background(), token)
	return deps.SS.AddText(ctxWithAuth, cmd.Name, cmd.SecretText, cmd.Metadata)
}
