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

type userRepository struct {
	DAO db.UserDAO
}

func NewUserRepository(dao db.UserDAO) contract.UserRepository {
	return &userRepository{dao}
}

func (r *userRepository) Register(ctx context.Context, login string, passwordHash string) (*int64, error) {
	id, err := r.DAO.Insert(ctx, entity.User{Login: login, PasswordHash: passwordHash})
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.IntegrityConstraintViolation {
		return nil, contract.ErrUserAlreadyRegistered
	}
	return id, err
}

func (r *userRepository) GetUserByLogin(ctx context.Context, login string) (*entity.User, error) {
	return r.DAO.GetByLogin(ctx, login)
}

func (r *userRepository) GetUserByID(ctx context.Context, userID int64) (*entity.User, error) {
	return r.DAO.GetByID(ctx, userID)
}
