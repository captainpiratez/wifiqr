package wifi

import (
	"fmt"
	"os/exec"
	"strings"
)

// GetCurrentSSID retrieves the currently connected WiFi network name
func GetCurrentSSID() (string, error) {
	cmd := exec.Command("netsh", "wlan", "show", "interfaces")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to run netsh command: %w", err)
	}

	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if strings.Contains(line, "SSID") && !strings.Contains(line, "BSSID") {
			parts := strings.Split(line, ":")
			if len(parts) > 1 {
				ssid := strings.TrimSpace(parts[1])
				if ssid != "" {
					return ssid, nil
				}
			}
		}
	}

	return "", fmt.Errorf("no connected WiFi network found")
}

// GetPassword retrieves the WiFi password for a given SSID
func GetPassword(ssid string) (string, error) {
	cmd := exec.Command("netsh", "wlan", "show", "profile", ssid, "key=clear")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to retrieve WiFi password (requires admin privileges): %w", err)
	}

	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if strings.Contains(line, "Key Content") {
			parts := strings.Split(line, ":")
			if len(parts) > 1 {
				password := strings.TrimSpace(parts[1])
				return password, nil
			}
		}
	}

	return "", fmt.Errorf("could not extract password for SSID: %s", ssid)
}
