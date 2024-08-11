package contract

import "errors"

// Ошибки биизнес-логики для пользователей
var (
	ErrUserAlreadyExists      = errors.New("user with provided login already exists")
	ErrUserInvalidCredentials = errors.New("wrong login or password")
)
