package manager

import "github.com/gin-gonic/gin"

type AdminRouter struct{}

func (ar *AdminRouter) InitAdminRouter(Router *gin.RouterGroup) {
	// Public router
	//adminRouterPublic := Router.Group("/admin")
	//{
	//	adminRouterPublic.POST("/login")
	//}

	// Private
	//adminRouterPrivate := Router.Group("/admin")
	//adminRouterPrivate.Use(limiter())
	//adminRouterPrivate.Use(Authen())
	//adminRouterPrivate.Use(Permission())
	//{
	//	adminRouterPrivate.POST("/active_user")
	//}
}
