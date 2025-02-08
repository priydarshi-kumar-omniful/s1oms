package repo

import (
	"context"
	"log"
	"oms/database"
	"oms/models"


	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateOrder inserts multiple orders into MongoDB efficiently
func CreateOrder(orders []models.Order) error {
	if len(orders) == 0 {
		log.Println("No orders to insert.")
		return nil
	}

	collection := database.GetCollection("orders")

	// Prepare orders for insertion
	var insertDocs []interface{}
	for i := range orders {
		orders[i].ID = primitive.NewObjectID() // Assign a unique ID
		orders[i].TotalAmount = calculateTotalAmount(orders[i].Items)
		insertDocs = append(insertDocs, orders[i])
	}

	// Perform bulk insertion
	result, err := collection.InsertMany(context.Background(), insertDocs)
	if err != nil {
		log.Printf("Error inserting orders: %v", err)
		return err
	}

	log.Printf("Successfully inserted %d orders", len(result.InsertedIDs))
	return nil
}

// calculateTotalAmount computes the total price of an order
func calculateTotalAmount(items []models.OrderItem) float64 {
	var total float64
	for _, item := range items {
		total += float64(item.Quantity) * item.Price
	}
	return total
}
