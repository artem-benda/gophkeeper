package contract

import "errors"

var (
	ErrSecretAlreadyExists = errors.New("secret already exists")
	ErrSecretNotFound      = errors.New("secret not found")
)
