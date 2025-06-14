package global

import (
	"database/sql"
	"github.com/poin4003/eCommerce_golang_api/pkg/logger"
	"github.com/poin4003/eCommerce_golang_api/pkg/settings"
	"github.com/redis/go-redis/v9"
	"github.com/segmentio/kafka-go"
)

var (
	Config        settings.Config
	Logger        *logger.LoggerZap
	Rdb           *redis.Client
	Mdbc          *sql.DB
	KafkaProducer *kafka.Writer
)
