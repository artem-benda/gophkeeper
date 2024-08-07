package service

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/artem-benda/gophkeeper/server/internal/domain/entity"
	mock_contract "github.com/artem-benda/gophkeeper/server/internal/testing/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestAdd_WithSuccess(t *testing.T) {
	assert := assert.New(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock_contract.NewMockSecretRepository(ctrl)
	id := int64(123)
	var err error
	m.EXPECT().Insert(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).Return(&id, err)

	svc := NewSecretService(m)
	guid, errt := svc.Add(context.Background(), 1, "test", []byte("test payload"))

	assert.NotEmpty(guid, "guid should be not empty when repo executes successfully")
	assert.Nil(errt, "error is not nil while should be")
}

func TestAdd_WithError(t *testing.T) {
	assert := assert.New(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock_contract.NewMockSecretRepository(ctrl)
	id := int64(123)
	err := errors.New("error")
	m.EXPECT().Insert(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).Return(&id, err)

	svc := NewSecretService(m)
	guid, errt := svc.Add(context.Background(), 1, "test", []byte("test payload"))

	assert.Equal("", guid, "guid should be empty when repo returns error")
	assert.Equal(err, errt, "error not returned while should be")
}

func TestEdit(t *testing.T) {
	assert := assert.New(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock_contract.NewMockSecretRepository(ctrl)
	err := errors.New("error")
	m.EXPECT().Update(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).Return(err)

	svc := NewSecretService(m)
	errt := svc.Edit(context.Background(), 1, "guid", "test", []byte("test payload"))

	assert.Equal(err, errt, "error is not the same as repo returns while should be")
}

func TestRemove(t *testing.T) {
	assert := assert.New(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock_contract.NewMockSecretRepository(ctrl)
	err := errors.New("error")
	m.EXPECT().Delete(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).Return(err)

	svc := NewSecretService(m)
	errt := svc.Remove(context.Background(), 1, "guid")

	assert.Equal(err, errt, "error is not the same as repo returns while should be")
}

func TestGet(t *testing.T) {
	assert := assert.New(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock_contract.NewMockSecretRepository(ctrl)
	s := &entity.Secret{ID: 1, GUID: "guid", Name: "name", EncPayload: []byte("payload"), CreatedAt: time.Time{}, UpdatedAt: time.Time{}}
	err := errors.New("error")
	m.EXPECT().Get(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).Return(s, err)

	svc := NewSecretService(m)
	st, errt := svc.Get(context.Background(), 1, "guid")

	assert.Equal(s, st, "secret is not the same as repo returns while should be")
	assert.Equal(err, errt, "error is not the same as repo returns while should be")
}

func TestGetAll(t *testing.T) {
	assert := assert.New(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock_contract.NewMockSecretRepository(ctrl)
	secrets := []entity.Secret{
		{ID: 1, GUID: "guid", Name: "name", EncPayload: []byte("payload"), CreatedAt: time.Time{}, UpdatedAt: time.Time{}},
		{ID: 2, GUID: "guid2", Name: "name2", EncPayload: []byte("payload2"), CreatedAt: time.Time{}, UpdatedAt: time.Time{}},
	}
	err := errors.New("error")
	m.EXPECT().GetByUserID(
		gomock.Any(),
		gomock.Any(),
	).Return(secrets, err)

	svc := NewSecretService(m)
	secretrst, errt := svc.GetByUserID(context.Background(), 1)

	assert.Equal(secrets, secretrst, "secrets sre not the same as repo returns while should be")
	assert.Equal(err, errt, "error is not the same as repo returns while should be")
}
