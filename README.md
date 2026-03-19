# Schoty

![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)
![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8.svg)

A TUI (Terminal User Interface) for monitoring AI subscription usage and API balances.

<!-- Screenshot placeholder -->

## Features

- Real-time balance tracking for multiple AI providers
- Subscription usage monitoring with quota alerts
- Terminal-native TUI for fast, efficient workflow
- Single dashboard view of all your AI service usages

## Supported Services

### API Balance Monitoring
- **Anthropic** - Track your Anthropic API credits and spending
- **OpenAI** - Monitor your OpenAI API usage and remaining balance
- **OpenRouter** - Keep an eye on your OpenRouter credits
- **Together.ai** - Monitor your Together.ai API balance

### Subscription Usage Monitoring
- **Anthropic (Claude Code)** - Track Claude Code subscription usage and quota
- **OpenAI (Codex)** - Monitor Codex subscription status and remaining calls
- **Z.ai (Coding Plan)** - Track Z.ai coding plan usage and limits
- **MiniMax (Token Plan)** - Monitor MiniMax token plan consumption

## Prerequisites

- Go 1.21 or higher

## Quick Start

```bash
# Clone the repository
git clone https://github.com/ChrisUFO/Schoty.git
cd Schoty

# Build
go build -o schoty

# Run
./schoty
```

## Configuration

Schoty uses a `config.yaml` file for API keys. See [Architecture.md](ARCHITECTURE.md) for configuration details.

## Documentation

- [Architecture](ARCHITECTURE.md) - Project structure and design
- [Roadmap](ROADMAP.md) - Development phases and progress

## Contributing

Contributions welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
