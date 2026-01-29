# WiFiQR

A Windows 10 command-line tool that generates QR codes for currently connected WiFi networks, making it easy to share your WiFi with others.

## Features

- 📡 Detects currently connected WiFi network
- 🔐 Retrieves WiFi password (requires admin privileges)
- 🎯 Generates QR code for easy phone connection
- 🖥️ Displays QR code in the terminal by default
- 💾 Saves QR code as PNG image with --image

## Requirements

- Windows 10 or later
- Go 1.21+
- Administrator privileges (to retrieve WiFi passwords)

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

### Generate QR for current WiFi

```bash
wifiqr
```

### Specify custom WiFi network

```bash
wifiqr --ssid "Your Network Name"
```

### Save QR code as an image

```bash
wifiqr --image --output wifi_qr.png
```

### Adjust terminal QR size

```bash
wifiqr --size 2
```

### Short flags

```bash
wifiqr -s "Your Network Name" -i -o wifi_qr.png
```

## Output

By default, the QR code is displayed in the terminal. Use `--image` to save a PNG file.

## Development

```bash
# Build
go build

# Test
go test ./...

# Run
./wifiqr.exe
```

## License

MIT
