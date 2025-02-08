package controllers

import (
	"context"
	"net/http"
	"oms/database"
	"oms/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

// ViewOrder retrieves all orders from the database
func ViewOrder(ctx *gin.Context) {
	// Connect to MongoDB collection
	collection :=database.GetCollection("orders")

	// Query all orders
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch orders"})
		return
	}
	defer cursor.Close(context.Background())

	// Iterate through the results and store them in a slice
	var orders []models.Order
	if err := cursor.All(context.Background(), &orders); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error parsing orders"})
		return
	}

	// Return orders as JSON
	ctx.JSON(http.StatusOK, gin.H{
		"message": "YOUR ORDERS:",
		"orders":  orders,
	})
}
