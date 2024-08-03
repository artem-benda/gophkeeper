package db

import (
	"context"
	"errors"
	"github.com/artem-benda/gophkeeper/server/internal/domain/entity"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type SecretDAO struct {
	DB *pgxpool.Pool
}

func (dao SecretDAO) getByID(ctx context.Context, id int64) (*entity.Secret, error) {
	secret := entity.Secret{}
	row := dao.DB.QueryRow(ctx, "SELECT id, name, enc_payload, created_at, updated_at FROM secret WHERE id = $1", id)
	err := row.Scan(&secret.ID, &secret.Name, &secret.EncPayload, &secret.CreatedAt, &secret.UpdatedAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &secret, nil
}

func (dao SecretDAO) getByUserID(ctx context.Context, userID int64) ([]entity.Secret, error) {
	rows, err := dao.DB.Query(ctx, "SELECT id, name, enc_payload, created_at, updated_at FROM secrets WHERE user_id = $1 ORDER BY id", userID)
	if err != nil {
		return nil, err
	}

	secrets := make([]entity.Secret, 0)

	for rows.Next() {
		secret := entity.Secret{}
		err := rows.Scan(&secret.ID, &secret.Name, &secret.EncPayload, &secret.CreatedAt, &secret.UpdatedAt)
		if err != nil {
			return nil, err
		}
		secrets = append(secrets, secret)
	}

	if rows.Err() != nil {
		return nil, err
	}
	return secrets, nil
}
