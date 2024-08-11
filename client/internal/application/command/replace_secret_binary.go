package command

import (
	"context"
	"fmt"
	"os"

	"github.com/artem-benda/gophkeeper/client/internal/application"
	"github.com/artem-benda/gophkeeper/client/internal/infrastructure/ctx"
	"github.com/jessevdk/go-flags"
)

// ReplaceSecretBinaryCommand replaces secret binary for specified secret
type ReplaceSecretBinaryCommand struct {
	Login          string `short:"l" long:"login" description:"login, non empty string" required:"true"`
	Password       string `short:"p" long:"password" description:"password, non empty string" required:"true"`
	GUID           string `short:"g" long:"guid" description:"secret guid, non empty string" required:"true"`
	Name           string `long:"name" description:"name, non empty string" required:"true"`
	SecretFilePath string `long:"secret-file-path" description:"file path to store secret binary from" required:"true"`
	Metadata       string `long:"metadata" description:"metadata"`
}

// IsReplaceSecretBinaryCommand returns true if command is ReplaceSecretBinaryCommand
func IsReplaceSecretBinaryCommand(parser *flags.Parser) bool {
	return parser.Active.Name == "replace-secret-binary"
}

// HandleReplaceSecretBinaryCommand handles ReplaceSecretBinaryCommand
func HandleReplaceSecretBinaryCommand(deps *application.AppDependencies, parser *flags.Parser, cmd *ReplaceSecretBinaryCommand) error {
	loginCtx := context.Background()
	token, err := deps.US.Login(loginCtx, cmd.Login, cmd.Password)
	if err != nil {
		return err
	}
	binary, err := os.ReadFile(cmd.SecretFilePath)
	if err != nil {
		return fmt.Errorf("error reding file: %w", err)
	}
	// Новый контекст для выполнения запроса, требующего авторизации
	ctxWithAuth := ctx.WithAuthToken(context.Background(), token)
	return deps.SS.ReplaceWithBinary(ctxWithAuth, cmd.GUID, cmd.Name, binary, cmd.Metadata)
}
