package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)



func ViewOrder(ctx *gin.Context){
	ctx.JSON(http.StatusAccepted,gin.H{"MESSAGE":"YOUR ORDER IS :"})
}