package ui

import (
	"github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	Ready     bool
	Tab       int
	Providers []string
}

func NewModel() Model {
	InitTheme()
	return Model{
		Ready:     false,
		Tab:       0,
		Providers: []string{},
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "tab":
			m.Tab = (m.Tab + 1) % 3
		case "r":
		}
	case tea.WindowSizeMsg:
		m.Ready = true
	}
	return m, nil
}

func (m Model) View() string {
	if !m.Ready {
		return "Initializing..."
	}

	header := HeaderStyle.Render("Schoty - AI Monitor")
	tabs := TabActiveStyle.Render("[API]") + TabInactiveStyle.Render(" [Subs]") + TabInactiveStyle.Render(" [Config]")
	content := BodyStyle.Render("Dashboard placeholder")

	return lipgloss.JoinVertical(
		lipgloss.Left,
		header,
		tabs,
		"",
		content,
		"",
		FooterStyle.Render("[q] quit  [r] refresh  [tab] switch tabs  [c] config"),
	)
}
