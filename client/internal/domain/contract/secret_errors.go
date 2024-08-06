package contract

import "errors"

var (
	ErrSecretAlreadyExists = errors.New("secret with provided guid already exists")
	ErrSecretNotFound      = errors.New("secret with provided guid does not exist")
)
