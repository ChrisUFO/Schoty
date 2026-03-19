# Schoty MVP Roadmap

> **Project Status: ARCHIVED** (2026-03-19)
> 
> After research, it was discovered that **OpenRouter is the only provider** that offers a public API endpoint to check account balance/credits. All other providers (Anthropic, OpenAI, Codex, Z.ai, MiniMax, Together.ai) require manual dashboard checking - no public API exists for programmatic access.
>
> Since the core value proposition was aggregating multiple providers in one view, and only OpenRouter supports this, the project does not provide sufficient value to continue.

## Milestone 1: Project Scaffold ✅
- [x] Initialize Go module with proper structure (`cmd/`, `internal/`, `pkg/`)
- [x] Set up Bubble Tea with basic TUI skeleton
- [x] Add Lip Gloss for styling
- [x] Create `main.go` entry point
- [x] Add pre-commit and pre-push hooks with prettier, lint checks, and tests
- [x] Add Makefile for standardized build/run commands
- [x] Add .editorconfig for consistent editor settings

## Milestone 2: TUI Foundation ✅
- [x] Build basic navigation model (tabs or list selection)
- [x] Create main dashboard layout
- [x] Add keyboard navigation (quit, refresh, tab switching)
- [x] Wire up re-render loop
- [x] Implement structured logging
- [x] Add graceful shutdown handling (SIGINT/SIGTERM)
- [x] Add version flag (-v, --version)
- [x] Add help flag (-h, --help)

## Milestone 3: Configuration Layer ✅
- [x] Define `config.yaml` structure for API keys
- [x] Load config on startup
- [x] Environment variable override support for keys
- [x] Service selection (enable/disable providers)

## Milestone 4: API Client Skeleton ✅
- [x] Create provider interface pattern
- [x] Build stub implementations for all 8 services
- [x] Implement refresh timer
- [x] Add basic error handling

## Milestone 5: Provider Implementation - API Balance ❌ (Not Feasible)
- [ ] **5.1** - ~~Implement Anthropic balance check~~ - No API available
- [ ] **5.2** - ~~Implement OpenAI balance check~~ - No API available
- [x] **5.3** - OpenRouter balance check - API EXISTS but not implemented (stub only)
- [ ] **5.4** - ~~Implement Together.ai balance check~~ - No API available

## Milestone 6: Provider Implementation - Subscriptions ❌ (Not Feasible)
- [ ] **6.1** - ~~Implement Claude Code usage check~~ - No API available
- [ ] **6.2** - ~~Implement Codex usage check~~ - No API available
- [ ] **6.3** - ~~Implement Z.ai usage check~~ - No API available
- [ ] **6.4** - ~~Implement MiniMax usage check~~ - No API available

## Milestone 7: Polish ❌ (Not applicable without working providers)
- [ ] Add loading states and spinners
- [ ] Add error display for failed API calls
- [ ] Add auto-refresh with configurable interval
- [ ] Format currency/numbers nicely

## Research Findings

### Providers with Public Balance/Usage APIs
| Provider | Balance API | Usage API | Notes |
|----------|-----------|----------|-------|
| OpenRouter | ✅ Yes | ✅ Yes | `GET /api/v1/credits` |
| Anthropic | ❌ No | ❌ No | Console only |
| OpenAI | ❌ No | ❌ No | Console only |
| Together.ai | ❌ No | ❌ No | Console only |
| Claude Code | ❌ No | ❌ No | Console only |
| Codex | ❌ No | ❌ No | Console only |
| Z.ai | ❌ No | ❌ No | Dashboard quota display only |
| MiniMax | ❌ No | ❌ No | Dashboard quota display only |

### Conclusion
Building a unified dashboard is not viable when only 1 of 8 providers supports the necessary API. The OpenRouter website already provides this functionality for OpenRouter alone.