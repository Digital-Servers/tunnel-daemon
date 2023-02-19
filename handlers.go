// handlers.go
package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/tunnel-api/utils"
)

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

func GetTunnels(c *gin.Context) {
	tunnels, err := utils.ListTunnels()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to list tunnels: %v", err)})
		return
	}

	c.JSON(http.StatusOK, tunnels)
}
