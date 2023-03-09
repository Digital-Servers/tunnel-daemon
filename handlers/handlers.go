// handlers.go
package handlers

import (
	"fmt"
    "io/ioutil"
    "net/http"
    "os/exec"
	"time"
	"bytes"

	"github.com/gin-gonic/gin"
	"github.com/Digital-Servers/tunnel-daemon/utils"
)

type VersionResponse struct {
	Version      string    `json:"version"`
	Uptime       time.Time `json:"uptime"`
	ResponseTime int64     `json:"response_time"`
}

var startTime = time.Now().UTC()

// CreateTunnel is the handler function for creating a new tunnel.
// It extracts the necessary parameters from the request and calls the corresponding utility function to create the tunnel.
// If any of the required parameters are missing, it returns an HTTP 400 Bad Request response.
// If there is an error creating the tunnel, it returns an HTTP 500 Internal Server Error response.
func CreateTunnel(c *gin.Context) {
	// Get the local IP, tunnel name, and remote IP from the request parameters
	localIP := c.PostForm("localIP")
	tunnelName := c.PostForm("tunnelName")
	remoteIP := c.PostForm("remoteIP")
 
	// Check if any of the required parameters are missing
	if localIP == "" || tunnelName == "" || remoteIP == "" {
	   // If so, return a bad request error
	   c.JSON(http.StatusBadRequest, gin.H{"error": "localIP, tunnelName, and remoteIP are required"})
	   return
	}
 
	// Call the utility function to create the tunnel
	err := utils.CreateTunnel(localIP, tunnelName, remoteIP)
	// If there is an error, return an internal server error
	if err != nil {
	   c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to create tunnel: %v", err)})
	   return
	}
 
	// If the tunnel is created successfully, return a success message
	c.JSON(http.StatusOK, gin.H{"message": "tunnel created successfully"})
 }

// DeleteTunnel is the handler function for deleting an existing tunnel.
// It extracts the tunnel name from the request and calls the corresponding utility function to delete the tunnel.
// If the tunnel name is missing, it returns an HTTP 400 Bad Request response.
// If there is an error deleting the tunnel, it returns an HTTP 500 Internal Server Error response.
func DeleteTunnel(c *gin.Context) {
	// Get the tunnel name from the request parameters
	tunnelName := c.Param("name")
 
	// Check if the tunnel name is missing
	if tunnelName == "" {
	   // If so, return a bad request error
	   c.JSON(http.StatusBadRequest, gin.H{"error": "tunnelName is required"})
	   return
	}
 
	// Call the utility function to delete the tunnel
	err := utils.DeleteTunnel(tunnelName)
	// If there is an error, return an internal server error
	if err != nil {
	   c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to delete tunnel: %v", err)})
	   return
	}
 
	// If the tunnel is deleted successfully, return a success message
	c.JSON(http.StatusOK, gin.H{"message": "tunnel deleted successfully"})
 }

// GetTunnels is the handler function for listing all existing tunnels.
// It calls the corresponding utility function to get a list of all the tunnels and returns them in the response.
// If there is an error getting the list of tunnels, it returns an HTTP 500 Internal Server Error response.
func GetTunnels(c *gin.Context) {
	// Call the utility function to get a list of all the tunnels
	tunnels, err := utils.ListTunnels()
	// If there is an error, return an internal server error
	if err != nil {
	   c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to list tunnels: %v", err)})
	   return
	}
 
	// If the list of tunnels is retrieved successfully, return them in the response
	c.JSON(http.StatusOK, tunnels)
 }

// GetVersion is the handler function for the version endpoint.
// It retrieves the version information for the application and returns it as a JSON response.
func GetVersion(c *gin.Context, appVersion string) {
	// Get the current time to calculate the response time
	start := time.Now()
 
	// Create a new version response struct with the current app version, start time, and response time
	response := VersionResponse{
	   Version:      appVersion,
	   Uptime:       startTime,
	   ResponseTime: time.Since(start).Nanoseconds() / int64(time.Millisecond),
	}
 
	// Return the version response as a JSON response
	c.JSON(http.StatusOK, response)
 }

// SetupInternal sends an HTTP GET request to the specified URL with a bearer token, then executes the downloaded script.
func SetupInternal(url string, token string) {
	// Create a new HTTP request with the specified URL
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		// If there was an error creating the request, print the error and return
		fmt.Println(err)
		return
	}

	// Add the bearer token to the request headers
	req.Header.Set("Authorization", "Bearer "+token)

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		// If there was an error sending the request, print the error and return
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	// Read the script contents from the response body
	script, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// If there was an error reading the response body, print the error and return
		fmt.Println(err)
		return
	}

	// Execute the script
	cmd := exec.Command("sudo", "bash", "-c", string(script))
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err = cmd.Run()
	if err != nil {
		// If there was an error executing the script, print the stderr output and return
		fmt.Println(stderr.String())
		return
	}

	// Print the stdout output
	fmt.Println(stdout.String())
}