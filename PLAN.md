# Project Strategy: Schoty - Milestone 2: TUI Foundation

## 1. High-Level Strategy

Milestone 2 completes the TUI foundation by wiring up the re-render loop, implementing auto-refresh, and establishing the config layer. This builds on the existing UI skeleton (Dashboard, Detail, Config, Help views) and keyboard navigation from Milestone 1.

**Key Objectives:**
- Implement auto-refresh timer mechanism in Bubble Tea model
- Wire `refreshAll()` to trigger actual data fetching via tick messages
- Create config loading infrastructure (`internal/config/`)
- Create data models (`internal/models/`)
- Add subscription tab content to the TUI
- Ensure proper state management for loading/fetching states

**Dependencies on Milestone 1:**
- UI model with ViewState, keyboard handling, and view rendering (COMPLETE)
- Styles system with theme detection (COMPLETE)
- Provider interface and stub implementations (COMPLETE)

## 2. Implementation Plan

### Phase 1: Config & Models Layer (COMPLETED)
- [x] Create `internal/config/config.go`:
  - [x] Config struct with API keys per provider
  - [x] Load from `config.yaml` using Viper
  - [x] Environment variable override support
  - [x] Provider enable/disable flags
- [x] Create `internal/models/models.go`:
  - [x] AppConfig model
  - [x] ProviderConfig model  
  - [x] RefreshState model for tracking refresh status
- [x] Add tests for config loading

### Phase 2: Auto-Refresh Timer Foundation (COMPLETED)
- [x] Add timer command to `Model.Init()` using `time.NewTicker`
- [x] Implement `tickTimerCmd()` that sends `tickMsg` periodically
- [x] Handle `tickMsg` in `Update()` to trigger refresh
- [x] Add `RefreshInterval` field to Model (default 60 seconds)
- [x] Clean up ticker on quit

### Phase 3: Provider Integration (COMPLETED)  
- [x] Create `internal/ui/provider_service.go`:
  - [x] Function to create provider instances from config
  - [x] Function to fetch all provider data concurrently
  - [x] Result channel to return provider states
- [x] Modify `refreshAll()` to return `tea.Cmd`
- [x] Wire provider results back to ProviderState in Update()
- [x] Handle provider errors and set ErrorMsg on ProviderState
- [x] Update Init() to load config and initialize providers

### Phase 4: Subscription Tab Content (COMPLETED)
- [x] Update `renderTabs()` to include "Subscriptions" tab
- [x] Create subscription-specific view rendering in dashboard
- [x] Add subscription data display (Used/Limit/Remaining)
- [x] Differentiate progress bar calculation for usage vs balance
- [x] Update footer to reflect subscription count

### Phase 5: Polish & Testing (COMPLETED)
- [x] Ensure ticker stops on quit
- [x] Add bounds checking for window resize
- [x] Create `internal/config/config_test.go` - test config loading, env overrides
- [x] Create `internal/ui/model_test.go` - test model init, refresh, timer
- [x] Run full test suite: `go test ./...`
- [x] Verify lint passes: `go vet ./...`

### Phase 6: Structured Logging (COMPLETED - Issue #32)
- [x] Add `log/slog` standard library for structured logging
- [x] Replace `fmt.Fprintf` with structured log statements
- [x] Add log levels (debug, info, warn, error)
- [x] Add contextual fields (provider name, request duration, etc.)
- [x] Create `internal/logging/logger.go` for centralized logger setup

### Phase 7: Graceful Shutdown (COMPLETED - Issue #33)
- [x] Add signal handling for SIGINT/SIGTERM
- [x] Create shutdown channel to coordinate cleanup
- [x] Stop ticker and cleanup resources on shutdown
- [x] Handle shutdown in main.go before tea.Program.Start()

### Phase 8: Help Flag (COMPLETED - Issue #35)
- [x] Add `flag` package for CLI flag parsing
- [x] Add `-h` and `--help` flags to main.go
- [x] Display usage information and exit cleanly
- [x] Suppress TUI startup when help is requested

### Phase 9: Hardening & Polish (COMPLETED)
- [x] Fix division by zero in Detail View for balance providers
- [x] Fix Detail View data mismatch (balance vs subscription rendering)
- [x] Add empty tab state message when no providers in filtered view
- [x] Add 30s network timeout to FetchAllProviders
- [x] Fix footer to show filtered provider count ("X of Y providers")
- [x] Fix signal handler goroutine cleanup with sync.WaitGroup

## 3. Execution Checklist

### Git Operations (COMPLETED)
- [x] Create feature branch: `git checkout -b milestone-2-tui-foundation`
- [x] Commit all Milestone 2 changes
- [x] Push branch: `git push -u origin milestone-2-tui-foundation`

### Files Created/Modified

**Phase 1:**
- `internal/config/config.go` (NEW)
- `internal/config/config_test.go` (NEW)
- `internal/models/models.go` (NEW)

**Phase 2:**
- `internal/ui/model.go` (MODIFIED)

**Phase 3:**
- `internal/ui/provider_service.go` (NEW)

**Phase 4:**
- `internal/ui/model.go` (MODIFIED - added provider types and filtering)

**Phase 5:**
- `internal/ui/model_test.go` (NEW)
- `go.mod`, `go.sum` (MODIFIED - added viper, testify deps)

**Phase 6 (NEW):**
- `internal/logging/logger.go` (NEW)
- `cmd/schoty/main.go` (MODIFIED - add structured logging)

**Phase 7 (NEW):**
- `cmd/schoty/main.go` (MODIFIED - add signal handling)

**Phase 8 (NEW):**
- `cmd/schoty/main.go` (MODIFIED - add flag parsing for help)

**Phase 9:**
- `internal/ui/model.go` (MODIFIED - detail view fixes, empty tab state, footer count)
- `internal/ui/provider_service.go` (MODIFIED - add timeout to FetchAllProviders)
- `cmd/schoty/main.go` (MODIFIED - signal handler cleanup)
