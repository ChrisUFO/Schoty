package ui

import (
	"testing"
	"time"

	"github.com/ChrisUFO/Schoty/internal/providers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewModel(t *testing.T) {
	model := NewModel()

	assert.False(t, model.Ready)
	assert.Equal(t, 80, model.Width)
	assert.Equal(t, 24, model.Height)
	assert.Equal(t, 0, model.Tab)
	assert.Equal(t, DashboardView, model.CurrentView)
	assert.Equal(t, 0, model.SelectedProvider)
	assert.NotNil(t, model.Providers)
	assert.Empty(t, model.Providers)
	assert.False(t, model.ShowHelp)
	assert.Equal(t, DefaultRefreshInterval, model.RefreshInterval)
	assert.Nil(t, model.ticker)
}

func TestProviderResultsToStates(t *testing.T) {
	results := []ProviderResult{
		{
			Name:    "OpenAI",
			Balance: &providers.Balance{Amount: 25.50},
			Usage:   nil,
			Error:   nil,
		},
		{
			Name:    "Claude Code",
			Balance: nil,
			Usage: &providers.Usage{
				Used:      500,
				Remaining: 4500,
				Limit:     5000,
			},
			Error: nil,
		},
	}

	states := ProviderResultsToStates(results)

	require.Len(t, states, 2)

	assert.Equal(t, "OpenAI", states[0].Name)
	assert.Equal(t, ProviderTypeBalance, states[0].Type)
	assert.Equal(t, 25.50, states[0].Balance)
	assert.Equal(t, "healthy", states[0].Status)
	assert.False(t, states[0].IsLoading)
	assert.True(t, states[0].IsConfigured)

	assert.Equal(t, "Claude Code", states[1].Name)
	assert.Equal(t, ProviderTypeSubscription, states[1].Type)
	assert.Equal(t, 500, states[1].Usage)
	assert.Equal(t, 4500, states[1].Remaining)
	assert.Equal(t, 5000, states[1].Limit)
	assert.Equal(t, "healthy", states[1].Status)
}

func TestProviderResultsToStatesWithError(t *testing.T) {
	results := []ProviderResult{
		{
			Name:    "OpenAI",
			Balance: nil,
			Usage:   nil,
			Error:   assert.AnError,
		},
	}

	states := ProviderResultsToStates(results)

	require.Len(t, states, 1)
	assert.Equal(t, "error", states[0].Status)
	assert.Equal(t, assert.AnError.Error(), states[0].ErrorMsg)
}

func TestGetDefaultProviderStates(t *testing.T) {
	states := GetDefaultProviderStates()

	require.Len(t, states, 8)

	assert.Equal(t, "OpenAI", states[0].Name)
	assert.Equal(t, ProviderTypeBalance, states[0].Type)
	assert.True(t, states[0].IsLoading)
	assert.False(t, states[0].IsConfigured)

	assert.Equal(t, "Claude Code", states[4].Name)
	assert.Equal(t, ProviderTypeSubscription, states[4].Type)
	assert.True(t, states[4].IsLoading)
}

func TestCalculateStatus(t *testing.T) {
	tests := []struct {
		name      string
		remaining int
		limit     int
		expected  string
	}{
		{"healthy high", 80, 100, "healthy"},
		{"healthy exactly 50", 50, 100, "healthy"},
		{"warning low", 30, 100, "warning"},
		{"warning exactly 20", 20, 100, "warning"},
		{"critical very low", 10, 100, "critical"},
		{"critical zero remaining", 0, 100, "critical"},
		{"critical zero limit", 50, 0, "critical"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CalculateStatus(tt.remaining, tt.limit)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestModelRefreshInterval(t *testing.T) {
	model := NewModel()
	assert.Equal(t, DefaultRefreshInterval, model.RefreshInterval)

	model.RefreshInterval = 30 * time.Second
	assert.Equal(t, 30*time.Second, model.RefreshInterval)
}

func TestTickMsg(t *testing.T) {
	now := time.Now()
	msg := tickMsg(now)
	assert.Equal(t, now, time.Time(msg))
}
