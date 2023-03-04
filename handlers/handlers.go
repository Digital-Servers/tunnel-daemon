// handlers.go
package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/Digital-Servers/tunnel-daemon/utils"
)

type VersionResponse struct {
	Version      string    `json:"version"`
	Uptime       time.Time `json:"uptime"`
	ResponseTime int64     `json:"response_time"`
}

var startTime = time.Now().UTC()

func authMiddleware(authToken string) gin.HandlerFunc {
	return func(c *gin.Context) {
	   authHeader := c.GetHeader("Authorization")
	   if authHeader != fmt.Sprintf("Bearer %s", authToken) {
		  c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization token"})
		  return
	   }
	   c.Next()
	}
 }

// CreateTunnel is the handler function for creating a new tunnel.
// It extracts the necessary parameters from the request and calls the corresponding utility function to create the tunnel.
// If there is an error creating the tunnel, it returns an HTTP 500 Internal Server Error response.
func CreateTunnel(c *gin.Context) {
	localIP := c.PostForm("localIP")
	tunnelName := c.PostForm("tunnelName")
	remoteIP := c.PostForm("remoteIP")

	if localIP == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "localIP is required"})
		return
	}

	if tunnelName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "tunnelName is required"})
		return
	}

	if remoteIP == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "remoteIP is required"})
		return
	}

	err := utils.CreateTunnel(localIP, tunnelName, remoteIP)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to create tunnel: %v", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "tunnel created successfully"})
}

// DeleteTunnel is the handler function for deleting an existing tunnel.
// It extracts the tunnel name from the request and calls the corresponding utility function to delete the tunnel.
// If there is an error deleting the tunnel, it returns an HTTP 500 Internal Server Error response.
func DeleteTunnel(c *gin.Context) {
	tunnelName := c.Param("name")

	if tunnelName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "tunnelName is required"})
		return
	}

	err := utils.DeleteTunnel(tunnelName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to delete tunnel: %v", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "tunnel deleted successfully"})
}

// GetTunnels is the handler function for listing all existing tunnels.
// It calls the corresponding utility function to get a list of all the tunnels and returns them in the response.
// If there is an error getting the list of tunnels, it returns an HTTP 500 Internal Server Error response.
func GetTunnels(c *gin.Context) {
	tunnels, err := utils.ListTunnels()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to list tunnels: %v", err)})
		return
	}

	c.JSON(http.StatusOK, tunnels)
}

// GetVersion is the handler function for the version endpoint.
// It retrieves the version information for the application and returns it as a JSON response.
func GetVersion(c *gin.Context, appVersion String) {
	start := time.Now()

	response := VersionResponse{
			Version:      appVersion,
			Uptime:       startTime,
			ResponseTime: time.Since(start).Nanoseconds() / int64(time.Millisecond),
	}

	c.JSON(http.StatusOK, response)
}