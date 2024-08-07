package contract

import "errors"

// доменные ошибки работы с пользователями
var (
	ErrUserAlreadyRegistered = errors.New("user already registered")
	ErrUserNotFound          = errors.New("user not found")
	ErrUnauthorized          = errors.New("unauthorized")
)
