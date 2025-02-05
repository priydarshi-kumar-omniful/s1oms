package main

import (
	"fmt"
	"oms/routes"
	"github.com/omniful/go_commons/http"
)

func main(){
	
	// Initialize the server
	server := http.InitializeServer(":8080", 0, 0, 0)

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











