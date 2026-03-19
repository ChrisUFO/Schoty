# Project Strategy: Schoty - Milestone 1

## 1. High-Level Strategy

Milestone 1 establishes the foundational project scaffold for Schoty, a TUI AI subscription usage and balance monitor built with Go and Bubble Tea. This phase creates the empty shell project with proper structure, tooling, and development infrastructure before implementing business logic.

**Key Objectives:**
- Establish Go module with directory structure matching ARCHITECTURE.md
- Set up Bubble Tea with minimal TUI skeleton
- Configure Lip Gloss for terminal styling following STYLEGUIDE.md
- Implement UI components per UIUX.md specifications
- Add development tooling (Makefile, pre-commit hooks, editorconfig)
- Create a runnable "Hello World" TUI application

**Planning Complete:**
- [x] UIUX.md - TUI layout, navigation, and interaction design
- [x] STYLEGUIDE.md - Visual design system and Lip Gloss styling

## 2. Implementation Plan

### Phase 0: UI/UX Planning (COMPLETED)
- [x] Create UIUX.md with:
  - Layout architecture (header, tabs, main content, footer)
  - Navigation patterns and keyboard shortcuts
  - Component specifications (cards, detail panels, states)
  - User workflows (first launch, normal use, refresh, error recovery)
  - Responsive behavior specifications
- [x] Create STYLEGUIDE.md with:
  - Color definitions (light/dark terminal support)
  - Typography styles
  - Component style definitions
  - Layout utilities
  - Animation guidelines

### Phase 1: Repository & Module Setup
- Initialize Go module (`go mod init`)
- Create directory structure per ARCHITECTURE.md:
  - `cmd/schoty/` - Application entry points
  - `internal/config/` - Configuration loading
  - `internal/models/` - Data structures
  - `internal/providers/` - API clients (stub placeholders)
  - `internal/ui/` - TUI components
  - `internal/ui/views/` - View components

### Phase 2: Dependencies
- Add Bubble Tea (`github.com/charmbracelet/bubbletea`)
- Add Lip Gloss (`github.com/charmbracelet/lipgloss`)
- Add Viper (`github.com/spf13/viper`) for configuration
- Add testing dependencies (stretchr/testify)

### Phase 3: UI Foundation (per UIUX.md and STYLEGUIDE.md)
- Create `cmd/schoty/main.go` entry point
- Implement minimal Bubble Tea model
- Create `internal/ui/styles.go` with all style definitions from STYLEGUIDE.md:
  - Color constants (light/dark mode)
  - Status colors (healthy/warning/critical/error/loading)
  - Typography styles (header, body, caption, status)
  - Component styles (tabs, cards, footer, error)
  - Progress bar helper function
  - Spinner animation helper
- Wire up basic re-render loop
- Implement theme detection (light/dark terminal)

### Phase 4: Development Tooling
- Create Makefile with targets: `build`, `run`, `test`, `lint`, `fmt`, `clean`
- Add `.editorconfig` for consistent formatting
- Set up pre-commit hooks (gofmt, golint, govet, tests)
- Set up pre-push hooks (run full test suite)

### Phase 5: Documentation
- Ensure ARCHITECTURE.md reflects actual structure
- Add placeholder provider files with interface documentation

## 3. Execution Checklist

### Planning (COMPLETED)
- [x] Create UIUX.md with detailed TUI design specifications
- [x] Create STYLEGUIDE.md with Lip Gloss styling system
- [x] Update README.md to reference UI/UX and Style Guide docs

### Repository & Module Setup
- [ ] Initialize Go module: `go mod init github.com/ChrisUFO/Schoty`
- [ ] Create `cmd/schoty/` directory
- [ ] Create `internal/config/` directory
- [ ] Create `internal/models/` directory
- [ ] Create `internal/providers/` directory
- [ ] Create `internal/ui/` directory
- [ ] Create `internal/ui/views/` directory

### Dependencies
- [ ] Add Bubble Tea: `go get github.com/charmbracelet/bubbletea`
- [ ] Add Lip Gloss: `go get github.com/charmbracelet/lipgloss`
- [ ] Add Viper: `go get github.com/spf13/viper`
- [ ] Add Testify: `go get github.com/stretchr/testify`

### UI Foundation (UIUX.md + STYLEGUIDE.md)
- [ ] Create `cmd/schoty/main.go` with Bubble Tea program structure
- [ ] Create `internal/ui/styles.go` with complete style system:
  - [ ] Color constants (light/dark mode)
  - [ ] Status color functions
  - [ ] Typography style definitions
  - [ ] Tab styles (active/inactive)
  - [ ] Card styles
  - [ ] Progress bar rendering function
  - [ ] Spinner animation function
  - [ ] Theme detection function
- [ ] Create `internal/ui/model.go` with minimal Tea model
- [ ] Implement keyboard handler stubs for navigation
- [ ] Verify TUI compiles and runs (`make run`)

### Development Tooling
- [ ] Create Makefile with `build`, `run`, `test`, `lint`, `fmt`, `clean` targets
- [ ] Create `.editorconfig` with Go formatting conventions
- [ ] Create `.git/hooks/pre-commit` with fmt, lint, vet, test checks
- [ ] Create `.git/hooks/pre-push` with full test suite
- [ ] Install git hooks: `git config core.hooksPath .git/hooks`

### Documentation
- [ ] Verify ARCHITECTURE.md matches created structure
- [ ] Create placeholder `internal/providers/*.go` stub files with interface documentation

### Testing
- [ ] Create `cmd/schoty/main_test.go` - smoke test for app initialization
- [ ] Create `internal/ui/model_test.go` - test model initialization
- [ ] Create `internal/ui/styles_test.go` - test style constants
- [ ] Run full test suite: `make test`
- [ ] Verify lint passes: `make lint`

### Git Operations
- [ ] Create feature branch: `git checkout -b milestone-1-project-scaffold`
- [ ] Commit all Milestone 1 changes
- [ ] Push branch: `git push -u origin milestone-1-project-scaffold`
