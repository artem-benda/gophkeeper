package service

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"github.com/artem-benda/gophkeeper/server/internal/domain/contract"
	"github.com/artem-benda/gophkeeper/server/internal/domain/entity"
	"golang.org/x/crypto/pbkdf2"
	"log/slog"
)

var _ contract.UserService = (*user)(nil)

func NewUserService(repo contract.UserRepository, salt []byte) contract.UserService {
	return &user{repo: repo, salt: salt}
}

type user struct {
	repo contract.UserRepository
	salt []byte
}

func (u *user) Register(ctx context.Context, login, password string) (*int64, error) {
	passwordHash, err := computeHash(password, u.salt)

	if err != nil {
		slog.Error("error registering user: ", err)
		return nil, err
	}

	return u.repo.Register(ctx, login, *passwordHash)
}

func (u *user) Login(ctx context.Context, login, password string) (*int64, error) {
	passwordHashString, err := computeHash(password, u.salt)
	if err != nil {
		return nil, err
	}
	user, err := u.repo.GetUserByLogin(ctx, login)

	if err != nil {
		slog.Error("error getting user by login: ", err)
		return nil, err
	}

	if user == nil {
		slog.Debug("user not found")
		return nil, contract.ErrUserNotFound
	}

	if user.PasswordHash != *passwordHashString {
		slog.Debug("hash mismatch: ", "expected: ", user.PasswordHash, ", actual: ", *passwordHashString)
		return nil, contract.ErrUnauthorized
	}

	return &user.ID, nil
}

func (u *user) GetUserByID(ctx context.Context, userID int64) (*entity.User, error) {
	user, err := u.repo.GetUserByID(ctx, userID)
	if err != nil {
		slog.Error("error getting user by ID: ", err)
		return nil, err
	}

	return user, nil
}

func computeHash(password string, salt []byte) (*string, error) {
	pwPbkdf2 := pbkdf2.Key([]byte(password), salt, 10240, 32, sha256.New)
	encodedHash := hex.EncodeToString(pwPbkdf2)

	return &encodedHash, nil
}
