package main

import (
	"fmt"
	"oms/routes"
	"oms/constants" //project level constants
	"oms/database"
	"github.com/omniful/go_commons/http"
	"oms/utils"

)

func main(){
	// Initialize the server
	server := http.InitializeServer(constants.PORT, constants.ReadTimeout, constants.WriteTimeout, constants.IdleTimeout)
	
	database.Connect()

	redisClient := utils.ConnectToRedis()
	if redisClient == nil {
		fmt.Println("❌ Redis connection failed. Exiting...")
		return
	}

	//checking the error
	if server == nil {
		fmt.Println("Failed to initialize server")
		return
	}
	
	routes.IncomingRoutes(server)

	// Start the server
	err := server.StartServer("oms")
	if err != nil {
		fmt.Println("Server error:", err)
	}

}











