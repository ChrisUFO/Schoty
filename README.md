# Schoty

![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)
![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8.svg)

> **⚠️ This project is no longer being developed.**

## Why

Schoty was built to provide a unified terminal view of AI subscription usage and API balances across multiple providers. However, after research, we discovered that **OpenRouter is the only provider** that offers a public API endpoint for checking account balance/credits. All other providers (Anthropic, OpenAI, Codex, Z.ai, MiniMax, Together.ai) require manual dashboard checking - no public API exists for programmatic access.

Since the core value proposition was aggregating multiple providers in one view, and only OpenRouter supports this, the project does not provide sufficient value over using each provider's existing dashboard.

## What Was Built

A functional TUI application with:
- Bubble Tea-based terminal interface
- Tab-based navigation (Dashboard, Detail, Config, Help)
- Keyboard shortcuts (quit, refresh, tab switching, etc.)
- Configuration loading with YAML and environment variable overrides
- Structured logging
- Graceful shutdown handling

## Project Status

**Archived** - This project is no longer actively maintained.

## Documentation

- [Architecture](ARCHITECTURE.md) - Project structure and design
- [UI/UX Specification](UIUX.md) - TUI layout, navigation, and interaction design
- [Style Guide](STYLEGUIDE.md) - Visual design system and Lip Gloss styling
- [Roadmap](ROADMAP.md) - Development phases and research findings

## License

This project is licensed under the MIT License - see the LICENSE file for details.