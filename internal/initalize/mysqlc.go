package initalize

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/poin4003/eCommerce_golang_api/global"
	"go.uber.org/zap"
)

func InitMysqlC() {
	m := global.Config.Mysql

	dsn := "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	var s = fmt.Sprintf(dsn, m.Username, m.Password, m.Host, m.Port, m.Dbname)

	db, err := sql.Open("mysql", s)

	checkErrorPanic(err, "InitMySql initialization error")

	global.Logger.Info("Initializing MySQL Successfully")

	global.Mdbc = db

	SetPool()
}

func checkErrorPanic(err error, errString string) {
	if err != nil {
		global.Logger.Error(errString, zap.Error(err))
		panic(err)
	}
}

func SetPool() {
	m := global.Config.Mysql

	sqlDb := global.Mdbc

	sqlDb.SetConnMaxIdleTime(time.Duration(m.MaxIdleConns))
	sqlDb.SetMaxOpenConns(m.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(time.Duration(m.ConnMaxLifetime))
}
