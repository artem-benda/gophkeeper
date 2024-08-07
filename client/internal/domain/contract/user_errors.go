package contract

import "errors"

var (
	ErrUserAlreadyExists      = errors.New("user with provided login already exists")
	ErrUserInvalidCredentials = errors.New("wrong login or password")
)
