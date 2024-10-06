package manager

import "github.com/gin-gonic/gin"

type UserRouter struct{}

func (ur *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	// Public router

	//userRouterPublic := Router.Group("/admin/user")
	//{
	//	//userRouterPublic.POST("/register")
	//	//userRouterPublic.POST("/login")
	//}

	// Private router
	//userRouter := Router.Group("/admin/user")
	//{
	//	userRouter.POST("/active_user")
	//}

}
