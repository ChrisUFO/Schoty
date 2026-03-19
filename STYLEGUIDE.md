# Schoty Style Guide

## Overview

This style guide defines the visual design system for Schoty's TUI using Lip Gloss.

## Styling Philosophy

- **Consistency**: Same styles applied everywhere for element type
- **Clarity**: High contrast, readable at various terminal sizes
- **Brevity**: Compact representations, no wasted space
- **Terminal-native**: Works in both light and dark terminals

## Color Definitions

### Brand Colors

```go
// Primary brand color for headers and emphasis
var BrandColor = lipgloss.Color("#0066CC")

// Brand for dark terminals
var BrandColorDark = lipgloss.Color("#4DA6FF")
```

### Status Colors

```go
// Healthy status (>50% remaining)
var StatusHealthy = lipgloss.Color("#008800")
var StatusHealthyDark = lipgloss.Color("#00CC00")

// Warning status (20-50% remaining)
var StatusWarning = lipgloss.Color("#CC8800")
var StatusWarningDark = lipgloss.Color("#FFB300")

// Critical status (<20% remaining)
var StatusCritical = lipgloss.Color("#CC0000")
var StatusCriticalDark = lipgloss.Color("#FF4444")

// Error status (API failure)
var StatusError = lipgloss.Color("#8800CC")
var StatusErrorDark = lipgloss.Color("#AA66FF")

// Loading state
var StatusLoading = lipgloss.Color("#666666")
var StatusLoadingDark = lipgloss.Color("#999999")
```

### Background Colors

```go
// Light terminal backgrounds
var BackgroundLight = lipgloss.Color("#FFFFFF")
var HeaderBgLight = lipgloss.Color("#E5E5E5")
var CardBgLight = lipgloss.Color("#F5F5F5")

// Dark terminal backgrounds
var BackgroundDark = lipgloss.Color("#1A1A1A")
var HeaderBgDark = lipgloss.Color("#2D2D2D")
var CardBgDark = lipgloss.Color("#2D2D2D")
```

### Text Colors

```go
// Light terminal text
var TextPrimaryLight = lipgloss.Color("#1A1A1A")
var TextSecondaryLight = lipgloss.Color("#666666")
var TextMutedLight = lipgloss.Color("#999999")

// Dark terminal text
var TextPrimaryDark = lipgloss.Color("#E5E5E5")
var TextSecondaryDark = lipgloss.Color("#AAAAAA")
var TextMutedDark = lipgloss.Color("#888888")
```

## Typography Styles

```go
// Header: 16pt bold
var HeaderStyle = lipgloss.NewStyle().
    Bold(true).
    FontSize(16)

// Subheader: 12pt bold
var SubheaderStyle = lipgloss.NewStyle().
    Bold(true).
    FontSize(12)

// Body: 12pt regular
var BodyStyle = lipgloss.NewStyle().
    FontSize(12)

// Caption: 10pt
var CaptionStyle = lipgloss.NewStyle().
    FontSize(10)

// Status: 8pt monospace
var StatusStyle = lipgloss.NewStyle().
    FontSize(8).
    Foreground(StatusHealthy)
```

## Component Styles

### Header Style

```go
var HeaderStyle = lipgloss.NewStyle().
    Background(HeaderBgLight).
    Foreground(TextPrimaryLight).
    Padding(0, 1).
    Width(width).
    Bold(true)
```

### Tab Styles

```go
var TabActiveStyle = lipgloss.NewStyle().
    Background(BrandColor).
    Foreground(BackgroundLight).
    Padding(0, 2).
    Bold(true)

var TabInactiveStyle = lipgloss.NewStyle().
    Background(BackgroundLight).
    Foreground(TextSecondaryLight).
    Padding(0, 2)
```

### Provider Card Style

```go
var CardStyle = lipgloss.NewStyle().
    Background(CardBgLight).
    Border(lipgloss.RoundedBorder()).
    Padding(1, 2).
    Width(40)

var CardTitleStyle = lipgloss.NewStyle().
    Bold(true).
    Foreground(TextPrimaryLight)

var CardValueStyle = lipgloss.NewStyle().
    Foreground(TextPrimaryLight).
    Align(lipgloss.Right)
```

### Progress Bar Style

```go
// Width: 10 characters + percentage
// Filled: █ (block character)
// Empty: ░

func ProgressBar(percent float64, width int) string {
    filled := int(float64(width) * percent / 100)
    empty := width - filled
    return strings.Repeat("█", filled) + strings.Repeat("░", empty)
}
```

### Footer Style

```go
var FooterStyle = lipgloss.NewStyle().
    Background(HeaderBgLight).
    Foreground(TextSecondaryLight).
    Padding(0, 1).
    Width(width)
```

### Error Style

```go
var ErrorStyle = lipgloss.NewStyle().
    Foreground(StatusError).
    Italic(true)
```

### Loading Spinner Style

```go
// Animated: ◐ ◓ ◑ ◒ (rotating)
var SpinnerFrames = []string{"◐", "◓", "◑", "◒"}
```

## Layout Utilities

### Screen Dimensions

```go
func GetScreenSize() (int, int) {
    size, _ := tea.GetTermSize()
    return size.Width, size.Height
}
```

### Padding Helpers

```go
const (
    ScreenPadding = 1
    SectionGap = 1
    CardPaddingH = 2
    CardPaddingV = 1
)
```

### Width Constants

```go
const (
    MinWidth = 60
    CardWidth = 40
    ProgressBarWidth = 10
)
```

## Dark/Light Mode Detection

```go
func DetectColorMode() string {
    // Check terminal background color or environment
    // Default to light if detection fails
    return "light"
}

var IsDarkMode = DetectColorMode() == "dark"
```

## Animation

### Spinner Animation

```go
func SpinnerTick(frame int) string {
    return SpinnerFrames[frame%len(SpinnerFrames)]
}
```

### Refresh Indicator

```go
var RefreshIndicatorStyle = lipgloss.NewStyle().
    Foreground(StatusLoading).
    SetString("⟳")
```

## Usage Examples

### Creating a Provider Card

```go
func RenderProviderCard(provider Provider, balance float64, percent float64) string {
    status := GetStatusColor(percent)
    
    card := lipgloss.NewStyle().
        Border(lipgloss.RoundedBorder()).
        Padding(1, 2).
        Width(CardWidth).
        Render(
            lipgloss.JoinVertical(
                lipgloss.Left,
                provider.Name()+"  "+formatCurrency(balance),
                ProgressBar(percent, ProgressBarWidth)+" "+fmt.Sprintf("%d%%", int(percent)),
            ),
        )
    
    return card
}
```

### Status Color Selection

```go
func GetStatusColor(percent float64) lipgloss.Color {
    switch {
    case percent >= 50:
        return StatusHealthy
    case percent >= 20:
        return StatusWarning
    default:
        return StatusCritical
    }
}
```

## Best Practices

1. **Always define both light and dark mode styles**
2. **Use constants for repeated colors and sizes**
3. **Test at 80x24 minimum and 120x40 optimal**
4. **Ensure text has enough contrast against backgrounds**
5. **Use symbols + color for status (not color alone)**
6. **Keep card widths consistent across views**
7. **Use lipgloss.JoinHorizontal/Vertical for layouts**
8. **Prefer rounded borders for cards, sharp for headers**

## File Organization

Styles should be defined in `internal/ui/styles.go` following the structure:

```go
package ui

// Color constants
var (...)

// Style constants
var (...)

// Helper functions
func (...)

// Theme detection
var IsDarkMode bool
```

## Testing Styles

Verify styles render correctly by testing:
- Light terminal backgrounds
- Dark terminal backgrounds
- Various terminal sizes (80x24, 120x40, 200x50)
- All status states (healthy, warning, critical, error, loading)
