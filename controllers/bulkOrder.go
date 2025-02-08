package controllers

import (
	_ "context"
	
	"net/http"
	"fmt"
	
	"oms/producer"
	"github.com/gin-gonic/gin"
	_ "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var MongoClient *mongo.Client

// BulkOrder handles the order creation and generates the order ID
func BulkOrder(ctx *gin.Context) {
	
	
	if err:=producer.PublishOrderMessage("D:\\impfolder\\Desktop\\finalproject\\s1oms\\service\\order.csv");err!= nil {
		fmt.Println("Error publishing order message:", err)
	} else {
		fmt.Println("Message sent successfully!")
	}
	// Respond with the order ID (and any other info you want to return)
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Order created successfully"})

}
