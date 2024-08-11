package db

import (
	"context"
	"errors"
	"time"

	"github.com/artem-benda/gophkeeper/server/internal/domain/contract"
	"github.com/artem-benda/gophkeeper/server/internal/domain/entity"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/pgtype"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

// SecretDAO - методы для работы с секретами
type SecretDAO struct {
	DB *pgxpool.Pool
}

// GetByID - Получить по идентификатору
func (dao *SecretDAO) GetByGUID(ctx context.Context, userID int64, guid string) (*entity.Secret, error) {
	secret := entity.Secret{}
	createdAt := new(pgtype.Timestamptz)
	updatedAt := new(pgtype.Timestamptz)
	row := dao.DB.QueryRow(ctx, "SELECT id, guid, name, enc_payload, created_at, updated_at FROM secrets WHERE guid = $1 and user_id = $2", guid, userID)
	err := row.Scan(&secret.ID, &secret.GUID, &secret.Name, &secret.EncPayload, createdAt, updatedAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, contract.ErrSecretNotFound
	}
	if err != nil {
		return nil, err
	}
	if createdAt.Status != pgtype.Null {
		secret.CreatedAt = createdAt.Time
	}
	if updatedAt.Status != pgtype.Null {
		secret.UpdatedAt = updatedAt.Time
	}
	return &secret, nil
}

// GetByUserID - Получить все по идентификатору пользователя
func (dao *SecretDAO) GetByUserID(ctx context.Context, userID int64) ([]entity.Secret, error) {
	rows, err := dao.DB.Query(ctx, "SELECT id, guid, name, enc_payload, created_at, updated_at FROM secrets WHERE user_id = $1 ORDER BY id", userID)
	if err != nil {
		return nil, err
	}

	secrets := make([]entity.Secret, 0)

	for rows.Next() {
		secret := entity.Secret{}
		createdAt := new(pgtype.Timestamptz)
		updatedAt := new(pgtype.Timestamptz)
		err := rows.Scan(&secret.ID, &secret.GUID, &secret.Name, &secret.EncPayload, createdAt, updatedAt)
		if err != nil {
			return nil, err
		}
		if createdAt.Status != pgtype.Null {
			secret.CreatedAt = createdAt.Time
		}
		if updatedAt.Status != pgtype.Null {
			secret.UpdatedAt = updatedAt.Time
		}
		secrets = append(secrets, secret)
	}

	if rows.Err() != nil {
		return nil, err
	}
	return secrets, nil
}

// Insert - Добавить секрет
func (dao *SecretDAO) Insert(ctx context.Context, userID int64, guid string, name string, encPayload []byte, clientTimestamp time.Time) (*int64, error) {
	secretID := new(int64)
	row := dao.DB.QueryRow(ctx, "insert into secrets(user_id, guid, name, enc_payload, created_at) values($1, $2, $3, $4, $5) returning id", userID, guid, name, encPayload, clientTimestamp)
	err := row.Scan(secretID)
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
		return nil, contract.ErrSecretAlreadyExists
	}
	if err != nil {
		return nil, err
	}
	return secretID, nil
}

// Update - Обновить секрет
func (dao *SecretDAO) Update(ctx context.Context, userID int64, guid string, name string, encPayload []byte, clientTimestamp time.Time) error {
	tag, err := dao.DB.Exec(ctx, "update secrets set name = $1, enc_payload = $2, updated_at = $3 where user_id = $4 and guid = $5", name, encPayload, clientTimestamp, userID, guid)
	if tag.RowsAffected() == 0 {
		return contract.ErrSecretNotFound
	}
	return err
}

// Delete - Удалить секрет
func (dao *SecretDAO) Delete(ctx context.Context, userID int64, guid string) error {
	tag, err := dao.DB.Exec(ctx, "delete from secrets where user_id = $1 AND guid = $1", userID, guid)
	if tag.RowsAffected() == 0 {
		return contract.ErrSecretNotFound
	}
	return err
}
