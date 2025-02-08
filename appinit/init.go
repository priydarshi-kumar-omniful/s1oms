package appinit

import (
	"context"
	"fmt"
	"oms/configs"
	"oms/consumer"
	"oms/controllers"
	"oms/database"
	"oms/internal"
)

func Initialize(ctx context.Context) {
	//database initialization
	database.Connect(ctx)
	controllers.MongoClient = database.Client

	//redis initialization
	redisClient := configs.ConnectToRedis(ctx)
	if redisClient == nil {
		fmt.Println("Redis connection failed. Exiting...")
		return
	}

	//sqs initialization
	configs.SQSInitialization()
	go consumer.StartConsumer()


	internal.InitInterSrvClient()
}
