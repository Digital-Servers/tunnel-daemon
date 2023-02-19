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
	cmd := exec.Command("ip", "tunnel", "add", tunnelName, "mode", "gre", "local", localIP, "remote", remoteIP)
	return cmd.Run()
}

// DeleteTunnel deletes the tunnel with the given name.
// It runs the corresponding Linux command to delete the tunnel.
// If the command returns an error, it is returned from the function.
func DeleteTunnel(tunnelName string) error {
	cmd := exec.Command("ip", "tunnel", "del", tunnelName)
	return cmd.Run()
}

// ListTunnels returns a map of all existing tunnels, where the keys are the tunnel names and the values are the peer IP addresses.
// It runs the corresponding Linux command to list all the tunnels and processes the output to build the map.
// If the command returns an error, it is returned from the function.
// Note: This function excludes any line containing "gre0" from the result map.
func ListTunnels() (map[string]string, error) {
	cmd := exec.Command("ip", "tunnel", "show")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	tunnels := make(map[string]string)

	for _, line := range strings.Split(string(output), "\n") {
		if strings.Contains(line, "gre/ip") && !strings.Contains(line, "gre0") {
			fields := strings.Fields(line)
			name := strings.TrimSuffix(fields[0], ":")
			peer := fields[3]
			tunnels[name] = peer
		}
	}

	return tunnels, nil
}
