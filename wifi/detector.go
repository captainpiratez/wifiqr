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
		trimmedLine := strings.TrimSpace(line)
		// Match exactly "SSID : <name>" to avoid false positives
		if strings.HasPrefix(trimmedLine, "SSID") && !strings.HasPrefix(trimmedLine, "BSSID") {
			parts := strings.SplitN(line, ":", 2)
			if len(parts) == 2 {
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
	output, err := cmd.CombinedOutput()
	if err != nil {
		message := strings.ToLower(string(output))
		if strings.Contains(message, "access is denied") {
			return "", fmt.Errorf("failed to retrieve WiFi password: access denied (try running as administrator)")
		}
		return "", fmt.Errorf("failed to retrieve WiFi password: %w", err)
	}

	lines := strings.Split(string(output), "\n")
	securityKeyAbsent := false
	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)
		if strings.HasPrefix(trimmedLine, "Security key") && strings.Contains(trimmedLine, "Absent") {
			securityKeyAbsent = true
		}
		if strings.HasPrefix(trimmedLine, "Authentication") && strings.Contains(strings.ToLower(trimmedLine), "open") {
			securityKeyAbsent = true
		}
		if strings.HasPrefix(trimmedLine, "Key Content") {
			parts := strings.SplitN(trimmedLine, ":", 2)
			if len(parts) == 2 {
				password := strings.TrimSpace(parts[1])
				return password, nil
			}
		}
	}

	if securityKeyAbsent {
		return "", nil
	}

	return "", fmt.Errorf("could not extract password for SSID: %s", ssid)
}
