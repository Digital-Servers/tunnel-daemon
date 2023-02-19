// main.go
package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"github.com/Digital-Servers/tunnel-daemon/handlers"
)

func main() {
	// Create a new default Gin router
	r := gin.Default()

	// Map the "/tunnel" POST route to the CreateTunnel function in handlers.go
	r.POST("/tunnel", handlers.CreateTunnel)

	// Map the "/tunnel/:name" DELETE route to the DeleteTunnel function in handlers.go
	r.DELETE("/tunnel/:name", handlers.DeleteTunnel)

	// Map the "/tunnels" GET route to the GetTunnels function in handlers.go
	r.GET("/tunnels", handlers.GetTunnels)

	// Start the server on port 8080
	err := r.Run(":8080")
	if err != nil {
		log.Fatal("Unable to start server:", err)
	}
}
