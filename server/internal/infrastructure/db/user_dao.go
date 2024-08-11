package db

import (
	"context"
	"errors"

	"github.com/artem-benda/gophkeeper/server/internal/domain/contract"
	"github.com/artem-benda/gophkeeper/server/internal/domain/entity"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

// UserDAO - методы для работы с пользователями
type UserDAO struct {
	DB *pgxpool.Pool
}

// GetByLogin - получение пользователя по логину
func (dao *UserDAO) GetByLogin(ctx context.Context, login string) (*entity.User, error) {
	user := entity.User{}

	row := dao.DB.QueryRow(ctx, "SELECT id, login, password_hash FROM users WHERE login = $1", login)
	err := row.Scan(&user.ID, &user.Login, &user.PasswordHash)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, contract.ErrUserNotFound
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetByID - получение пользователя по ID
func (dao *UserDAO) GetByID(ctx context.Context, userID int64) (*entity.User, error) {
	user := entity.User{}

	row := dao.DB.QueryRow(ctx, "SELECT id, login, password_hash FROM users WHERE id = $1", userID)
	err := row.Scan(&user.ID, &user.Login, &user.PasswordHash)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Insert - добавление пользователя
func (dao *UserDAO) Insert(ctx context.Context, user entity.User) (*int64, error) {
	userID := new(int64)
	row := dao.DB.QueryRow(ctx, "insert into users(login, password_hash) values($1, $2) returning id", user.Login, user.PasswordHash)
	err := row.Scan(userID)
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
		return nil, contract.ErrUserAlreadyRegistered
	}
	if err != nil {
		return nil, err
	}
	return userID, nil
}
