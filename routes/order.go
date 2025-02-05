// this package contain the routes for creating and fetching the orders details

package routes

import (
	"github.com/omniful/go_commons/http"
	_"github.com/gin-gonic/gin"
	"oms/controllers"
)

func IncomingRoutes(r *http.Server){
	r.POST("/createorder",controllers.BulkOrder)
	r.GET("/vieweorder",controllers.ViewOrder)
}


