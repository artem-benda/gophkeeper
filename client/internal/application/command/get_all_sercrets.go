package command

import (
	"context"

	"github.com/artem-benda/gophkeeper/client/internal/application"
	"github.com/artem-benda/gophkeeper/client/internal/infrastructure/ctx"
	"github.com/jessevdk/go-flags"
)

// GetAllSecretsCommand - command for getting all secrets
type GetAllSecretsCommand struct {
	Login    string `short:"l" long:"login" description:"login, non empty string" required:"true"`
	Password string `short:"p" long:"password" description:"password, non empty string" required:"true"`
}

// IsGetAllSecretsCommand - returns true if command is GetAllSecretsCommand
func IsGetAllSecretsCommand(parser *flags.Parser) bool {
	return parser.Active.Name == "get-all-secrets"
}

// HandleGetAllSecretsCommand - handles GetAllSecretsCommand
func HandleGetAllSecretsCommand(deps *application.AppDependencies, parser *flags.Parser, cmd *GetAllSecretsCommand) error {
	loginCtx := context.Background()
	token, err := deps.US.Login(loginCtx, cmd.Login, cmd.Password)
	if err != nil {
		return err
	}
	// Новый контекст для выполнения запроса, требующего авторизации
	ctxWithAuth := ctx.WithAuthToken(context.Background(), token)
	secrets, err := deps.SS.GetAll(ctxWithAuth)
	if err != nil {
		return err
	}
	for _, secret := range secrets {
		printSecret(&secret)
	}
	return nil
}
