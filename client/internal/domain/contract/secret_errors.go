package contract

import "errors"

// Ошибки бизнес логики для секретных данных
var (
	ErrSecretAlreadyExists = errors.New("secret with provided guid already exists")
	ErrSecretNotFound      = errors.New("secret with provided guid does not exist")
)
