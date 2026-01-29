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
		// Match exactly "SSID :" field to avoid false positives
		if strings.HasPrefix(trimmedLine, "SSID") && strings.Contains(trimmedLine, ":") && !strings.HasPrefix(trimmedLine, "BSSID") {
			parts := strings.SplitN(trimmedLine, ":", 2)
			if len(parts) == 2 {
				ssid := strings.TrimSpace(parts[1])
				if ssid != "" && ssid != "None" {
					return ssid, nil
				}
			}
		}
	}

	return "", fmt.Errorf("no connected WiFi network found")
}

// GetSecurityType detects the WiFi security type from netsh output
func GetSecurityType(output []byte) string {
	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)
		if strings.HasPrefix(trimmedLine, "Authentication") {
			parts := strings.SplitN(trimmedLine, ":", 2)
			if len(parts) == 2 {
				auth := strings.ToLower(strings.TrimSpace(parts[1]))
				if strings.Contains(auth, "wpa3") {
					return "WPA3"
				} else if strings.Contains(auth, "wpa2") {
					return "WPA2"
				} else if strings.Contains(auth, "wpa") {
					return "WPA"
				} else if strings.Contains(auth, "open") {
					return "nopass"
				}
			}
		}
	}
	return "WPA"
}

// PasswordResult holds password and security info
type PasswordResult struct {
	Password     string
	SecurityType string
}

// GetPasswordAndType retrieves the WiFi password and security type for a given SSID
func GetPasswordAndType(ssid string) (*PasswordResult, error) {
	cmd := exec.Command("netsh", "wlan", "show", "profile", ssid, "key=clear")
	output, err := cmd.CombinedOutput()
	if err != nil {
		message := strings.ToLower(string(output))
		if strings.Contains(message, "access is denied") {
			return nil, fmt.Errorf("failed to retrieve WiFi password: access denied (try running as administrator)")
		}
		return nil, fmt.Errorf("failed to retrieve WiFi password: %w", err)
	}

	result := &PasswordResult{
		SecurityType: GetSecurityType(output),
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
				result.Password = strings.TrimSpace(parts[1])
				return result, nil
			}
		}
	}

	if securityKeyAbsent {
		result.SecurityType = "nopass"
		return result, nil
	}

	return nil, fmt.Errorf("could not extract password for SSID: %s", ssid)
}
