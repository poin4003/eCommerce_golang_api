package initalize

import (
	"github.com/poin4003/eCommerce_golang_api/global"
	"github.com/segmentio/kafka-go"
	"log"
)

// Init Kafka Producer
var KafkaProducer *kafka.Writer

func InitKafka() {
	global.KafkaProducer = &kafka.Writer{
		Addr:     kafka.TCP("localhost:9092"),
		Topic:    "otp-auth-topic",
		Balancer: &kafka.LeastBytes{},
	}
}

func CloseKafka() {
	if err := global.KafkaProducer.Close(); err != nil {
		log.Fatal(err)
	}
}
