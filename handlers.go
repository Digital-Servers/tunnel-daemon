// handlers.go
package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/Digital-Servers/tunnel-daemon/utils"
)

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
