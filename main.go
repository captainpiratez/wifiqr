package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/captainpiratez/wifiqr/qrcode"
	"github.com/captainpiratez/wifiqr/wifi"
)

func main() {
	ssidFlag := flag.String("ssid", "", "WiFi network name (optional)")
	ssidShortFlag := flag.String("s", "", "WiFi network name (optional) (shorthand)")
	outputFlag := flag.String("output", "wifi_qr.png", "Output file path (used with -image)")
	outputShortFlag := flag.String("o", "wifi_qr.png", "Output file path (used with -image) (shorthand)")
	imageFlag := flag.Bool("image", false, "Save QR code as an image instead of terminal output")
	imageShortFlag := flag.Bool("i", false, "Save QR code as an image instead of terminal output (shorthand)")
	sizeFlag := flag.Int("size", 1, "Terminal QR size multiplier (1 = default)")
	helpFlag := flag.Bool("help", false, "Show help")
	helpShortFlag := flag.Bool("h", false, "Show help")
	flag.Usage = func() {
		fmt.Println("wifiqr - Generate WiFi QR codes on Windows")
		fmt.Println()
		fmt.Println("Usage:")
		fmt.Println("  wifiqr [options]")
		fmt.Println()
		fmt.Println("Options:")
		fmt.Println("  --ssid, -s <name>      WiFi network name (optional)")
		fmt.Println("  --image, -i            Save QR code as an image (default is terminal output)")
		fmt.Println("  --output, -o <path>    Output file path (used with --image)")
		fmt.Println("  --size <n>             Terminal QR size multiplier (1 = default)")
		fmt.Println("  --help, -h             Show help")
		fmt.Println()
		fmt.Println("Examples:")
		fmt.Println("  wifiqr")
		fmt.Println("  wifiqr --ssid MyNetwork")
		fmt.Println("  wifiqr --image --output wifi_qr.png")
		fmt.Println("  wifiqr --size 2")
	}
	flag.Parse()

	if *helpFlag || *helpShortFlag {
		flag.Usage()
		return
	}

	var ssid string
	var password string
	var err error

	if *ssidFlag == "" && *ssidShortFlag != "" {
		*ssidFlag = *ssidShortFlag
	}

	if *ssidFlag != "" {
		// Use specified SSID
		ssid = *ssidFlag
		password, err = wifi.GetPassword(ssid)
	} else {
		// Get current connected WiFi
		ssid, err = wifi.GetCurrentSSID()
		if err != nil {
			log.Fatalf("Failed to get current WiFi: %v", err)
		}

		password, err = wifi.GetPassword(ssid)
	}

	if err != nil {
		log.Fatalf("Failed to retrieve WiFi details: %v", err)
	}

	fmt.Printf("Network: %s\n", ssid)
	fmt.Println("Generating QR code...")

	if *outputFlag == "wifi_qr.png" && *outputShortFlag != "wifi_qr.png" {
		*outputFlag = *outputShortFlag
	}

	if *imageFlag || *imageShortFlag {
		// Save to file
		err = qrcode.GenerateWiFiQR(ssid, password, *outputFlag)
		if err != nil {
			log.Fatalf("Failed to generate QR code: %v", err)
		}
		fmt.Printf("✓ QR code saved to: %s\n", *outputFlag)
	} else {
		// Display in terminal (default)
		qrcode.DisplayQRCodeTerminal(ssid, password, *sizeFlag)
	}
}
