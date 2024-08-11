package repository

import (
	"context"
	"errors"

	"github.com/artem-benda/gophkeeper/server/internal/domain/contract"
	"github.com/artem-benda/gophkeeper/server/internal/domain/entity"
	"github.com/artem-benda/gophkeeper/server/internal/infrastructure/db"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
)

var _ contract.UserRepository = (*userRepository)(nil)

// userRepository implements contract.UserRepository
type userRepository struct {
	DAO db.UserDAO
}

// NewUserRepository creates new instance of userRepository
func NewUserRepository(dao db.UserDAO) contract.UserRepository {
	return &userRepository{dao}
}

// Register registers new user
func (r *userRepository) Register(ctx context.Context, login string, passwordHash string) (*int64, error) {
	id, err := r.DAO.Insert(ctx, entity.User{Login: login, PasswordHash: passwordHash})
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.IntegrityConstraintViolation {
		return nil, contract.ErrUserAlreadyRegistered
	}
	return id, err
}

// GetUserByLogin returns user by login
func (r *userRepository) GetUserByLogin(ctx context.Context, login string) (*entity.User, error) {
	return r.DAO.GetByLogin(ctx, login)
}

// GetUserByID returns user by id
func (r *userRepository) GetUserByID(ctx context.Context, userID int64) (*entity.User, error) {
	return r.DAO.GetByID(ctx, userID)
}
