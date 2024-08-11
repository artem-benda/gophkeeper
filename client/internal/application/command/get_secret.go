package command

import (
	"context"
	"log/slog"

	"github.com/artem-benda/gophkeeper/client/internal/application"
	"github.com/artem-benda/gophkeeper/client/internal/domain/entity"
	"github.com/artem-benda/gophkeeper/client/internal/infrastructure/ctx"
	"github.com/jessevdk/go-flags"
)

// GetSecretCommand - Команда получения секрета
type GetSecretCommand struct {
	Login    string `short:"l" long:"login" description:"login, non empty string" required:"true"`
	Password string `short:"p" long:"password" description:"password, non empty string" required:"true"`
	GUID     string `short:"g" long:"guid" description:"secret guid, non empty string" required:"true"`
}

// IssGetSecretCommand - returns true if command is GetSecretCommand
func IsGetSecretCommand(parser *flags.Parser) bool {
	return parser.Active.Name == "get-secret"
}

// HandleGetSecretCommand - обработчик команды get-secret
func HandleGetSecretCommand(deps *application.AppDependencies, parser *flags.Parser, cmd *GetSecretCommand) error {
	loginCtx := context.Background()
	token, err := deps.US.Login(loginCtx, cmd.Login, cmd.Password)
	if err != nil {
		return err
	}
	// Новый контекст для выполнения запроса, требующего авторизации
	ctxWithAuth := ctx.WithAuthToken(context.Background(), token)
	secret, err := deps.SS.GetByGUID(ctxWithAuth, cmd.GUID)
	if err != nil {
		return err
	}
	printSecret(secret)
	return nil
}

// printSecret - Функция вывода секрета в лог
func printSecret(secret *entity.Secret) {
	switch v := secret.Payload.Secret.(type) {
	case *entity.SecretPayload_LoginPassword:
		slog.Info(
			"get secret result (secret login and password)",
			slog.String("GUID", secret.GUID),
			slog.String("Name", secret.Name),
			slog.String("Metadata", secret.Payload.Metadata),
			slog.String("Login", v.LoginPassword.Login),
			slog.String("Password", v.LoginPassword.Password),
			slog.Time("CreatedAt", secret.CreatedAt),
			slog.Time("UpdatedAt", secret.UpdatedAt),
		)
	case *entity.SecretPayload_Text:
		slog.Info(
			"get secret result (secret text)",
			slog.String("GUID", secret.GUID),
			slog.String("Name", secret.Name),
			slog.String("Metadata", secret.Payload.Metadata),
			slog.String("Text", v.Text.Text),
			slog.Time("CreatedAt", secret.CreatedAt),
			slog.Time("UpdatedAt", secret.UpdatedAt),
		)
	case *entity.SecretPayload_BankingCard:
		slog.Info(
			"get secret result (secret banking card)",
			slog.String("GUID", secret.GUID),
			slog.String("Name", secret.Name),
			slog.String("Metadata", secret.Payload.Metadata),
			slog.String("CardNumber", v.BankingCard.CardNumber),
			slog.String("CardValidThru", v.BankingCard.ValidThru),
			slog.String("CardOwner", v.BankingCard.OwnerName),
			slog.String("CVV/CV2", v.BankingCard.CVV),
			slog.Time("CreatedAt", secret.CreatedAt),
			slog.Time("UpdatedAt", secret.UpdatedAt),
		)
	case *entity.SecretPayload_Binary:
		slog.Info(
			"get secret result (secret binary)",
			slog.String("GUID", secret.GUID),
			slog.String("Name", secret.Name),
			slog.String("Metadata", secret.Payload.Metadata),
			slog.String("Binary", v.Binary.String()),
			slog.Time("CreatedAt", secret.CreatedAt),
			slog.Time("UpdatedAt", secret.UpdatedAt),
		)
	}
}
