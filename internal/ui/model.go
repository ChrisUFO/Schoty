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
}

func NewModel() Model {
	InitTheme()
	width, height := GetScreenSize()
	return Model{
		Ready:            false,
		Width:            width,
		Height:           height,
		Tab:              0,
		CurrentView:      DashboardView,
		SelectedProvider: 0,
		LastRefresh:      time.Now(),
		Providers:        []ProviderState{},
		ShowHelp:         false,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		return m.handleKeyPress(msg)
	case tea.WindowSizeMsg:
		m.Ready = true
		m.Width = msg.Width
		m.Height = msg.Height
	case tickMsg:
		m.LastRefresh = time.Now()
	}
	return m, nil
}

func (m Model) handleKeyPress(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "q", "ctrl+c":
		return m, tea.Quit
	case "?":
		m.ShowHelp = !m.ShowHelp
	case "c":
		if m.CurrentView == ConfigView {
			m.CurrentView = DashboardView
		} else {
			m.CurrentView = ConfigView
		}
	case "r":
		return m, m.refreshAll()
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
	return m, nil
}

type tickMsg time.Time

func (m Model) refreshAll() tea.Cmd {
	return func() tea.Msg {
		time.Sleep(100 * time.Millisecond)
		return tickMsg(time.Now())
	}
}

func (m Model) View() string {
	if !m.Ready {
		return "Initializing Schoty..."
	}

	if m.Width < MinWidth || m.Height < 10 {
		return lipgloss.JoinVertical(
			lipgloss.Left,
			HeaderStyle.Render("Schoty - AI Monitor"),
			"",
			ErrorStyle.Render("Terminal too small. Please resize to at least 60x10"),
		)
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

func (m Model) renderHeader() string {
	refreshStr := fmt.Sprintf("Last refresh: %s", m.LastRefresh.Format("15:04:05"))
	headerContent := lipgloss.JoinHorizontal(
		lipgloss.Left,
		HeaderStyle.Render("Schoty - AI Monitor"),
		HeaderStyle.Width(m.Width-30).Render(refreshStr),
	)
	return headerContent
}

func (m Model) renderTabs() string {
	tabs := []string{"[API Balances]", "[Subscriptions]", "[Config]"}
	tabStr := ""
	for i, tab := range tabs {
		if i == m.Tab {
			tabStr += TabActiveStyle.Render(tab)
		} else {
			tabStr += TabInactiveStyle.Render(tab)
		}
		tabStr += " "
	}
	return tabStr
}

func (m Model) renderDashboard() string {
	var content string

	if len(m.Providers) == 0 {
		content = m.renderEmptyState()
	} else {
		content = m.renderProviderList()
	}

	footer := FooterStyle.Width(m.Width).Render(
		fmt.Sprintf("[q] quit  [r] refresh  [tab] tabs  [c] config  [?] help  │ %d providers", len(m.Providers)),
	)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		m.renderHeader(),
		m.renderTabs(),
		"",
		content,
		"",
		footer,
	)
}

func (m Model) renderProviderList() string {
	var lines []string

	for i, p := range m.Providers {
		if p.IsLoading {
			status := StatusIndicator("loading")
			nameStyle := lipgloss.NewStyle().Foreground(TextSecondary)
			lines = append(lines, fmt.Sprintf("%s %s  %s", status, nameStyle.Render(p.Name), CaptionStyle.Render("Fetching...")))
		} else if p.ErrorMsg != "" {
			status := StatusIndicator("error")
			lines = append(lines, fmt.Sprintf("%s %s  %s", status, CardTitleStyle.Render(p.Name), ErrorStyle.Render(p.ErrorMsg)))
		} else if p.IsConfigured {
			status := StatusIndicator(p.Status)
			balanceStr := fmt.Sprintf("$%.2f remaining", p.Balance)
			progress := ProgressBarSimple(float64(p.Remaining*100) / float64(p.Limit))

			row := lipgloss.JoinHorizontal(
				lipgloss.Left,
				status,
				CardTitleStyle.Render(p.Name),
				lipgloss.NewStyle().Width(15).Align(lipgloss.Right).Render(balanceStr),
			)
			lines = append(lines, row)
			lines = append(lines, lipgloss.JoinHorizontal(
				lipgloss.Left,
				"  "+CaptionStyle.Render("API Balance"),
				"  "+CaptionStyle.Render(progress),
			))
		} else {
			status := StatusIndicator("loading")
			lines = append(lines, fmt.Sprintf("%s %s  %s", status, CardTitleStyle.Render(p.Name), CaptionStyle.Render("Not configured")))
		}

		if i < len(m.Providers)-1 {
			lines = append(lines, "")
		}
	}

	return lipgloss.JoinVertical(
		lipgloss.Left,
		lines...,
	)
}

func (m Model) renderEmptyState() string {
	emptyBox := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Padding(2, 4).
		Width(40).
		Align(lipgloss.Center).
		Render(
			lipgloss.JoinVertical(
				lipgloss.Center,
				BodyStyle.Render("No providers configured"),
				"",
				CaptionStyle.Render("Press [c] to add API keys"),
				CaptionStyle.Render("and configure your providers"),
			),
		)

	return lipgloss.JoinVertical(
		lipgloss.Center,
		"",
		emptyBox,
	)
}

func (m Model) renderDetailView() string {
	if m.SelectedProvider >= len(m.Providers) {
		m.CurrentView = DashboardView
		return m.renderDashboard()
	}

	p := m.Providers[m.SelectedProvider]

	header := CardTitleStyle.Render(p.Name)
	balance := fmt.Sprintf("Balance: $%.2f", p.Balance)
	usage := fmt.Sprintf("Usage: %d / %d", p.Limit-p.Remaining, p.Limit)
	remaining := fmt.Sprintf("Remaining: %d", p.Remaining)
	lastUpdated := fmt.Sprintf("Last Updated: %s", m.LastRefresh.Format("15:04:05"))

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		header,
		"",
		BodyStyle.Render(balance),
		BodyStyle.Render(usage),
		BodyStyle.Render(remaining),
		"",
		CaptionStyle.Render(lastUpdated),
	)

	back := CaptionStyle.Render("[←] Back  [r] Refresh  [e] Edit Config")

	footer := FooterStyle.Width(m.Width).Render(
		fmt.Sprintf("[q] quit  [esc] back  [r] refresh  [?] help"),
	)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		m.renderHeader(),
		m.renderTabs(),
		"",
		content,
		"",
		back,
		"",
		footer,
	)
}

func (m Model) renderConfigView() string {
	title := CardTitleStyle.Render("Configuration")

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		"",
		BodyStyle.Render("Provider configuration"),
		CaptionStyle.Render("(Coming in Milestone 3)"),
	)

	footer := FooterStyle.Width(m.Width).Render(
		fmt.Sprintf("[q] quit  [esc] back  [?] help"),
	)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		m.renderHeader(),
		m.renderTabs(),
		"",
		content,
		"",
		footer,
	)
}

func (m Model) renderHelpView() string {
	helpText := `
Keyboard Shortcuts:

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
  d            Toggle detail mode
`

	helpBox := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Padding(2, 4).
		Render(
			lipgloss.JoinVertical(
				lipgloss.Left,
				CardTitleStyle.Render("Keyboard Shortcuts"),
				BodyStyle.Render(helpText),
			),
		)

	footer := FooterStyle.Width(m.Width).Render(
		fmt.Sprintf("[q] quit  [esc] close help"),
	)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		m.renderHeader(),
		"",
		helpBox,
		"",
		footer,
	)
}
