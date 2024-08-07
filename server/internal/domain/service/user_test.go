package service

import (
	"context"
	"encoding/base64"
	"errors"
	"testing"

	"github.com/artem-benda/gophkeeper/server/internal/domain/contract"
	"github.com/artem-benda/gophkeeper/server/internal/domain/entity"
	mock_contract "github.com/artem-benda/gophkeeper/server/internal/testing/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	assert := assert.New(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock_contract.NewMockUserRepository(ctrl)
	id := int64(123)
	err := errors.New("error")
	m.EXPECT().Register(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).Return(&id, err)

	svc := NewUserService(m, []byte("test salt"))
	idt, errt := svc.Register(context.Background(), "login", "passwordHash")

	assert.Equal(&id, idt, "id is not the same as repo returns while should be")
	assert.Equal(err, errt, "error is not the same as repo returns while should be")
}

func TestLogin_UserNotFound(t *testing.T) {
	assert := assert.New(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock_contract.NewMockUserRepository(ctrl)
	user := entity.User{ID: 123, Login: "login", PasswordHash: "passwordHash"}
	err := errors.New("error")
	m.EXPECT().GetUserByLogin(
		gomock.Any(),
		gomock.Any(),
	).Return(&user, err)

	svc := NewUserService(m, []byte("test salt"))
	idt, errt := svc.Login(context.Background(), "login", "password")

	assert.Nil(idt, "id is not nil while should be")
	assert.Equal(err, errt, "error is not the same as repo returns while should be")
}

func TestLogin_Successful(t *testing.T) {
	assert := assert.New(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock_contract.NewMockUserRepository(ctrl)
	user := entity.User{ID: 123, Login: "testuser123", PasswordHash: "36efaf1bb59c69dee2225be7a7f89fac3d314259f6c8e49e2516818fbabc7fde"}
	var err error = nil
	m.EXPECT().GetUserByLogin(
		gomock.Any(),
		gomock.Any(),
	).Return(&user, err)

	saltBase64 := "BPjkLEqJfARvsYGW++WRcnCjxHyZsrnxXd/qdzpWIaE="
	salt, _ := base64.StdEncoding.DecodeString(saltBase64)
	svc := NewUserService(m, salt)
	idt, errt := svc.Login(context.Background(), "testuser", "testuser123")

	assert.Equal(&user.ID, idt, "user is not the same as repo returns while should be")
	assert.Nil(errt, "error is not nil while should be: %s", errt)
}

func TestLogin_WrongPassword(t *testing.T) {
	assert := assert.New(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock_contract.NewMockUserRepository(ctrl)
	user := entity.User{ID: 123, Login: "testuser123", PasswordHash: "36efaf1bb59c69dee2225be7a7f89fac3d314259f6c8e49e2516818fbabc7fde"}
	var err error = nil
	m.EXPECT().GetUserByLogin(
		gomock.Any(),
		gomock.Any(),
	).Return(&user, err)

	saltBase64 := "BPjkLEqJfARvsYGW++WRcnCjxHyZsrnxXd/qdzpWIaE="
	salt, _ := base64.StdEncoding.DecodeString(saltBase64)
	svc := NewUserService(m, salt)
	idt, errt := svc.Login(context.Background(), "testuser", "testuser1234")

	assert.Nil(idt, "user is not nil while should be")
	assert.Equal(contract.ErrUnauthorized, errt, "error is not ErrUnauthorized while should be")
}
