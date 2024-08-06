package main

import (
	"github.com/artem-benda/gphkeeper/client/internal/application"
	"github.com/artem-benda/gphkeeper/client/internal/domain/service"
)

func mustCreateAppDependencies() *application.AppDependencies {
	userService := service.User
	return &application.AppDependencies{
		US: userService,
	}
}
