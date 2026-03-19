package main

import (
	"fmt"
	"os"

	"github.com/ChrisUFO/Schoty/internal/ui"
	"github.com/charmbracelet/bubbletea"
)

func main() {
	model := ui.NewModel()
	p := tea.NewProgram(&model)

	if err := p.Start(); err != nil {
		fmt.Fprintf(os.Stderr, "Error starting Schoty: %v\n", err)
		os.Exit(1)
	}
}
