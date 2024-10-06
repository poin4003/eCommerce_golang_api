package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
	"time"

	"github.com/segmentio/kafka-go"
)

var (
	kafkaConsumer *kafka.Reader
	kafkaProducer *kafka.Writer
)

const (
	kafkaURL   = "127.0.0.1:9092"
	kafkaTopic = "user_topic_vip"
)

// For producer
func getKafkaWriter(kafkaURL, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(kafkaURL),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
}

// For consumer
func getKafkaReader(kafkaURL, topic, groupID string) *kafka.Reader {
	brokers := strings.Split(kafkaURL, ",")
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:        brokers,
		GroupID:        groupID,
		Topic:          topic,
		MinBytes:       10e3,
		MaxBytes:       10e6,
		CommitInterval: time.Second,
		StartOffset:    kafka.LastOffset,
	})
}

type StockInfo struct {
	Message string `json:"message"`
	Type    string `json:"type"`
}

func newStock(msg, typeMsg string) *StockInfo {
	s := StockInfo{}
	s.Message = msg
	s.Type = typeMsg

	return &s
}

func actionStock(c *gin.Context) {
	s := newStock(c.Query("msg"), c.Query("type"))
	body := map[string]interface{}{
		"action": "action",
		"info":   s,
	}

	jsonBody, _ := json.Marshal(body)

	msg := kafka.Message{
		Key:   []byte("action"),
		Value: []byte(jsonBody),
	}

	if err := kafkaProducer.WriteMessages(context.Background(), msg); err != nil {
		fmt.Printf("Error sending message: %v\n", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"msg": "action Successfully"})
}

func RegisterConsumerATC(id int) {
	kafkaGroupId := fmt.Sprintf("consumer-group-%d", id)
	reader := getKafkaReader(kafkaURL, kafkaTopic, kafkaGroupId)
	defer reader.Close()

	fmt.Printf("Consumer(%d) connected to topic: %s\n", id, kafkaTopic)

	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			fmt.Printf("Consumer(%d) error: %v\n", id, err)
			continue
		}
		fmt.Printf("Consumer(%d), topic:%v, partition:%v, offset:%v, time:%d %s = %s\n",
			id, m.Topic, m.Partition, m.Offset, m.Time.Unix(), string(m.Key), string(m.Value))
	}
}

func main() {
	r := gin.Default()
	kafkaProducer = getKafkaWriter(kafkaURL, kafkaTopic)
	defer kafkaProducer.Close()

	r.POST("action/stock", actionStock)

	go RegisterConsumerATC(1)
	go RegisterConsumerATC(2)

	r.Run(":8999")
}
