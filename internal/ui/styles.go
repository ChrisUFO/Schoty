package ui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

var IsDarkMode bool

func DetectColorMode() bool {
	return false
}

var (
	BackgroundColor = lipgloss.Color("#FFFFFF")
	HeaderBgColor   = lipgloss.Color("#E5E5E5")
	CardBgColor     = lipgloss.Color("#F5F5F5")
	TextPrimary     = lipgloss.Color("#1A1A1A")
	TextSecondary   = lipgloss.Color("#666666")
	TextMuted       = lipgloss.Color("#999999")
	BrandColor      = lipgloss.Color("#0066CC")
	TabActive       = lipgloss.Color("#0066CC")
	TabInactive     = lipgloss.Color("#666666")
)

var (
	StatusHealthy  = lipgloss.Color("#008800")
	StatusWarning  = lipgloss.Color("#CC8800")
	StatusCritical = lipgloss.Color("#CC0000")
	StatusError    = lipgloss.Color("#8800CC")
	StatusLoading  = lipgloss.Color("#666666")
)

func InitTheme() {
	IsDarkMode = DetectColorMode()
	if IsDarkMode {
		BackgroundColor = lipgloss.Color("#1A1A1A")
		HeaderBgColor = lipgloss.Color("#2D2D2D")
		CardBgColor = lipgloss.Color("#2D2D2D")
		TextPrimary = lipgloss.Color("#E5E5E5")
		TextSecondary = lipgloss.Color("#AAAAAA")
		TextMuted = lipgloss.Color("#888888")
		BrandColor = lipgloss.Color("#4DA6FF")
		TabActive = lipgloss.Color("#4DA6FF")
		TabInactive = lipgloss.Color("#888888")
		StatusHealthy = lipgloss.Color("#00CC00")
		StatusWarning = lipgloss.Color("#FFB300")
		StatusCritical = lipgloss.Color("#FF4444")
		StatusError = lipgloss.Color("#AA66FF")
		StatusLoading = lipgloss.Color("#999999")
	}
}

var HeaderStyle = lipgloss.NewStyle().
	Background(HeaderBgColor).
	Foreground(TextPrimary).
	Padding(0, 1).
	Bold(true)

var TabActiveStyle = lipgloss.NewStyle().
	Background(BrandColor).
	Foreground(BackgroundColor).
	Padding(0, 2).
	Bold(true)

var TabInactiveStyle = lipgloss.NewStyle().
	Background(BackgroundColor).
	Foreground(TabInactive).
	Padding(0, 2)

var CardStyle = lipgloss.NewStyle().
	Background(CardBgColor).
	Border(lipgloss.RoundedBorder()).
	Padding(1, 2).
	Width(40)

var CardTitleStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(TextPrimary)

var CardValueStyle = lipgloss.NewStyle().
	Foreground(TextPrimary).
	Align(lipgloss.Right)

var FooterStyle = lipgloss.NewStyle().
	Background(HeaderBgColor).
	Foreground(TextSecondary).
	Padding(0, 1)

var ErrorStyle = lipgloss.NewStyle().
	Foreground(StatusError).
	Italic(true)

var BodyStyle = lipgloss.NewStyle().
	Foreground(TextPrimary)

var CaptionStyle = lipgloss.NewStyle().
	Foreground(TextMuted)

var SpinnerFrames = []string{"◐", "◓", "◑", "◒"}

func SpinnerTick(frame int) string {
	return SpinnerFrames[frame%len(SpinnerFrames)]
}

func ProgressBar(percent float64, width int) string {
	if percent < 0 {
		percent = 0
	}
	if percent > 100 {
		percent = 100
	}
	return fmt.Sprintf("%d%%", int(percent))
}

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
