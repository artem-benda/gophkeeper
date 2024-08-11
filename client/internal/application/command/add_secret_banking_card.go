package command

import (
	"context"

	"github.com/artem-benda/gophkeeper/client/internal/application"
	"github.com/artem-benda/gophkeeper/client/internal/infrastructure/ctx"
	"github.com/jessevdk/go-flags"
)

// AddSecretBankingCardCommand is a command for adding a secret banking card
type AddSecretBankingCardCommand struct {
	Login     string `short:"l" long:"login" description:"login, non empty string" required:"true"`
	Password  string `short:"p" long:"password" description:"password, non empty string" required:"true"`
	Name      string `long:"name" description:"name, non empty string" required:"true"`
	Owner     string `long:"card-owner" description:"banking card owner"`
	Number    string `long:"card-number" description:"banking card number"`
	ValidThru string `long:"card-valid-thru" description:"banking card valid thru"`
	CVV       string `long:"card-cvv" description:"banking card cvv/cv2"`
	Metadata  string `long:"metadata" description:"metadata"`
}

// IsAddSecretBankingCardCommand returns true if command is AddSecretBankingCardCommand
func IsAddSecretBankingCardCommand(parser *flags.Parser) bool {
	return parser.Active.Name == "add-secret-banking-card"
}

// HandleAddSecretBankingCardCommand executes AddSecretBankingCardCommand
func HandleAddSecretBankingCardCommand(deps *application.AppDependencies, parser *flags.Parser, cmd *AddSecretBankingCardCommand) (string, error) {
	loginCtx := context.Background()
	token, err := deps.US.Login(loginCtx, cmd.Login, cmd.Password)
	if err != nil {
		return "", err
	}
	// Новый контекст для выполнения запроса, требующего авторизации
	ctxWithAuth := ctx.WithAuthToken(context.Background(), token)
	return deps.SS.AddBankingCard(ctxWithAuth, cmd.Name, cmd.Number, cmd.Owner, cmd.ValidThru, cmd.CVV, cmd.Metadata)
}
