package qrcode

import (
	"fmt"
	"os"
	"strings"

	"github.com/mdp/qrterminal/v3"
	qr "github.com/skip2/go-qrcode"
)

// GenerateWiFiQR creates a QR code for WiFi connection details
// WiFi QR format: WIFI:T:WPA;S:SSID;P:PASSWORD;;
func GenerateWiFiQR(ssid, password, outputPath string) error {
	// Build WiFi connection string
	wifiString := fmt.Sprintf("WIFI:T:WPA;S:%s;P:%s;;", ssid, password)

	// Generate QR code
	err := qr.WriteFile(wifiString, qr.Medium, 256, outputPath)
	if err != nil {
		return fmt.Errorf("failed to generate QR code: %w", err)
	}

	return nil
}

// DisplayQRCodeTerminal displays the QR code in the terminal
func DisplayQRCodeTerminal(ssid, password string, size int) {
	wifiString := fmt.Sprintf("WIFI:T:WPA;S:%s;P:%s;;", ssid, password)

	fmt.Println("\n📱 Scan this QR code to connect:")
	fmt.Println()

	if size < 1 {
		size = 1
	}

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
