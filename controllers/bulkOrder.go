package controllers

import (
	_ "context"
	"log"
	"net/http"
	
	csv_order "oms/orderCSV"

	"github.com/gin-gonic/gin"
	_ "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var MongoClient *mongo.Client

// BulkOrder handles the order creation and generates the order ID
func BulkOrder(ctx *gin.Context) {
	csv_order.MongoClient = MongoClient
	err := csv_order.ParseAndCreateOrdersFromCSV()

	if err != nil {
		log.Printf("Error inserting order: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
		return
	}

	// Respond with the order ID (and any other info you want to return)
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Order created successfully"})

}
