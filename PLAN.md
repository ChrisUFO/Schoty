# Project Strategy: Schoty - Milestone 2: TUI Foundation

## 1. High-Level Strategy

Milestone 2 completes the TUI foundation by wiring up the re-render loop, implementing auto-refresh, and establishing the config layer. This builds on the existing UI skeleton (Dashboard, Detail, Config, Help views) and keyboard navigation from Milestone 1.

**Key Objectives:**
- Implement auto-refresh timer mechanism in Bubble Tea model
- Wire up `refreshAll()` to trigger actual data fetching via tick messages
- Create config loading infrastructure (`internal/config/`)
- Create data models (`internal/models/`)
- Add subscription tab content to the TUI
- Ensure proper state management for loading/fetching states

**Dependencies on Milestone 1:**
- UI model with ViewState, keyboard handling, and view rendering (COMPLETE)
- Styles system with theme detection (COMPLETE)
- Provider interface and stub implementations (COMPLETE)

## 2. Implementation Plan

### Phase 1: Config & Models Layer
- Create `internal/config/config.go`:
  - Config struct with API keys per provider
  - Load from `config.yaml` using Viper
  - Environment variable override support
  - Provider enable/disable flags
- Create `internal/models/models.go`:
  - AppConfig model
  - ProviderConfig model  
  - RefreshState model for tracking refresh status

### Phase 2: Auto-Refresh Timer Foundation
- Add timer command to `Model.Init()` using `time.NewTicker`
- Implement `tickTimerCmd()` that sends `tickMsg` periodically
- Handle `tickMsg` in `Update()` to trigger refresh
- Add `RefreshInterval` field to Model (default 60 seconds)
- Clean up ticker on quit

### Phase 3: Provider Integration  
- Wire `refreshAll()` to actually call provider `CheckBalance()` and `CheckUsage()` concurrently
- Add `tea.Cmd` return from Update to run providers in background
- Implement result handling to update ProviderState from provider responses
- Add error handling that sets ErrorMsg on ProviderState
- Create `internal/ui/provider_service.go` to bridge UI and providers

### Phase 4: Subscription Tab Content
- Implement subscription tab rendering (currently just shows "Subscriptions" tab)
- Create subscription-specific view rendering in dashboard
- Add subscription status indicators
- Differentiate balance vs usage display between tabs

### Phase 5: Polish & Testing
- Add proper cleanup on quit (stop ticker)
- Add window resize optimization
- Create unit tests for config loading
- Create unit tests for auto-refresh timer
- Create integration tests for provider service

## 3. Execution Checklist

### Phase 1: Config & Models Layer
- [ ] Create `internal/config/config.go` with Viper-based config loading
- [ ] Define Config struct with provider API keys and settings
- [ ] Implement `LoadConfig()` function
- [ ] Implement environment variable override support
- [ ] Create `internal/models/models.go` with AppConfig, ProviderConfig types
- [ ] Add GetProviderConfigByName helper
- [ ] Write tests for config loading in `internal/config/config_test.go`

### Phase 2: Auto-Refresh Timer Foundation  
- [ ] Add timer command to `Model.Init()` using `time.NewTicker`
- [ ] Create `tickTimerCmd()` function that sends tickMsg periodically
- [ ] Add `RefreshInterval` field to Model (default 60 seconds)
- [ ] Handle tickMsg in Update() to trigger refreshAll
- [ ] Stop ticker properly on quit (cleanup in Update when model quits)
- [ ] Write tests for timer functionality

### Phase 3: Provider Integration
- [ ] Create `internal/ui/provider_service.go`:
  - [ ] Function to create provider instances from config
  - [ ] Function to fetch all provider data concurrently
  - [ ] Result channel to return provider states
- [ ] Modify `refreshAll()` to return `tea.Cmd`
- [ ] Wire provider results back to ProviderState in Update()
- [ ] Handle provider errors and set ErrorMsg on ProviderState
- [ ] Update Init() to load config and initialize providers

### Phase 4: Subscription Tab Content
- [ ] Update `renderTabs()` to include "Subscriptions" tab
- [ ] Create subscription-specific row rendering in `renderProviderRow()`
- [ ] Add subscription data display (Used/Limit/Remaining)
- [ ] Differentiate progress bar calculation for usage vs balance
- [ ] Update footer to reflect subscription count

### Phase 5: Polish & Testing
- [ ] Ensure ticker stops on quit
- [ ] Add bounds checking for window resize
- [ ] Create `internal/config/config_test.go` - test config loading, env overrides
- [ ] Create `internal/ui/provider_service_test.go` - test provider creation, fetching
- [ ] Create `internal/ui/model_test.go` - test model init, refresh, timer
- [ ] Run full test suite: `go test ./...`
- [ ] Verify lint passes: `golangci-lint run`

### Git Operations
- [ ] Create feature branch: `git checkout -b milestone-2-tui-foundation`
- [ ] Commit all Milestone 2 changes
- [ ] Push branch: `git push -u origin milestone-2-tui-foundation`
