// package consumer

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"oms/config"
// 	csv_order "oms/orderCSV"
// 	"runtime/debug"
// 	"github.com/omniful/go_commons/sqs"
// )

// // OrderCSVHandler processes messages from the queue
// type OrderCSVHandler struct{}

// // Process handles incoming SQS messages
// func (h *OrderCSVHandler) Process(ctx context.Context, messages *[]sqs.Message) (err error) {
// 	defer func() {
// 		if r := recover(); r != nil {
// 			log.Printf(" PANIC Recovered in Process: %v\nStack trace: %s", r, debug.Stack())
// 			err = fmt.Errorf("panic recovered: %v", r)
// 		}
// 	}()

// 	// Check if messages are nil or empty
// 	if messages == nil || len(*messages) == 0 {
// 		log.Println(" No messages received, skipping processing.")
// 		return nil
// 	}

// 	// Log received messages
// 	log.Printf(" Received %d messages", len(*messages))

// 	for _, message := range *messages {

// 		// Extract file path from message
// 		filePath := string(message.Value)
// 		log.Printf(" Processing CSV file from path: %s", filePath)

// 		// Ensure filePath is valid
// 		if filePath == "" {
// 			log.Println("Received empty file path, skipping...")
// 			continue
// 		}

// 		// Process the CSV file
// 		err := csv_order.ParseAndCreateOrdersFromCSV(filePath)
// 		if err != nil {
// 			log.Printf(" Error processing CSV file: %v", err)
// 			return err
// 		}

// 		// Remove the processed message from the queue
// 		err = config.SQSQueue.Remove(message.ReceiptHandle)
// 		if err != nil {
// 			log.Printf("Failed to delete message from queue: %v", err)
// 			return err
// 		}

// 		log.Println(" CSV processing complete. Message deleted from queue.")
// 	}

// 	return nil
// }

// // StartConsumer initializes and starts the SQS consumer
// func StartConsumer() {
// 	if config.SQSQueue == nil {
// 		log.Fatal("SQS queue is not initialized")
// 	}

// 	handler := &OrderCSVHandler{}
// 	consumer, err := sqs.NewConsumer(
// 		config.SQSQueue, // Queue reference from config
// 		1,               // Number of workers (max 2)
// 		1,               // Concurrency per worker
// 		handler,         // Handler for processing messages
// 		1,               // Max messages count (FIFO queues allow only 1)
// 		30,              // Visibility timeout (in seconds)
// 		true,            // Async processing
// 		false,           // Disable batch processing for FIFO
// 	)

// 	if err != nil {
// 		log.Fatalf("Failed to create SQS consumer: %v", err)
// 	}

// 	// Start consumer
// 	consumer.Start(context.Background())
// 	fmt.Println("SQS Consumer started and listening for messages...")
// }

// -------------------------OWN CONSUMER IMPLEMENTATION----------------

package consumer

import (
	"context"
	"fmt"
	"log"
	omsConfig "oms/config"
	"oms/orderCSV"
	"strings"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

type MessageOutput struct {
	Message string
}

func StartConsumer() {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(omsConfig.Region))
	if err != nil {
		log.Fatalf("failed to load AWS SDK config: %v", err)
	}

	sqsClient := sqs.NewFromConfig(cfg)

	for {
		output, err := sqsClient.ReceiveMessage(context.TODO(), &sqs.ReceiveMessageInput{
			QueueUrl:            aws.String(omsConfig.QueueURL),
			MaxNumberOfMessages: 10,
			WaitTimeSeconds:     5,
		})
		if err != nil {
			log.Printf("failed to receive messages: %v", err)
			continue
		}

		for _, msg := range output.Messages {
			messageOutput := MessageOutput{Message: *msg.Body}
			fmt.Printf("Received message: %s\n", messageOutput.Message)
			filePath := strings.TrimPrefix(messageOutput.Message, "OrderCSVFilePath: ")
			filePath = strings.TrimSpace(filePath) // Remove any extra spaces
			if err := orderCSV.ParseAndCreateOrdersFromCSV(filePath); err != nil {
				fmt.Println("error in creating order",err)
			}

			// Delete message after processing
			_, err = sqsClient.DeleteMessage(context.TODO(), &sqs.DeleteMessageInput{
				QueueUrl:      aws.String(omsConfig.QueueURL),
				ReceiptHandle: msg.ReceiptHandle,
			})
			if err != nil {
				log.Printf("failed to delete message: %v", err)
			} else {
				fmt.Println("Message deleted successfully")
			}
		}
	}
}
