package global

import (
	"database/sql"
	"github.com/poin4003/eCommerce_golang_api/pkg/logger"
	"github.com/poin4003/eCommerce_golang_api/pkg/settings"
	"github.com/redis/go-redis/v9"
)

var (
	Config settings.Config
	Logger *logger.LoggerZap
	Rdb    *redis.Client
	Mdbc   *sql.DB
)
