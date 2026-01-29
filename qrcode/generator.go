package qrcode

import (
	"fmt"
	"os"
	"strings"

	"github.com/mdp/qrterminal/v3"
	qr "github.com/skip2/go-qrcode"
)

// escapeWiFiString escapes special characters for WiFi QR code format
// Special chars: \ ; , : "
func escapeWiFiString(s string) string {
	replacer := strings.NewReplacer(
		"\\", "\\\\",
		";", "\\;",
		",", "\\,",
		":", "\\:",
		"\"", "\\\"",
	)
	return replacer.Replace(s)
}

// GenerateWiFiQR creates a QR code for WiFi connection details
// WiFi QR format:
//   - WPA/WPA2: WIFI:T:WPA;S:SSID;P:PASSWORD;;
//   - Open:    WIFI:T:nopass;S:SSID;;
func GenerateWiFiQR(ssid, password, outputPath string) error {
	// Build WiFi connection string with escaped special characters
	wifiString := buildWiFiString(ssid, password)

	// Generate QR code
	err := qr.WriteFile(wifiString, qr.Medium, 256, outputPath)
	if err != nil {
		// Check for permission-related errors
		if strings.Contains(err.Error(), "permission") || strings.Contains(err.Error(), "access") {
			return fmt.Errorf("failed to write QR code: permission denied (check directory permissions)")
		}
		return fmt.Errorf("failed to generate QR code: %w", err)
	}

	return nil
}

// DisplayQRCodeTerminal displays the QR code in the terminal
func DisplayQRCodeTerminal(ssid, password string, size int) {
	wifiString := buildWiFiString(ssid, password)

	fmt.Println("\n📱 Scan this QR code to connect:")
	fmt.Println()

	block := strings.Repeat("██", size)
	space := strings.Repeat("  ", size)

	config := qrterminal.Config{
		Level:     qrterminal.L,
		Writer:    os.Stdout,
		BlackChar: block,
		WhiteChar: space,
		QuietZone: 0,
	}

	qrterminal.GenerateWithConfig(wifiString, config)

	fmt.Println()
}

func buildWiFiString(ssid, password string) string {
	escapedSSID := escapeWiFiString(ssid)
	escapedPassword := escapeWiFiString(password)
	if password == "" {
		return fmt.Sprintf("WIFI:T:nopass;S:%s;;", escapedSSID)
	}
	return fmt.Sprintf("WIFI:T:WPA;S:%s;P:%s;;", escapedSSID, escapedPassword)
}
