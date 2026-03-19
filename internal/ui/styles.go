package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

const (
	ScreenPadding    = 1
	SectionGap       = 1
	CardPaddingH     = 2
	CardPaddingV     = 1
	MinWidth         = 60
	CardWidth        = 40
	ProgressBarWidth = 10
)

func bgColor() lipgloss.Color {
	return lipgloss.Color("#1A1A1A")
}

func fgColor() lipgloss.Color {
	return lipgloss.Color("#E5E5E5")
}

func headerBgColor() lipgloss.Color {
	return lipgloss.Color("#2D2D2D")
}

func cardBgColor() lipgloss.Color {
	return lipgloss.Color("#2D2D2D")
}

func secondaryColor() lipgloss.Color {
	return lipgloss.Color("#AAAAAA")
}

func mutedColor() lipgloss.Color {
	return lipgloss.Color("#9CA3AF")
}

func brandColor() lipgloss.Color {
	return lipgloss.Color("#EA580C")
}

func tabInactiveColor() lipgloss.Color {
	return lipgloss.Color("#6B7280")
}

func statusHealthyColor() lipgloss.Color {
	return lipgloss.Color("#22C55E")
}

func statusWarningColor() lipgloss.Color {
	return lipgloss.Color("#FACC15")
}

func statusCriticalColor() lipgloss.Color {
	return lipgloss.Color("#EF4444")
}

func statusErrorColor() lipgloss.Color {
	return lipgloss.Color("#FB7185")
}

func statusLoadingColor() lipgloss.Color {
	return lipgloss.Color("#78716C")
}

func HeaderStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Background(headerBgColor()).
		Foreground(fgColor()).
		Padding(0, 1).
		Bold(true)
}

func TabActiveStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Background(brandColor()).
		Foreground(bgColor()).
		Padding(0, 2).
		Bold(true)
}

func TabInactiveStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Background(bgColor()).
		Foreground(tabInactiveColor()).
		Padding(0, 2)
}

func CardStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Background(cardBgColor()).
		Border(lipgloss.RoundedBorder()).
		Padding(1, 2).
		Width(CardWidth)
}

func CardTitleStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Bold(true).
		Foreground(fgColor())
}

func CardValueStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(fgColor()).
		Align(lipgloss.Right)
}

func FooterStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Background(headerBgColor()).
		Foreground(secondaryColor()).
		Padding(0, 1)
}

func ErrorStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(statusErrorColor()).
		Italic(true)
}

func BodyStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(fgColor())
}

func CaptionStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(mutedColor())
}

var SpinnerFrames = []string{"◐", "◓", "◑", "◒"}

func SpinnerTick(frame int) string {
	return SpinnerFrames[frame%len(SpinnerFrames)]
}

func ProgressBarSimple(percent float64) string {
	if percent < 0 {
		percent = 0
	}
	if percent > 100 {
		percent = 100
	}
	filled := int(float64(ProgressBarWidth) * percent / 100)
	empty := ProgressBarWidth - filled
	filledStr := strings.Repeat("█", filled)
	emptyStr := strings.Repeat("░", empty)
	return fmt.Sprintf("%s%s %d%%", filledStr, emptyStr, int(percent))
}

func GetStatusColor(percent float64) lipgloss.Color {
	switch {
	case percent >= 50:
		return statusHealthyColor()
	case percent >= 20:
		return statusWarningColor()
	default:
		return statusCriticalColor()
	}
}

func StatusIndicator(state string) string {
	style := lipgloss.NewStyle()
	switch state {
	case "healthy":
		style.Foreground(statusHealthyColor())
		return style.Render("●")
	case "warning":
		style.Foreground(statusWarningColor())
		return style.Render("●")
	case "critical":
		style.Foreground(statusCriticalColor())
		return style.Render("●")
	case "error":
		style.Foreground(statusErrorColor())
		return style.Render("✗")
	case "loading":
		style.Foreground(statusLoadingColor())
		return style.Render("○")
	default:
		style.Foreground(statusLoadingColor())
		return style.Render("○")
	}
}

func GetScreenSize() (int, int) {
	return 80, 24
}
