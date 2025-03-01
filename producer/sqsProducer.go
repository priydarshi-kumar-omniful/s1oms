package producer

import (
	"context"
	"fmt"
	"log"

	"oms/configs"

	"github.com/google/uuid"
	"github.com/omniful/go_commons/sqs"
)

func PublishOrderMessage(orderFilePath string) error {
	if configs.SQSPublisher == nil {
		return fmt.Errorf("SQS publisher is not initialized")
	}

	message := &sqs.Message{
		Value:           []byte(fmt.Sprintf("OrderCSVFilePath: %s", orderFilePath)),
		GroupId:         "order-processing-group",
		DeduplicationId: uuid.New().String(), // Unique ID for FIFO queue
	}

	err := configs.SQSPublisher.Publish(context.Background(), message)
	if err != nil {
		log.Printf("Failed to publish order message: %v", err)
		return err
	}

	fmt.Println("Successfully published message:", orderFilePath)
	return nil
}
