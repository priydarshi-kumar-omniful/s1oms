package consumer

import (
	"context"
	"log"
	"oms/configs"
	"oms/service"
	"strings"

	"github.com/omniful/go_commons/sqs"
)

// OrderMessageHandler processes incoming SQS messages.
type OrderMessageHandler struct{}

// Process handles the messages from SQS.
func (h *OrderMessageHandler) Process(ctx context.Context, messages *[]sqs.Message) (err error) {

	for _, message := range *messages {

		// Convert byte slice to string
		messageString := string(message.Value)
		log.Printf("[DEBUG] Raw message content: %s", messageString)

		// Extract file path
		filePath := strings.TrimSpace(strings.TrimPrefix(messageString, "OrderCSVFilePath: "))

		// Process the CSV file and create orders
		if err := service.ParseAndCreateOrdersFromCSV(filePath); err != nil {
			log.Printf("[ERROR] Failed to process order CSV: %v", err)
			return err // This ensures the message remains in the queue if processing fails
		}

		log.Println("[SUCCESS] Order processed successfully.")
	}
	return nil
}

// StartConsumer initializes and runs the go_commons SQS consumer
func StartConsumer() {
	if configs.SQSQueue == nil {
		log.Fatal("[FATAL] SQS queue is not initialized. Ensure SQSInitialization() is called first.")
	}

	// Initialize the SQS consumer
	consumer, err := sqs.NewConsumer(configs.SQSQueue, 2, 2, &OrderMessageHandler{}, 10, 30, true, false)
	if err != nil {
		log.Fatalf("[FATAL] Failed to create SQS consumer: %v", err)
	}

	log.Println("[INFO] Starting SQS Consumer for MyOrders.fifo...")
	consumer.Start(context.Background())

	// Keep the consumer running
	select {}
}
