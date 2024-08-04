package db

import (
	"context"
	"errors"

	"github.com/artem-benda/gophkeeper/server/internal/domain/entity"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserDAO struct {
	DB *pgxpool.Pool
}

func (dao *UserDAO) GetByLogin(ctx context.Context, login string) (*entity.User, error) {
	user := entity.User{}

	row := dao.DB.QueryRow(ctx, "SELECT id, login, password_hash FROM users WHERE login = $1", login)
	err := row.Scan(&user.ID, &user.Login, &user.PasswordHash)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (dao *UserDAO) GetByID(ctx context.Context, userID int64) (*entity.User, error) {
	user := entity.User{}

	row := dao.DB.QueryRow(ctx, "SELECT id, login, password_hash FROM users WHERE id = $1", userID)
	err := row.Scan(&user.ID, &user.Login, &user.PasswordHash)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (dao *UserDAO) Insert(ctx context.Context, user entity.User) (*int64, error) {
	userID := new(int64)
	row := dao.DB.QueryRow(ctx, "insert into users(login, password_hash) values($1, $2) returning id", user.Login, user.PasswordHash)
	err := row.Scan(userID)
	if err != nil {
		return nil, err
	}
	return userID, nil
}
