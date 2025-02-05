package routes

import (
	"github.com/gin-gonic/gin"
	"oms/controllers"
)

func IncomingRoutes(r *gin.Engine){
	r.POST("/createorder",controllers.BulkOrder)
}