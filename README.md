# Cliforge

[![View on KikPlate](https://img.shields.io/static/v1?label=KikPlate&message=golang-cli-starter&color=0366d6&style=flat-square)](https://kikplate.dev/plates/golang-cli-starter)

> **Forge your CLI tools with confidence.**

Cliforge is a production-ready Go CLI boilerplate built on [Cobra](https://github.com/spf13/cobra), [Viper](https://github.com/spf13/viper), and [Zap](https://github.com/uber-go/zap). It gives you a clean, scalable foundation for building professional command-line tools — with config management, structured logging, versioning, and tests already wired up.

```bash
go install github.com/kikplate/golang-cli-starter@latest
```

---

## Features

- **Cobra** subcommand tree with persistent flags
- **Viper** config: file + environment variables (`CLIFORGE_*`)
- **Zap** structured logger with `--verbose` debug mode
- **Build-time version injection** via `ldflags`
- **Unit + integration tests** with race detector
- Clean separation: `cmd/`, `internal/`, `pkg/`
- `Makefile` for build, test, and lint shortcuts

---

## Installation

### From source

```bash
git clone https://github.com/kikplate/golang-cli-starter.git
cd golang-cli-starter
make build
./bin/cliforge --help
```

### With `go install`

```bash
go install github.com/kikplate/golang-cli-starter@latest
```

> Requires Go 1.22+

---

## Usage

```
cliforge [command] [flags]
```

### Global flags

| Flag | Shorthand | Default | Description |
|------|-----------|---------|-------------|
| `--config` | | `~/.cliforge.yaml` | Path to config file |
| `--verbose` | `-v` | `false` | Enable debug logging |
| `--help` | `-h` | | Show help |

---

### Commands

#### `greet`

Greet someone by name.

```bash
cliforge greet Alice
cliforge greet --name Bob
cliforge greet Alice --shout
cliforge greet -v Alice          # with debug logging
```

| Flag | Default | Description |
|------|---------|-------------|
| `--name` | `""` | Name of the person to greet |
| `--shout` | `false` | Print greeting in uppercase |

#### `version`

Print the build version, git commit, and build date.

```bash
cliforge version
# cliforge v1.2.0 (commit: a3f9c12, built: 2024-03-15T10:00:00Z)
```

---

## Configuration

Cliforge loads configuration in this priority order (highest wins):

1. Command-line flags
2. Environment variables (`CLIFORGE_<KEY>`)
3. Config file (`~/.cliforge.yaml` or `--config` path)
4. Defaults

### Example config file

```yaml
# ~/.cliforge.yaml
verbose: false
```

### Environment variables

```bash
CLIFORGE_VERBOSE=true cliforge greet Alice
```

---

## Project Structure

```
cliforge/
├── cmd/
│   ├── root.go          # Root command, persistent flags, global setup
│   ├── version.go       # `cliforge version` subcommand
│   ├── greet.go         # Example subcommand: `cliforge greet`
│   └── root_test.go     # Integration tests for commands
├── internal/
│   ├── config/
│   │   └── config.go    # Config loading via Viper
│   └── logger/
│       └── logger.go    # Zap logger initialisation
├── pkg/
│   └── greeting/
│       ├── greeting.go       # Reusable greeting logic
│       └── greeting_test.go  # Unit tests
├── docs/
│   ├── README.md            # This file
│   ├── CONTRIBUTING.md      # How to contribute
│   ├── CHANGELOG.md         # Version history
│   └── commands/
│       ├── greet.md         # greet command reference
│       └── version.md       # version command reference
├── Makefile
├── go.mod
└── main.go
```

---

## Testing

```bash
# Run all tests
make test

# Run with verbose output
go test -v -race ./...

# Run only unit tests (pkg/)
go test ./pkg/...

# Run only command integration tests (cmd/)
go test ./cmd/...
```

---

## 🔨 Development

```bash
# Build binary
make build

# Build with version info
make build VERSION=1.0.0

# Run linter (requires golangci-lint)
make lint

# Clean build artifacts
make clean
```

### Adding a new subcommand

1. Create `cmd/mycommand.go`
2. Define a `*cobra.Command` and register it with `rootCmd.AddCommand(myCmd)` in `init()`
3. Put reusable logic in `pkg/mypackage/`
4. Write tests in `cmd/root_test.go` and `pkg/mypackage/mypackage_test.go`

---

## Requirements

- Go 1.22+
- [golangci-lint](https://golangci-lint.run/) (for `make lint`)
