package ui

import (
	"fmt"
	"os"
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

var IsDarkMode bool

func DetectColorMode() bool {
	if os.Getenv("TERM_PROGRAM") == "Apple_Terminal" {
		return true
	}
	if os.Getenv("ColorTerm") == "truecolor" || os.Getenv("ColorTerm") == "24bit" {
		return true
	}
	if os.Getenv("TERM") == "xterm-256color" || os.Getenv("TERM") == "screen-256color" {
		return true
	}
	return false
}

func bgColor() lipgloss.Color {
	if os.Getenv("TERM_PROGRAM") == "Apple_Terminal" {
		return lipgloss.Color("#1A1A1A")
	}
	if os.Getenv("ColorTerm") == "truecolor" || os.Getenv("ColorTerm") == "24bit" {
		return lipgloss.Color("#1A1A1A")
	}
	return lipgloss.Color("#FFFFFF")
}

func fgColor() lipgloss.Color {
	if os.Getenv("TERM_PROGRAM") == "Apple_Terminal" {
		return lipgloss.Color("#E5E5E5")
	}
	if os.Getenv("ColorTerm") == "truecolor" || os.Getenv("ColorTerm") == "24bit" {
		return lipgloss.Color("#E5E5E5")
	}
	return lipgloss.Color("#1A1A1A")
}

func headerBgColor() lipgloss.Color {
	if os.Getenv("TERM_PROGRAM") == "Apple_Terminal" {
		return lipgloss.Color("#2D2D2D")
	}
	if os.Getenv("ColorTerm") == "truecolor" || os.Getenv("ColorTerm") == "24bit" {
		return lipgloss.Color("#2D2D2D")
	}
	return lipgloss.Color("#E5E5E5")
}

func cardBgColor() lipgloss.Color {
	if os.Getenv("TERM_PROGRAM") == "Apple_Terminal" {
		return lipgloss.Color("#2D2D2D")
	}
	if os.Getenv("ColorTerm") == "truecolor" || os.Getenv("ColorTerm") == "24bit" {
		return lipgloss.Color("#2D2D2D")
	}
	return lipgloss.Color("#F5F5F5")
}

func secondaryColor() lipgloss.Color {
	if os.Getenv("TERM_PROGRAM") == "Apple_Terminal" {
		return lipgloss.Color("#AAAAAA")
	}
	if os.Getenv("ColorTerm") == "truecolor" || os.Getenv("ColorTerm") == "24bit" {
		return lipgloss.Color("#AAAAAA")
	}
	return lipgloss.Color("#666666")
}

func mutedColor() lipgloss.Color {
	if os.Getenv("TERM_PROGRAM") == "Apple_Terminal" {
		return lipgloss.Color("#888888")
	}
	if os.Getenv("ColorTerm") == "truecolor" || os.Getenv("ColorTerm") == "24bit" {
		return lipgloss.Color("#888888")
	}
	return lipgloss.Color("#999999")
}

func brandColor() lipgloss.Color {
	if os.Getenv("TERM_PROGRAM") == "Apple_Terminal" {
		return lipgloss.Color("#4DA6FF")
	}
	if os.Getenv("ColorTerm") == "truecolor" || os.Getenv("ColorTerm") == "24bit" {
		return lipgloss.Color("#4DA6FF")
	}
	return lipgloss.Color("#0066CC")
}

func tabInactiveColor() lipgloss.Color {
	if os.Getenv("TERM_PROGRAM") == "Apple_Terminal" {
		return lipgloss.Color("#888888")
	}
	if os.Getenv("ColorTerm") == "truecolor" || os.Getenv("ColorTerm") == "24bit" {
		return lipgloss.Color("#888888")
	}
	return lipgloss.Color("#666666")
}

func statusHealthyColor() lipgloss.Color {
	if os.Getenv("TERM_PROGRAM") == "Apple_Terminal" {
		return lipgloss.Color("#00CC00")
	}
	if os.Getenv("ColorTerm") == "truecolor" || os.Getenv("ColorTerm") == "24bit" {
		return lipgloss.Color("#00CC00")
	}
	return lipgloss.Color("#008800")
}

func statusWarningColor() lipgloss.Color {
	if os.Getenv("TERM_PROGRAM") == "Apple_Terminal" {
		return lipgloss.Color("#FFB300")
	}
	if os.Getenv("ColorTerm") == "truecolor" || os.Getenv("ColorTerm") == "24bit" {
		return lipgloss.Color("#FFB300")
	}
	return lipgloss.Color("#CC8800")
}

func statusCriticalColor() lipgloss.Color {
	if os.Getenv("TERM_PROGRAM") == "Apple_Terminal" {
		return lipgloss.Color("#FF4444")
	}
	if os.Getenv("ColorTerm") == "truecolor" || os.Getenv("ColorTerm") == "24bit" {
		return lipgloss.Color("#FF4444")
	}
	return lipgloss.Color("#CC0000")
}

func statusErrorColor() lipgloss.Color {
	if os.Getenv("TERM_PROGRAM") == "Apple_Terminal" {
		return lipgloss.Color("#AA66FF")
	}
	if os.Getenv("ColorTerm") == "truecolor" || os.Getenv("ColorTerm") == "24bit" {
		return lipgloss.Color("#AA66FF")
	}
	return lipgloss.Color("#8800CC")
}

func statusLoadingColor() lipgloss.Color {
	if os.Getenv("TERM_PROGRAM") == "Apple_Terminal" {
		return lipgloss.Color("#999999")
	}
	if os.Getenv("ColorTerm") == "truecolor" || os.Getenv("ColorTerm") == "24bit" {
		return lipgloss.Color("#999999")
	}
	return lipgloss.Color("#666666")
}

func InitTheme() {
	IsDarkMode = os.Getenv("TERM_PROGRAM") == "Apple_Terminal" ||
		os.Getenv("ColorTerm") == "truecolor" ||
		os.Getenv("ColorTerm") == "24bit"
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
