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
│   ├── logging/            # Structured logging setup
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
│       ├── provider_service.go  # Provider creation and fetching
│       ├── model_test.go   # Model tests
│       └── styles.go       # Lip Gloss styling
├── config.yaml.example     # Configuration template
├── go.mod
└── go.sum
```

## Core Components

### Config (`internal/config/`)
Handles loading and managing configuration from `config.yaml` with environment variable overrides.

**Environment Variable Overrides:**
- Environment variables take precedence over `config.yaml` values
- Format: `SCHOTY_<PROVIDER>_API_KEY`
- Provider names are normalized: dots (`.`), spaces (` `), and dashes (`-`) are replaced with underscores
- Examples:
  - OpenAI → `SCHOTY_OPENAI_API_KEY`
  - Claude Code → `SCHOTY_CLAUDE_CODE_API_KEY`
  - Z.ai → `SCHOTY_Z_AI_API_KEY`

**Config Loading:**
- Searches for `config.yaml` in current directory and `$HOME/.schoty/`
- Missing config file returns empty config (no error)
- `enabled` field controls whether a provider is active

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

## CLI Flags

```
-h, --help      display help information
-v, --version   display version information
--log-level     set log level: debug, info, warn, error (default: info)
```
