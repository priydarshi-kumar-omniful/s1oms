// this package contain the routes for creating and fetching the orders details

package routes

import (
	"oms/controllers"
	"oms/internal"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/omniful/go_commons/http"
)

type Response struct {
	Message string `json:"message"`
}

func IncomingRoutes(r *http.Server){
	r.POST("/createorder",controllers.BulkOrder)
	r.GET("/vieworder",controllers.ViewOrder)
	r.GET("/internal",func(ctx *gin.Context){
		var response Response
    result, err := internal.GetReq(context.Background(), &response, "/get/hub")
    
    if err != nil {
        ctx.JSON(500, gin.H{"error": err.Message})
        return
    }

    ctx.JSON(200, result) // Now returns the correct Response struct
	})
}


