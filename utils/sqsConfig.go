package utils

import (
	"context"
	"fmt"
	"log"
	"os"

	
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/omniful/go_commons/sqs"
	"github.com/omniful/go_commons/compression"
)

func SQSInitialization() {
	region := os.Getenv("AWS_REGION") 

	if region == "" {
		log.Fatalf("Missing AWS region environment variable")
	}

	// Load the AWS configuration using the default credentials chain
	_, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion(region),
	)
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	// Create FIFO Queue connection
	config := &sqs.Config{
		Account:   "062260873674", // Your AWS account number
		Region:    region,
		Compression: compression.None,
	}

	queueName := "MyOrders.fifo"  // The name of your FIFO queue
	queue, err := sqs.NewFifoQueue(context.Background(), queueName, config)
	if err != nil {
		log.Fatalf("Error connecting to FIFO queue: %v", err)
	}

	// Confirm successful queue connection
	fmt.Printf(" Successfully connected to FIFO Queue: %s\nQueue URL: %s\n", queue.Name, *queue.Url)
}
