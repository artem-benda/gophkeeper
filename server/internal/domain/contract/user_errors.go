package contract

import "errors"

var (
	ErrUserAlreadyRegistered = errors.New("user already registered")
	ErrUserNotFound          = errors.New("user not found")
	ErrUnauthorized          = errors.New("unauthorized")
)
