package command

import (
	"context"

	"github.com/artem-benda/gophkeeper/client/internal/application"
	"github.com/artem-benda/gophkeeper/client/internal/infrastructure/ctx"
	"github.com/jessevdk/go-flags"
)

// ReplaceSecretTextCommand - replace secret text for specified secret
type ReplaceSecretTextCommand struct {
	Login      string `short:"l" long:"login" description:"login, non empty string" required:"true"`
	Password   string `short:"p" long:"password" description:"password, non empty string" required:"true"`
	GUID       string `short:"g" long:"guid" description:"secret guid, non empty string" required:"true"`
	Name       string `long:"name" description:"name, non empty string" required:"true"`
	SecretText string `long:"secret-text" description:"secret text"`
	Metadata   string `long:"metadata" description:"metadata"`
}

// IsReplaceSecretTextCommand - check if command is ReplaceSecretTextCommand
func IsReplaceSecretTextCommand(parser *flags.Parser) bool {
	return parser.Active.Name == "add-secret-text"
}

// HandleReplaceSecretTextCommand - handle ReplaceSecretTextCommand
func HandleReplaceSecretTextCommand(deps *application.AppDependencies, parser *flags.Parser, cmd *ReplaceSecretTextCommand) error {
	loginCtx := context.Background()
	token, err := deps.US.Login(loginCtx, cmd.Login, cmd.Password)
	if err != nil {
		return err
	}
	// Новый контекст для выполнения запроса, требующего авторизации
	ctxWithAuth := ctx.WithAuthToken(context.Background(), token)
	return deps.SS.ReplaceWithText(ctxWithAuth, cmd.GUID, cmd.Name, cmd.SecretText, cmd.Metadata)
}
