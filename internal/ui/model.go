package ui

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type ViewState int

const (
	DashboardView ViewState = iota
	DetailView
	ConfigView
	HelpView
)

const numTabs = 3

type ProviderState struct {
	Name         string
	Balance      float64
	Usage        int
	Remaining    int
	Limit        int
	Status       string
	ErrorMsg     string
	IsLoading    bool
	IsConfigured bool
}

type Model struct {
	Ready            bool
	Width            int
	Height           int
	Tab              int
	CurrentView      ViewState
	SelectedProvider int
	LastRefresh      time.Time
	Providers        []ProviderState
	ShowHelp         bool
	frame            int
}

func NewModel() Model {
	InitTheme()
	return Model{
		Ready:            false,
		Width:            80,
		Height:           24,
		Tab:              0,
		CurrentView:      DashboardView,
		SelectedProvider: 0,
		LastRefresh:      time.Now(),
		Providers:        []ProviderState{},
		ShowHelp:         false,
		frame:            0,
	}
}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		m.handleKeyPress(msg)
	case tea.WindowSizeMsg:
		m.Ready = true
		m.Width = msg.Width
		m.Height = msg.Height
	case tickMsg:
		m.LastRefresh = time.Now()
	}
	return m, nil
}

func (m *Model) handleKeyPress(msg tea.KeyMsg) {
	switch msg.String() {
	case "q", "ctrl+c":
		return
	case "?":
		m.ShowHelp = !m.ShowHelp
	case "c":
		if m.CurrentView == ConfigView {
			m.CurrentView = DashboardView
		} else {
			m.CurrentView = ConfigView
		}
	case "r":
		m.refreshAll()
	case "tab":
		m.Tab = (m.Tab + 1) % 3
	case "1", "2", "3", "4", "5", "6", "7", "8":
		idx := int(msg.String()[0] - '1')
		if idx < len(m.Providers) {
			m.SelectedProvider = idx
			m.CurrentView = DetailView
		}
	case "enter":
		if len(m.Providers) > 0 && m.CurrentView == DashboardView {
			m.CurrentView = DetailView
		}
	case "esc":
		if m.ShowHelp {
			m.ShowHelp = false
		} else if m.CurrentView == DetailView {
			m.CurrentView = DashboardView
		} else if m.CurrentView == ConfigView {
			m.CurrentView = DashboardView
		}
	case "up", "k":
		if m.SelectedProvider > 0 {
			m.SelectedProvider--
		}
	case "down", "j":
		if m.SelectedProvider < len(m.Providers)-1 {
			m.SelectedProvider++
		}
	case "left", "h":
		if m.Tab > 0 {
			m.Tab--
		}
	case "right", "l":
		if m.Tab < 2 {
			m.Tab++
		}
	case "d":
		if m.CurrentView == DashboardView && len(m.Providers) > 0 {
			m.CurrentView = DetailView
		}
	}
}

type tickMsg time.Time

func (m *Model) refreshAll() {
	for i := range m.Providers {
		m.Providers[i].IsLoading = true
	}
}

func (m *Model) View() string {
	if !m.Ready {
		return lipgloss.NewStyle().
			Width(m.Width).
			Height(m.Height).
			Align(lipgloss.Center, lipgloss.Center).
			Render("Initializing Schoty...")
	}

	if m.Width < MinWidth || m.Height < 10 {
		return m.renderTooSmallView()
	}

	if m.ShowHelp {
		return m.renderHelpView()
	}

	switch m.CurrentView {
	case ConfigView:
		return m.renderConfigView()
	case DetailView:
		return m.renderDetailView()
	default:
		return m.renderDashboard()
	}
}

func (m *Model) appStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Width(m.Width).
		Height(m.Height)
}

func (m *Model) renderTooSmallView() string {
	msg := ErrorStyle.Render("Terminal too small. Please resize to at least 60x10")
	return lipgloss.NewStyle().
		Width(m.Width).
		Height(m.Height).
		Align(lipgloss.Center, lipgloss.Center).
		Render(msg)
}

func (m *Model) renderHeader() string {
	refreshStr := fmt.Sprintf("Last refresh: %s", m.LastRefresh.Format("15:04:05"))
	title := lipgloss.NewStyle().
		Background(HeaderBgColor).
		Foreground(TextPrimary).
		Bold(true).
		Padding(0, 1).
		Render(" Schoty - AI Monitor ")

	refresh := lipgloss.NewStyle().
		Background(HeaderBgColor).
		Foreground(TextSecondary).
		Padding(0, 1).
		Render(refreshStr)

	content := lipgloss.JoinHorizontal(
		lipgloss.Left,
		title,
		refresh,
	)

	bordered := lipgloss.NewStyle().
		Border(lipgloss.Border{Bottom: "─"}).
		Width(m.Width).
		Render(content)

	return bordered
}

func (m *Model) renderTabs() string {
	tabs := []string{"API Balances", "Subscriptions", "Config"}
	result := ""

	for i, tab := range tabs {
		isActive := i == m.Tab
		bg := BrandColor
		fg := BackgroundColor
		if !isActive {
			bg = BackgroundColor
			fg = TabInactive
		}
		style := lipgloss.NewStyle().
			Background(bg).
			Foreground(fg).
			Padding(0, 2).
			Bold(isActive)

		bracketFg := BrandColor
		if isActive {
			bracketFg = BackgroundColor
		}
		bracketLeft := lipgloss.NewStyle().
			Background(bg).
			Foreground(bracketFg).
			Render("[")

		bracketRight := lipgloss.NewStyle().
			Background(bg).
			Foreground(bracketFg).
			Render("]")

		result += bracketLeft + style.Render(tab) + bracketRight + " "
	}

	return result
}

func (m *Model) renderDashboard() string {
	var content string

	if len(m.Providers) == 0 {
		content = m.renderEmptyState()
	} else {
		content = m.renderProviderList()
	}

	footer := m.renderFooter()

	borderBox := lipgloss.NewStyle().
		Border(lipgloss.Border{Top: "─"}).
		Width(m.Width).
		Render(footer)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		m.renderHeader(),
		m.renderTabs(),
		content,
		borderBox,
	)
}

func (m *Model) renderProviderList() string {
	var rows []string

	for i, p := range m.Providers {
		row := m.renderProviderRow(p, i)
		rows = append(rows, row)
		if i < len(m.Providers)-1 {
			rows = append(rows, "")
		}
	}

	lines := lipgloss.JoinVertical(
		lipgloss.Left,
		rows...,
	)

	contentStyle := lipgloss.NewStyle().
		Width(m.Width).
		Height(m.Height-6).
		Align(lipgloss.Left, lipgloss.Top)

	return contentStyle.Render(lines)
}

func (m *Model) renderProviderRow(p ProviderState, idx int) string {
	status := StatusIndicator(p.Status)
	nameStyle := lipgloss.NewStyle().Foreground(TextPrimary).Bold(idx == m.SelectedProvider)

	name := nameStyle.Render(p.Name)

	var valueStr string
	var progressStr string

	if p.IsLoading {
		valueStr = CaptionStyle.Render("Fetching...")
		progressStr = SpinnerTick(m.frame)
	} else if p.ErrorMsg != "" {
		valueStr = ErrorStyle.Render(p.ErrorMsg)
		progressStr = ""
	} else if p.IsConfigured {
		valueStr = fmt.Sprintf("$%.2f remaining", p.Balance)
		percent := float64(p.Remaining*100) / float64(p.Limit)
		progressStr = ProgressBarSimple(percent)
	} else {
		valueStr = CaptionStyle.Render("Not configured")
		progressStr = CaptionStyle.Render("—")
	}

	row := lipgloss.JoinHorizontal(
		lipgloss.Left,
		status,
		" "+name,
		lipgloss.NewStyle().Width(m.Width-40).Align(lipgloss.Right).Render(valueStr),
	)

	progress := lipgloss.JoinHorizontal(
		lipgloss.Left,
		"  "+CaptionStyle.Render("API Balance"),
		"  "+CaptionStyle.Render(progressStr),
	)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		row,
		progress,
	)
}

func (m *Model) renderEmptyState() string {
	box := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Padding(3, 5).
		Width(45).
		Align(lipgloss.Center, lipgloss.Center).
		BorderForeground(BrandColor).
		Render(
			lipgloss.JoinVertical(
				lipgloss.Center,
				BodyStyle.Render("No providers configured"),
				"",
				CaptionStyle.Render("Press [c] to add API keys"),
				CaptionStyle.Render("and configure your providers"),
			),
		)

	centered := lipgloss.NewStyle().
		Width(m.Width).
		Height(m.Height-6).
		Align(lipgloss.Center, lipgloss.Center).
		Render(box)

	return centered
}

func (m *Model) renderDetailView() string {
	if m.SelectedProvider >= len(m.Providers) {
		m.CurrentView = DashboardView
		return m.renderDashboard()
	}

	p := m.Providers[m.SelectedProvider]

	header := lipgloss.NewStyle().
		Foreground(BrandColor).
		Bold(true).
		Render(fmt.Sprintf(" %s ", p.Name))

	borderTop := lipgloss.NewStyle().
		Border(lipgloss.Border{Bottom: "─"}).
		Width(m.Width).
		Render(header)

	balance := fmt.Sprintf("Balance:     $%.2f", p.Balance)
	usage := fmt.Sprintf("Usage:       %d / %d", p.Limit-p.Remaining, p.Limit)
	remaining := fmt.Sprintf("Remaining:   %d", p.Remaining)
	lastUpdated := fmt.Sprintf("Last Updated: %s", m.LastRefresh.Format("15:04:05"))

	percent := float64(p.Remaining*100) / float64(p.Limit)
	progress := ProgressBarSimple(percent)

	details := lipgloss.JoinVertical(
		lipgloss.Left,
		BodyStyle.Render(balance),
		BodyStyle.Render(usage),
		BodyStyle.Render(remaining),
		"",
		CaptionStyle.Render(progress),
		"",
		CaptionStyle.Render(lastUpdated),
	)

	back := CaptionStyle.Render("[←] Back   [r] Refresh   [e] Edit Config")

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		borderTop,
		"",
		details,
		"",
		"",
		back,
	)

	centered := lipgloss.NewStyle().
		Width(m.Width).
		Height(m.Height-6).
		Align(lipgloss.Center, lipgloss.Center).
		Render(content)

	footer := m.renderFooter()

	borderFooter := lipgloss.NewStyle().
		Border(lipgloss.Border{Top: "─"}).
		Width(m.Width).
		Render(footer)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		m.renderHeader(),
		m.renderTabs(),
		centered,
		borderFooter,
	)
}

func (m *Model) renderConfigView() string {
	title := lipgloss.NewStyle().
		Foreground(BrandColor).
		Bold(true).
		Render(" Configuration ")

	borderTop := lipgloss.NewStyle().
		Border(lipgloss.Border{Bottom: "─"}).
		Width(m.Width).
		Render(title)

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		"",
		BodyStyle.Render("Provider configuration"),
		CaptionStyle.Render("(Coming in Milestone 3)"),
	)

	centered := lipgloss.NewStyle().
		Width(m.Width).
		Height(m.Height-6).
		Align(lipgloss.Center, lipgloss.Center).
		Render(content)

	footer := m.renderFooter()

	borderFooter := lipgloss.NewStyle().
		Border(lipgloss.Border{Top: "─"}).
		Width(m.Width).
		Render(footer)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		borderTop,
		m.renderTabs(),
		centered,
		borderFooter,
	)
}

func (m *Model) renderHelpView() string {
	helpText := `Keyboard Shortcuts:

Global:
  q, ctrl+c    Quit application
  r            Refresh all data
  c            Toggle config view
  ?            Toggle this help

Navigation:
  tab          Cycle through tabs
  1-8          Quick jump to provider
  up/down      Navigate list
  left/right   Switch tabs

Views:
  enter        Open detail view
  esc          Go back / close
  d            Toggle detail mode`

	helpBox := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Padding(2, 4).
		BorderForeground(BrandColor).
		Render(
			lipgloss.JoinVertical(
				lipgloss.Left,
				CardTitleStyle.Render("Keyboard Shortcuts"),
				BodyStyle.Render(helpText),
			),
		)

	centered := lipgloss.NewStyle().
		Width(m.Width).
		Height(m.Height-4).
		Align(lipgloss.Center, lipgloss.Center).
		Render(helpBox)

	footer := m.renderFooter()

	borderFooter := lipgloss.NewStyle().
		Border(lipgloss.Border{Top: "─"}).
		Width(m.Width).
		Render(footer)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		m.renderHeader(),
		"",
		centered,
		borderFooter,
	)
}

func (m *Model) renderFooter() string {
	status := fmt.Sprintf("[q] quit  [r] refresh  [tab] tabs  [c] config  [?] help  │ %d providers", len(m.Providers))
	return lipgloss.NewStyle().
		Background(HeaderBgColor).
		Foreground(TextSecondary).
		Width(m.Width).
		Padding(0, 1).
		Render(status)
}
