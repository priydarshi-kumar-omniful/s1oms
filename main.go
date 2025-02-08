package main

import (
	"fmt"
	"oms/appinit"
	"oms/constants" //project level constants
	"oms/routes"
	
	"github.com/omniful/go_commons/http"
	"github.com/omniful/go_commons/config"
	"log"
	"time"
	"os"
)

func main() {

	// Initialize config
	os.Setenv("CONFIG_PATH", "local")

	err := config.Init(time.Second * 10)
	if err != nil {
		log.Panicf("Error while initialising config, err: %v", err)
		panic(err)
	}

	ctx, err := config.TODOContext()
	if err != nil {
		log.Panicf("Error while getting context from config, err: %v", err)
		panic(err)
	}

	appinit.Initialize(ctx)

	server := http.InitializeServer(constants.PORT, constants.ReadTimeout, constants.WriteTimeout, constants.IdleTimeout)


	//checking the error
	if server == nil {
		fmt.Println("Failed to initialize server")
		return
	}

	routes.IncomingRoutes(server)

	// Start the server
	if err :=server.StartServer(config.GetString(ctx,"server.port"));err!= nil {
		fmt.Println("Server error:", err)
	}

}
