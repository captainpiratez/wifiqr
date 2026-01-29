package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/captainpiratez/wifiqr/qrcode"
	"github.com/captainpiratez/wifiqr/wifi"
)

func main() {
	ssidFlag := flag.String("ssid", "", "WiFi network name (optional)")
	ssidShortFlag := flag.String("s", "", "WiFi network name (optional) (shorthand)")
	outputFlag := flag.String("output", "wifi_qr.png", "Output file path (used with -image)")
	outputShortFlag := flag.String("o", "", "Output file path (used with -image) (shorthand)")
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

	if *sizeFlag < 1 || *sizeFlag > 5 {
		log.Fatalf("Size must be between 1 and 5 (got %d)", *sizeFlag)
	}

	var ssid string
	var password string
	var err error

	ssid = *ssidFlag
	if ssid == "" {
		ssid = *ssidShortFlag
	}

	imageChoice := *imageFlag || *imageShortFlag
	outputPath := *outputFlag
	if *outputShortFlag != "" {
		outputPath = *outputShortFlag
	}

	if imageChoice {
		trimmedOutput := strings.TrimSpace(outputPath)
		if trimmedOutput == "" {
			log.Fatal("Output path cannot be empty when using --image")
		}
		outputPath = trimmedOutput
		outputDir := filepath.Dir(outputPath)
		if outputDir != "." && outputDir != "" {
			info, statErr := os.Stat(outputDir)
			if statErr != nil || !info.IsDir() {
				log.Fatalf("Output directory does not exist: %s", outputDir)
			}
		}
	}

	if ssid != "" {
		// Use specified SSID
		result, err := wifi.GetPasswordAndType(ssid)
		if err != nil {
			log.Fatalf("Failed to retrieve WiFi details: %v", err)
		}
		password = result.Password
	} else {
		// Get current connected WiFi
		ssid, err = wifi.GetCurrentSSID()
		if err != nil {
			fmt.Println("⚠️  No connected WiFi network found.")
			fmt.Println()
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Enter WiFi network name (SSID): ")
			ssid, _ = reader.ReadString('\n')
			ssid = strings.TrimSpace(ssid)
			if ssid == "" {
				log.Fatal("SSID cannot be empty")
			}
		}

		result, err := wifi.GetPasswordAndType(ssid)
		if err != nil {
			log.Fatalf("Failed to retrieve WiFi details: %v", err)
		}
		password = result.Password
	}

	fmt.Printf("Network: %s\n", ssid)
	fmt.Println("Generating QR code...")

	if imageChoice {
		// Save to file
		err = qrcode.GenerateWiFiQR(ssid, password, outputPath)
		if err != nil {
			log.Fatalf("Failed to generate QR code: %v", err)
		}
		fmt.Printf("✓ QR code saved to: %s\n", outputPath)
	} else {
		// Display in terminal (default)
		qrcode.DisplayQRCodeTerminal(ssid, password, *sizeFlag)
	}
}
