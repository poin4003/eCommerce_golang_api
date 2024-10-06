package initalize

import (
	"fmt"

	"github.com/poin4003/eCommerce_golang_api/global"
	"go.uber.org/zap"
)

func Run() {
	// Load configuration
	LoadConfig()
	m := global.Config.Mysql
	fmt.Println("Loading configuration mysql", m.Username, m.Password)
	InitLogger()
	global.Logger.Info("Config log ok!!", zap.String("ok", "success"))
	InitRedis()
	InitMysqlC()
	InitServiceInterface()

	r := InitRouter()

	r.Run(":8000")
}
