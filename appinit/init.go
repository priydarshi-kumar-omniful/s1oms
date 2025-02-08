package appinit

import (
	"context"
	"fmt"
	"oms/configs"
	"oms/consumer"
	"oms/controllers"
	"oms/database"
	"oms/internal"
	"oms/kafka"
)

func Initialize(ctx context.Context) {
	//database initialization
	database.Connect(ctx)
	controllers.MongoClient = database.Client

	// redis initialization
	redisClient := configs.ConnectToRedis(ctx)
	if redisClient == nil {
		fmt.Println("Redis connection failed. Exiting...")
		return
	}

	// sqs initialization
	configs.SQSInitialization()
	go consumer.StartConsumer()

	// internal service client initialization
	internal.InitInterSrvClient()

	// kafka initialization
	kafka.InitializeKafkaProducer()

	// Initialize Kafka Consumer
	go kafka.InitializeKafkaConsumer(ctx)
}