package application

import "github.com/artem-benda/gophkeeper/client/internal/domain/contract"

// AppDependencies - application dependencies
type AppDependencies struct {
	US contract.UserService
	SS contract.SecretService
}
