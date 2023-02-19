// main.go
package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/Digital-Servers/tunnel-daemon/handlers"
)

func main() {
	r := gin.Default()

	r.POST("/tunnel", handlers.CreateTunnel)
	r.DELETE("/tunnel/:name", handlers.DeleteTunnel)
	r.GET("/tunnels", handlers.GetTunnels)

	err := r.Run(":8080")
	if err != nil {
		log.Fatal("Unable to start server:", err)
	}
}
