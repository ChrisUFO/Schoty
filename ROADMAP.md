# Schoty MVP Roadmap

## Milestone 1: Project Scaffold
- [ ] Initialize Go module with proper structure (`cmd/`, `internal/`, `pkg/`)
- [ ] Set up Bubble Tea with basic TUI skeleton
- [ ] Add Lip Gloss for styling
- [ ] Create `main.go` entry point
- [ ] Add pre-commit and pre-push hooks with prettier, lint checks, and tests
- [ ] Add Makefile for standardized build/run commands
- [ ] Add .editorconfig for consistent editor settings

## Milestone 2: TUI Foundation
- [ ] Build basic navigation model (tabs or list selection)
- [ ] Create main dashboard layout
- [ ] Add keyboard navigation (quit, refresh, tab switching)
- [ ] Wire up re-render loop
- [ ] Implement structured logging
- [ ] Add graceful shutdown handling (SIGINT/SIGTERM)
- [ ] Add version flag (-v, --version)
- [ ] Add help flag (-h, --help)

## Milestone 3: Configuration Layer
- [ ] Define `config.yaml` structure for API keys
- [ ] Load config on startup
- [ ] Environment variable override support for keys
- [ ] Service selection (enable/disable providers)

## Milestone 4: API Client Skeleton
- [ ] Create provider interface pattern
- [ ] Build stub implementations for all 8 services
- [ ] Implement refresh timer
- [ ] Add basic error handling

## Milestone 5: Provider Implementation - API Balance

- [ ] **5.1** - Implement Anthropic balance check
- [ ] **5.2** - Implement OpenAI balance check
- [ ] **5.3** - Implement OpenRouter balance check
- [ ] **5.4** - Implement Together.ai balance check

## Milestone 6: Provider Implementation - Subscriptions

- [ ] **6.1** - Implement Claude Code usage check
- [ ] **6.2** - Implement Codex usage check
- [ ] **6.3** - Implement Z.ai usage check
- [ ] **6.4** - Implement MiniMax usage check

## Milestone 7: Polish
- [ ] Add loading states and spinners
- [ ] Add error display for failed API calls
- [ ] Add auto-refresh with configurable interval
- [ ] Format currency/numbers nicely

## Future Enhancements
- [ ] Persistent history and trend charts
- [ ] Alert thresholds and notifications
- [ ] Export data to CSV/JSON
- [ ] Multiple account support
