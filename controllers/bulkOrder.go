package controllers

import (
	_ "context"
	
	"net/http"
	"fmt"
	csv_order "oms/orderCSV"
	"oms/producer"
	"github.com/gin-gonic/gin"
	_ "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var MongoClient *mongo.Client

// BulkOrder handles the order creation and generates the order ID
func BulkOrder(ctx *gin.Context) {
	csv_order.MongoClient = MongoClient
	
	if err:=producer.PublishOrderMessage("./orderCSV/order.csv");err!= nil {
		fmt.Println("Error publishing order message:", err)
	} else {
		fmt.Println("Message sent successfully!")
	}
	// Respond with the order ID (and any other info you want to return)
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Order created successfully"})

}
