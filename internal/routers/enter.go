package routers

import (
	"github.com/poin4003/eCommerce_golang_api/internal/routers/manager"
	"github.com/poin4003/eCommerce_golang_api/internal/routers/user"
)

type RouterGroup struct {
	User    user.UserRouterGroup
	Manager manager.ManagerRouterGroup
}

var RouterGroupApp = new(RouterGroup)
