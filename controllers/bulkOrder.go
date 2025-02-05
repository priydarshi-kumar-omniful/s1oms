package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
func BulkOrder(ctx *gin.Context){
	ctx.JSON(http.StatusAccepted,gin.H{"MESSAGE":"ORDER CREATED"})
}

