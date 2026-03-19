# Schoty Architecture

## Overview

Schoty is built with Go and Bubble Tea for a reactive, terminal-native user interface.

## Project Structure

```
Schoty/
├── cmd/                    # Application entry points
│   └── schoty/             # Main CLI application
├── internal/               # Private application code
│   ├── config/             # Configuration loading and management
│   ├── models/             # Data structures and types
│   ├── providers/          # API client implementations
│   │   ├── anthropic.go
│   │   ├── openai.go
│   │   ├── openrouter.go
│   │   ├── togetherai.go
│   │   ├── claude_code.go
│   │   ├── codex.go
│   │   ├── zai.go
│   │   └── minimax.go
│   └── ui/                 # TUI components
│       ├── model.go        # Main Bubble Tea model
│       ├── views/          # View components
│       └── styles.go       # Lip Gloss styling
├── config.yaml             # Configuration file
├── go.mod
└── go.sum
```

## Core Components

### Config (`internal/config/`)
Handles loading and managing configuration from `config.yaml` with environment variable overrides.

### Models (`internal/models/`)
Defines data structures for:
- `Balance` - API balance information
- `Usage` - Subscription usage data
- `Provider` - Provider metadata and status

### Providers (`internal/providers/`)
Each provider implements a common interface:

```go
type Provider interface {
    Name() string
    CheckBalance() (*Balance, error)
    CheckUsage() (*Usage, error)
}
```

### UI (`internal/ui/`)
Bubble Tea model and views that render the TUI.

## Data Flow

1. App starts → Config loads API keys
2. Model initializes → Providers created
3. On refresh → Each provider fetches data concurrently
4. Data returned → UI updates reactively
5. On error → Error displayed inline, no crash

## Dependencies

- [Bubble Tea](https://github.com/charmbracelet/bubbletea) - TUI framework
- [Lip Gloss](https://github.com/charmbracelet/lipgloss) - Styling
- [Viper](https://github.com/spf13/viper) - Configuration management
