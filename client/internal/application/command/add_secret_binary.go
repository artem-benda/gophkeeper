package command

import (
	"context"
	"os"

	"github.com/artem-benda/gophkeeper/client/internal/application"
	"github.com/artem-benda/gophkeeper/client/internal/infrastructure/ctx"
	"github.com/jessevdk/go-flags"
)

// AddSecretBinaryCommand is a command for adding secret binary
type AddSecretBinaryCommand struct {
	Login          string `short:"l" long:"login" description:"login, non empty string" required:"true"`
	Password       string `short:"p" long:"password" description:"password, non empty string" required:"true"`
	Name           string `long:"name" description:"name, non empty string" required:"true"`
	SecretFilePath string `long:"secret-file-path" description:"file path to store secret binary from" required:"true"`
	Metadata       string `long:"metadata" description:"metadata"`
}

// IsAddSecretBinaryCommand returns true if command is AddSecretBinaryCommand
func IsAddSecretBinaryCommand(parser *flags.Parser) bool {
	return parser.Active.Name == "add-secret-binary"
}

// HandleAddSecretBinaryCommand handles AddSecretBinaryCommand
func HandleAddSecretBinaryCommand(deps *application.AppDependencies, parser *flags.Parser, cmd *AddSecretBinaryCommand) (string, error) {
	loginCtx := context.Background()
	token, err := deps.US.Login(loginCtx, cmd.Login, cmd.Password)
	if err != nil {
		return "", err
	}
	binary, err := os.ReadFile(cmd.SecretFilePath)
	if err != nil {
		return "", err
	}
	// Новый контекст для выполнения запроса, требующего авторизации
	ctxWithAuth := ctx.WithAuthToken(context.Background(), token)
	return deps.SS.AddBinary(ctxWithAuth, cmd.Name, binary, cmd.Metadata)
}
