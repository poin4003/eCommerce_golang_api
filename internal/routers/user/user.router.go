package user

import (
	"github.com/gin-gonic/gin"
	"github.com/poin4003/eCommerce_golang_api/internal/controller/account"
	"github.com/poin4003/eCommerce_golang_api/internal/middlewares"
	//"github.com/poin4003/eCommerce_golang_api/internal/wire"
)

type UserRouter struct{}

func (pr *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	// Public router
	userRouterPublic := Router.Group("/user")
	{
		userRouterPublic.POST("/register", account.Login.Register)
		userRouterPublic.POST("/verify_account", account.Login.VerifyOTP)
		userRouterPublic.POST("/update_pass_register", account.Login.UpdatePasswordRegister)
		userRouterPublic.POST("/login", account.Login.Login)
	}

	// Private router
	userRouterPrivate := Router.Group("/user")
	userRouterPrivate.Use(middlewares.AuthenMiddleware())
	{
		userRouterPrivate.GET("/get_info")
		userRouterPrivate.POST("/two_factor/setup", account.TwoFA.SetupTwoFactorAuth)
	}
}
