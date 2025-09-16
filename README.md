# ttrk

A command-line interface tool built with Go.

## Project Structure

```
ttrk/
├── cmd/
│   └── ttrk/           # Main application entry point
│       └── main.go
├── internal/           # Private application code
│   ├── cli/           # CLI commands and logic
│   │   ├── root.go
│   │   └── version.go
│   └── config/        # Configuration management
│       └── config.go
├── pkg/               # Public library code
│   └── version/       # Version information
│       ├── version.go
│       └── version_test.go
├── scripts/           # Build and development scripts
│   └── build.sh
├── docs/              # Documentation
├── Makefile          # Build automation
├── go.mod            # Go module definition
└── README.md         # This file
```

## Installation

### From Source

```bash
git clone https://github.com/MBH999/ttrk.git
cd ttrk
make build
```

The binary will be available in the `bin/` directory.

## Usage

```bash
# Run the application
./bin/ttrk

# Show help
./bin/ttrk --help

# Show version
./bin/ttrk version
```

## Configuration

Create `~/.config/ttrk/config.ini` to customise runtime behaviour. The file uses a simple `key=value` format.

- `data_dir` — directory where tracker data is stored. Defaults to `~/.config/ttrk`.

## Development

### Prerequisites

- Go 1.19 or later
- Make (optional, for using Makefile commands)

### Building

```bash
# Using Makefile
make build

# Or using Go directly
go build -o bin/ttrk ./cmd/ttrk

# Or using the build script
./scripts/build.sh
```

### Testing

```bash
# Run all tests
make test

# Or using Go directly
go test ./...
```

### Other Development Commands

```bash
# Format code
make fmt

# Run linter (requires golangci-lint)
make lint

# Run go vet
make vet

# Clean build artifacts
make clean

# Show all available make targets
make help
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.
