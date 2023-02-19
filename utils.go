// utils.go
package utils

import (
	"os/exec"
	"strings"
)

func CreateTunnel(localIP string, tunnelName string, remoteIP string) error {
	cmd := exec.Command("ip", "tunnel", "add", tunnelName, "mode", "gre", "local", localIP, "remote", remoteIP)
	return cmd.Run()
}

func DeleteTunnel(tunnelName string) error {
	cmd := exec.Command("ip", "tunnel", "del", tunnelName)
	return cmd.Run()
}

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
