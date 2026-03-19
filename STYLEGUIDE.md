# Schoty Style Guide

## Overview

This style guide defines the visual design system for Schoty's TUI using Lip Gloss.

## Styling Philosophy

- **Consistency**: Same styles applied everywhere for element type
- **Clarity**: High contrast, readable at various terminal sizes
- **Brevity**: Compact representations, no wasted space
- **Dark-first**: Optimized for dark terminal backgrounds (the common case)

## Color Definitions

All colors are optimized for dark terminal backgrounds. Most terminals use dark themes by default, so we prioritize readability on dark backgrounds.

### Brand Colors

```go
// Primary brand color for headers and emphasis
var BrandColor = lipgloss.Color("#3B82F6")
```

### Status Colors

```go
// Healthy status (>50% remaining)
var StatusHealthy = lipgloss.Color("#22C55E")

// Warning status (20-50% remaining)
var StatusWarning = lipgloss.Color("#F59E0B")

// Critical status (<20% remaining)
var StatusCritical = lipgloss.Color("#EF4444")

// Error status (API failure)
var StatusError = lipgloss.Color("#EC4899")

// Loading state
var StatusLoading = lipgloss.Color("#06B6D4")
```

### Background Colors

```go
// Dark terminal backgrounds
var Background = lipgloss.Color("#1A1A1A")
var HeaderBg = lipgloss.Color("#2D2D2D")
var CardBg = lipgloss.Color("#2D2D2D")
```

### Text Colors

```go
// Text for dark terminals
var TextPrimary = lipgloss.Color("#E5E5E5")
var TextSecondary = lipgloss.Color("#AAAAAA")
var TextMuted = lipgloss.Color("#9CA3AF")
var TabInactive = lipgloss.Color("#6B7280")
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
    Background(HeaderBg).
    Foreground(TextPrimary).
    Padding(0, 1).
    Width(width).
    Bold(true)
```

### Tab Styles

```go
var TabActiveStyle = lipgloss.NewStyle().
    Background(BrandColor).
    Foreground(Background).
    Padding(0, 2).
    Bold(true)

var TabInactiveStyle = lipgloss.NewStyle().
    Background(Background).
    Foreground(TabInactive).
    Padding(0, 2)
```

### Provider Card Style

```go
var CardStyle = lipgloss.NewStyle().
    Background(CardBg).
    Border(lipgloss.RoundedBorder()).
    Padding(1, 2).
    Width(40)

var CardTitleStyle = lipgloss.NewStyle().
    Bold(true).
    Foreground(TextPrimary)

var CardValueStyle = lipgloss.NewStyle().
    Foreground(TextPrimary).
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
    Background(HeaderBg).
    Foreground(TextSecondary).
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

1. **Always use high-contrast colors for text on dark backgrounds**
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

// Color functions (return dark-theme optimized colors)
func bgColor() lipgloss.Color { ... }
func fgColor() lipgloss.Color { ... }
func brandColor() lipgloss.Color { ... }
// ... etc

// Style functions
func HeaderStyle() lipgloss.Style { ... }
func CardStyle() lipgloss.Style { ... }
// ... etc

// Helper functions
func (...)
```

## Testing Styles

Verify styles render correctly by testing:
- Dark terminal backgrounds (primary target)
- Various terminal sizes (80x24, 120x40, 200x50)
- All status states (healthy, warning, critical, error, loading)
