package main

import(
	"fmt"
	"github.com/omniful/go_commons/http"
	"github.com/gin-gonic/gin"
	"time"
)

func main(){
	
	// Initialize the server
	server := http.InitializeServer(":8080", 0, 0, 0)

	//checking the error
	if server == nil {
		fmt.Println("Failed to initialize server")
		return
	}
	
	server.GET("/",func(ctx *gin.Context){
		ctx.JSON(int(http.StatusOK),gin.H{"MESSAGE":"ACCEPTED REQUES"})
	})

	// Start the server
	err := server.StartServer("oms")
	if err != nil {
		fmt.Println("Server error:", err)
	}

}











