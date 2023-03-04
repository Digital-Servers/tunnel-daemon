// utils.go
package utils

import (
	"os/exec"
	"strings"
)

// CreateTunnel creates a new tunnel with the given local IP, tunnel name, and remote IP.
// It runs the corresponding Linux command to create the tunnel.
// If the command returns an error, it is returned from the function.
func CreateTunnel(localIP string, tunnelName string, remoteIP string) error {
	// Execute the "ip tunnel add" command with the specified parameters
	cmd := exec.Command("ip", "tunnel", "add", tunnelName, "mode", "gre", "local", localIP, "remote", remoteIP)
	// Run the command and return any errors
	return cmd.Run()
 }

// DeleteTunnel deletes the tunnel with the given name.
// It runs the corresponding Linux command to delete the tunnel.
// If the command returns an error, it is returned from the function.
func DeleteTunnel(tunnelName string) error {
	// Execute the "ip tunnel del" command with the specified tunnel name
	cmd := exec.Command("ip", "tunnel", "del", tunnelName)
	// Run the command and return any errors
	return cmd.Run()
 }

// ListTunnels returns a map of all existing tunnels, where the keys are the tunnel names and the values are the peer IP addresses.
// It runs the corresponding Linux command to list all the tunnels and processes the output to build the map.
// If the command returns an error, it is returned from the function.
// Note: This function excludes any line containing "gre0" from the result map.
func ListTunnels() (map[string]string, error) {
	// Execute the "ip tunnel show" command
	cmd := exec.Command("ip", "tunnel", "show")
	output, err := cmd.Output()
	if err != nil {
	   // If the command fails, return an error
	   return nil, err
	}
 
	// Create an empty map to store the tunnel names and peer IP addresses
	tunnels := make(map[string]string)
 
	// Iterate over each line of the command output
	for _, line := range strings.Split(string(output), "\n") {
	   // If the line contains "gre/ip" and does not contain "gre0"
	   if strings.Contains(line, "gre/ip") && !strings.Contains(line, "gre0") {
		  // Split the line into fields and extract the tunnel name and peer IP address
		  fields := strings.Fields(line)
		  name := strings.TrimSuffix(fields[0], ":")
		  peer := fields[3]
		  // Add the tunnel name and peer IP address to the map
		  tunnels[name] = peer
	   }
	}
 
	// Return the map of tunnel names and peer IP addresses
	return tunnels, nil
 }