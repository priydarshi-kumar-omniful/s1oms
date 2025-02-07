package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/omniful/go_commons/compression"
	"github.com/omniful/go_commons/sqs"
)

var (
	SQSQueue     *sqs.Queue
	SQSPublisher *sqs.Publisher // Store Publisher instance
	QueueURL     string
	Region       string
)

func SQSInitialization() {
	Region = os.Getenv("AWS_REGION")
	if Region == "" {
		log.Fatalf("Missing AWS region environment variable")
	}

	config := &sqs.Config{
		Account:     "062260873674",
		Region:      Region,
		Compression: compression.None,
	}

	queueName := "MyOrders.fifo"
	queue, err := sqs.NewFifoQueue(context.Background(), queueName, config)
	if err != nil {
		log.Fatalf("Error connecting to FIFO queue: %v", err)
	}

	// Store queue and publisher
	SQSQueue = queue
	SQSPublisher = sqs.NewPublisher(queue) // ✅ Initialize publisher
	QueueURL = *queue.Url
	fmt.Printf("Successfully connected to FIFO Queue: %s\nQueue URL: %s\n", queue.Name, *queue.Url)
}
