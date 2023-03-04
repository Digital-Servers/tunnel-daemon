// main.go
package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"github.com/Digital-Servers/tunnel-daemon/handlers"
)

const authToken = "YOUR-BEARER-TOKEN"
const appVersion = "0.0.1"

func main() {
	// Create a new default Gin router
	r := gin.Default()

	// Map the "/api/tunnel" POST route to the CreateTunnel function in handlers.go
	r.POST("/api/tunnel", handlers.CreateTunnel)

	// Map the "/api/tunnel/:name" DELETE route to the DeleteTunnel function in handlers.go
	r.DELETE("/api/tunnel/:name", handlers.DeleteTunnel)

	// Map the "/api/tunnels" GET route to the GetTunnels function in handlers.go
	r.GET("/api/tunnels", handlers.GetTunnels)

	// Map the "/api/version" GET route to the GetVersion function in handlers.go
	r.GET("/api/version", handlers.GetVersion)

	// Start the server on port 8080
	err := r.Run(":8080")
	if err != nil {
		log.Fatal("Unable to start server:", err)
	}
}
