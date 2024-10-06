package initalize

import (
	"github.com/poin4003/eCommerce_golang_api/global"
	"github.com/poin4003/eCommerce_golang_api/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}
