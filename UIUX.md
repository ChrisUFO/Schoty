# Schoty UI/UX Specification

## Overview

Schoty is a terminal-native TUI application for monitoring AI subscription usage and API balances. The design prioritizes information density, quick scanning, and keyboard-driven navigation.

## Design Principles

1. **Information Density** - Show maximum relevant data without clutter
2. **Scanability** - Status visible at a glance from across the room
3. **Keyboard-First** - All actions accessible without mouse
4. **Graceful Degradation** - Useful even when APIs are unavailable
5. **Responsive Feedback** - Every action has immediate visual response

## Layout Architecture

### Screen Structure

```
┌─────────────────────────────────────────────────────────────┐
│  HEADER: Title | Status Indicator | Last Refresh Timestamp   │
├─────────────────────────────────────────────────────────────┤
│  TABS: [API Balances] [Subscriptions] [Config]              │
├─────────────────────────────────────────────────────────────┤
│                                                             │
│  MAIN CONTENT AREA                                          │
│  - Provider cards/list in selected tab                      │
│  - Scrollable if content exceeds viewport                  │
│                                                             │
├─────────────────────────────────────────────────────────────┤
│  FOOTER: Keyboard Hints | Provider Count | Refresh Status   │
└─────────────────────────────────────────────────────────────┘
```

### View Types

#### 1. Dashboard View (Default)
- Grid/list of all configured providers
- Shows: Provider name, balance/usage, status indicator
- Click/Enter to drill into detail view
- Color-coded status (healthy/warning/critical)

#### 2. Detail View
- Single provider expanded information
- Historical data if available
- Refresh button
- Error details if failed

#### 3. Config View
- List of providers with enable/disable toggle
- API key input (masked)
- Environment variable hints
- Save/Cancel actions

## Navigation

### Global Hotkeys

| Key | Action |
|-----|--------|
| `q` / `Ctrl+C` | Quit application |
| `r` | Refresh all data |
| `1-8` | Quick jump to provider |
| `Tab` | Cycle through tabs |
| `Esc` | Back to previous view / Close modal |
| `c` | Open config view |
| `?` | Show keyboard shortcuts help |

### Tab-Specific Navigation

| Key | Action |
|-----|--------|
| `↑` / `↓` | Navigate list items |
| `Enter` | Select item / Drill down |
| `←` / `→` | Switch tabs |
| `d` | Toggle detail view mode |

## Visual Design

### Color Scheme

#### Light Terminal
| Element | Color | Usage |
|---------|-------|-------|
| Background | `#FFFFFF` | Main background |
| Primary Text | `#1A1A1A` | Body text |
| Header BG | `#E5E5E5` | Header area |
| Tab Active | `#0066CC` | Active tab |
| Tab Inactive | `#666666` | Inactive tab |

#### Dark Terminal
| Element | Color | Usage |
|---------|-------|-------|
| Background | `#1A1A1A` | Main background |
| Primary Text | `#E5E5E5` | Body text |
| Header BG | `#2D2D2D` | Header area |
| Tab Active | `#4DA6FF` | Active tab |
| Tab Inactive | `#888888` | Inactive tab |

### Status Colors

| Status | Light | Dark | Meaning |
|--------|-------|------|---------|
| Healthy | `#008800` | `#00CC00` | >50% remaining |
| Warning | `#CC8800` | `#FFB300` | 20-50% remaining |
| Critical | `#CC0000` | `#FF4444` | <20% remaining |
| Error | `#8800CC` | `#AA66FF` | API failed |
| Loading | `#666666` | `#999999` | Fetching data |

### Typography

- **Font**: Terminal default (monospace)
- **Header Size**: 16pt bold
- **Body Size**: 12pt regular
- **Footer Size**: 10pt
- **Status Indicators**: 8pt

### Spacing System

| Element | Value |
|---------|-------|
| Screen Padding | 1 character |
| Section Gap | 1 line |
| Card Padding | 2 characters horizontal, 1 vertical |
| List Item Gap | 0 lines |

## Component Specifications

### Provider Card

```
┌─────────────────────────────────────────┐
│ ● Anthropic           $12.45 remaining  │
│   API Balance         ████████░░ 80%   │
└─────────────────────────────────────────┘
```

- Status indicator dot (color per status)
- Provider name (left-aligned)
- Balance/usage value (right-aligned)
- Progress bar showing percentage
- Subtle border or background differentiation

### Detail Panel

```
┌─────────────────────────────────────────┐
│  ANTHROPIC                              │
│                                         │
│  Balance:     $12.45                    │
│  Daily Avg:   $1.23                     │
│  Days Left:   ~10 (at current rate)     │
│                                         │
│  Last Updated: 2 minutes ago            │
│                                         │
│  [r] Refresh  [e] Edit Config  [←] Back │
└─────────────────────────────────────────┘
```

### Loading State

```
○ Anthropic        Fetching...
○ OpenAI           Fetching...
○ OpenRouter       Fetching...
```

- Spinning indicator (○) animation
- "Fetching..." text
- Stops on error with error indicator

### Error State

```
✗ Anthropic        Error: Invalid API key
✗ OpenAI            $24.50 remaining
✗ OpenRouter       $123.45 remaining
```

- Error indicator (✗)
- Error message on same line
- Other providers still show data

### Empty State (No Providers Configured)

```
┌─────────────────────────────────────────┐
│                                         │
│         No providers configured         │
│                                         │
│    Press [c] to add API keys and        │
│    configure your providers             │
│                                         │
└─────────────────────────────────────────┘
```

## User Workflows

### First Launch Flow

1. App starts → Empty state displayed
2. User presses `c` → Config view opens
3. User enters API keys → Keys validated on save
4. User returns to dashboard → Data fetches automatically
5. Dashboard shows all configured providers

### Normal Usage Flow

1. User runs `schoty`
2. Dashboard displays with cached data
3. Background refresh triggers automatically
4. User glances at status
5. If attention needed → drill down with Enter
6. Press `q` to exit

### Refresh Flow

1. User presses `r` or auto-refresh triggers
2. All providers show loading state
3. Data returns progressively
4. UI updates reactively
5. Footer shows "Last refresh: just now"

### Error Recovery Flow

1. Provider API call fails
2. Error state shown with message
3. User can press `r` to retry single provider
4. Or wait for auto-refresh
5. Stale data remains visible with timestamp

## Responsive Behavior

### Viewport < 80 columns
- Single column layout
- Compact cards
- Truncated provider names if needed

### Viewport >= 80 columns
- Two-column grid for cards
- Full provider names
- Expanded detail views

### Viewport < 24 rows
- Header condensed
- Footer on single line
- Minimal padding

## Accessibility Considerations

- Color is not sole indicator (use symbols + color)
- High contrast between text and background
- Status perceivable without exact color distinction
- All actions keyboard-accessible

## Technical Implementation

### Full-Window TUI Behavior

Schoty takes over the entire terminal window using Bubble Tea's full-screen renderer:

- **Renderer**: Uses Bubble Tea's built-in terminal renderer for ANSI/VT100 compatible terminals
- **Window Size**: Dynamically adapts to terminal dimensions via `tea.WindowSizeMsg`
- **Minimum Size**: Warns user if terminal is smaller than 60x10
- **Resize Handling**: Re-renders automatically on window resize events
- **Cursor**: Hidden during TUI display for cleaner appearance

### Rendering Pipeline

1. `tea.Program` initializes with `Model`
2. `Model.Update()` handles window resize and keyboard events
3. `Model.View()` renders the complete UI string for each frame
4. Bubble Tea's renderer draws the UI and handles cursor management

### Style Application

Styles use the actual terminal dimensions:

```go
func (m *Model) appStyle() lipgloss.Style {
    return lipgloss.NewStyle().
        Width(m.Width).
        Height(m.Height)
}
```

Each component is styled relative to the current viewport, ensuring the UI fills the entire terminal.

## Future Enhancements

- [ ] Customizable refresh intervals
- [ ] Multiple theme support
- [ ] Sort/filter providers
- [ ] Data export
- [ ] Alert notifications
- [ ] History graphs
