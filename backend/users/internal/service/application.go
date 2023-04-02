package service

import (
	"github.com/dstopka/notebook-app/backend/users/internal/adapters"
	"github.com/dstopka/notebook-app/backend/users/internal/app"
)

func NewApplication() *app.Application {
	userRepository := adapters.NewMemoryUserRepository()

	app := app.New(userRepository)
	return app
}
