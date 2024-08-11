package contract

import "errors"

// доменные ошибки работы с секретами
var (
	ErrSecretAlreadyExists = errors.New("secret already exists")
	ErrSecretNotFound      = errors.New("secret not found")
)
