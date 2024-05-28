package mq

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"tiny-tiktok/video_service/config"
	"tiny-tiktok/video_service/pkg/logger"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Printf("%s: %s", msg, err)
	}
}

func InitMQ() *amqp.Connection {
	// 连接到RabbitMQ服务器
	url := config.InitRabbitMQUrl()
	conn, err := amqp.Dial(url)
	logger.Log.Info("url:", url)
	failOnError(err, "Failed to connect to RabbitMQ")

	return conn
}

// ConsumeMessage 消费消息
func ConsumeMessage() {

}
