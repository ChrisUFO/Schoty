package main

import (
	"os"

	"github.com/ChrisUFO/Schoty/internal/ui"
	"github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(ui.NewModel())
	if _, err := p.Run(); err != nil {
		os.Exit(1)
	}
}
