# Schoty

A TUI (Terminal User Interface) for monitoring AI subscription usage and API balances.

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

## Features

- Real-time balance tracking for multiple AI providers
- Subscription usage monitoring with quota alerts
- Terminal-native TUI for fast, efficient workflow
- Single dashboard view of all your AI service usages

## Installation

```bash
# Clone the repository
git clone https://github.com/ChrisUFO/Schoty.git
cd Schoty

# Build and run
go build -o schoty
./schoty
```

## Tech Stack

Built with Go and [Bubble Tea](https://github.com/charmbracelet/bubbletea) for the TUI.

## License

MIT
