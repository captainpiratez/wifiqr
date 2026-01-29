# WiFiQR

A Windows 10 command-line tool that generates QR codes for WiFi networks, making it easy to share your WiFi with others.

## Features

- 📡 Auto-detects currently connected WiFi network
- ⌨️ Interactive input if not connected to any network
- 🔐 Retrieves WiFi password (admin privileges may be required)
- 🎯 Detects WiFi security type (WPA/WPA2/WPA3/Open)
- 🖥️ Displays QR code in the terminal by default
- 💾 Saves QR code as PNG image with `--image`
- 🎨 Customizable terminal QR size (1-5x)
- 📋 Short and long flag options for convenience

## Requirements

- Windows 10 or later
- Go 1.21+ (to build from source)
- Administrator privileges (optional, may be needed to retrieve WiFi passwords)

## Installation

```bash
go install github.com/captainpiratez/wifiqr@latest
```

Or build from source:

```bash
git clone https://github.com/captainpiratez/wifiqr.git
cd wifiqr
go build
```

## Usage

### Basic Usage - Generate QR for current WiFi

```bash
wifiqr
```

### Specify a custom WiFi network

```bash
wifiqr --ssid "Your Network Name"
wifiqr -s "Your Network Name"
```

### Save QR code as a PNG image

```bash
wifiqr --image --output wifi_qr.png
wifiqr -i -o wifi_qr.png
```

### Adjust terminal QR code size

```bash
wifiqr --size 2
```

### Show help

```bash
wifiqr --help
wifiqr -h
```

## Command-Line Options

```
--ssid, -s <name>      WiFi network name (optional)
--image, -i            Save QR code as an image (default is terminal output)
--output, -o <path>    Output file path (used with --image)
--size <n>             Terminal QR size multiplier (1-5, default: 1)
--help, -h             Show help
```

## Examples

### Display terminal QR for current network
```bash
wifiqr
```

### Display larger terminal QR
```bash
wifiqr --size 3
```

### Specify network and save as PNG
```bash
wifiqr -s "MyNetwork" -i -o qr.png
```

### Generate QR for network not currently connected to
```bash
wifiqr --ssid "GuestNetwork" --image --output guest_qr.png
```

### Interactive mode (when not connected)
If you're not connected to any WiFi, the tool will prompt you to enter a network name:
```
⚠️  No connected WiFi network found.

Enter WiFi network name (SSID): MyNetwork
```

## Output

By default, the QR code is displayed in the terminal as ASCII art. Use `--image` to save a PNG file instead.

## Development

```bash
# Build
go build

# Run
./wifiqr.exe

# Build release (optimized, smaller size)
go build -ldflags="-s -w"
```

## Notes

- Requires administrator privileges on some systems to retrieve WiFi passwords
- SSID and password special characters are properly escaped for QR code compatibility
- Supports WPA, WPA2, WPA3, and open networks
- Terminal QR codes work best on dark backgrounds with light text

## License

MIT
