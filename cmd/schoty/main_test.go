package main

import (
	"testing"

	"github.com/ChrisUFO/Schoty/internal/ui"
)

func TestNewModel(t *testing.T) {
	model := ui.NewModel()
	if model.Ready {
		t.Error("Expected Ready to be false initially")
	}
	if model.Tab != 0 {
		t.Errorf("Expected Tab to be 0, got %d", model.Tab)
	}
	if model.CurrentView != ui.DashboardView {
		t.Errorf("Expected CurrentView to be DashboardView")
	}
	if len(model.Providers) != 0 {
		t.Errorf("Expected empty Providers, got %d", len(model.Providers))
	}
}

func TestModelViewStates(t *testing.T) {
	if ui.DashboardView != 0 {
		t.Error("DashboardView should be 0")
	}
	if ui.DetailView != 1 {
		t.Error("DetailView should be 1")
	}
	if ui.ConfigView != 2 {
		t.Error("ConfigView should be 2")
	}
	if ui.HelpView != 3 {
		t.Error("HelpView should be 3")
	}
}
