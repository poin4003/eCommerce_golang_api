package initalize

import (
	"github.com/poin4003/eCommerce_golang_api/global"
	"github.com/poin4003/eCommerce_golang_api/internal/database"
	"github.com/poin4003/eCommerce_golang_api/internal/service"
	"github.com/poin4003/eCommerce_golang_api/internal/service/implement"
)

func InitServiceInterface() {
	queries := database.New(global.Mdbc)

	// User Service interface
	service.InitUserLogin(implement.NewUserLoginImplement(queries))
}
