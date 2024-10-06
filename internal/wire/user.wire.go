//go:build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/poin4003/eCommerce_golang_api/internal/controller"
	"github.com/poin4003/eCommerce_golang_api/internal/service"
)

func InitUserRouterHandler() (*controller.UserController, error) {
	wire.Build(
		repo.NewUserRepository,
		service.NewUserService,
		controller.NewUserController,
	)

	return new(controller.UserController), nil
}
