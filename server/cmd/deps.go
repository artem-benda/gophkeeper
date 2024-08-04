package main

import (
	"encoding/base64"
	"github.com/artem-benda/gophkeeper/server/internal/domain/contract"
	"github.com/artem-benda/gophkeeper/server/internal/domain/service"
	"github.com/artem-benda/gophkeeper/server/internal/infrastructure/db"
	"github.com/artem-benda/gophkeeper/server/internal/infrastructure/repository"
	"github.com/jackc/pgx/v5/pgxpool"
)

type appDependencies struct {
	userService   contract.UserService
	secretService contract.SecretService
}

func mustCreateAppDependencies(dbPool *pgxpool.Pool, cfg Config) *appDependencies {
	userDAO := db.UserDAO{DB: dbPool}
	userRepo := repository.NewUserRepository(userDAO)
	userService := service.NewUserService(userRepo, cfg.mustGetSalt())
	secretDAO := db.SecretDAO{DB: dbPool}
	secretRepo := repository.NewSecretRepository(secretDAO)
	secretService := service.NewSecretService(secretRepo)
	return &appDependencies{userService, secretService}
}

func (c Config) mustGetSalt() []byte {
	salt, err := base64.StdEncoding.DecodeString(c.Salt)
	if err != nil {
		panic(err)
	}

	return salt
}
