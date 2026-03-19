package ui

import (
	"context"
	"sync"
	"time"

	"github.com/ChrisUFO/Schoty/internal/config"
	"github.com/ChrisUFO/Schoty/internal/providers"
)

type ProviderResult struct {
	Name    string
	Balance *providers.Balance
	Usage   *providers.Usage
	Error   error
}

func CreateProvidersFromConfig(cfg *config.Config) []providers.Provider {
	var result []providers.Provider
	for _, pcfg := range cfg.Providers {
		if !pcfg.Enabled {
			continue
		}
		var p providers.Provider
		switch pcfg.Name {
		case "OpenAI":
			p = &providers.OpenAIProvider{}
		case "Anthropic":
			p = &providers.AnthropicProvider{}
		case "OpenRouter":
			p = &providers.OpenRouterProvider{}
		case "TogetherAI":
			p = &providers.TogetherAIProvider{}
		case "Claude Code":
			p = &providers.ClaudeCodeProvider{}
		case "Codex":
			p = &providers.CodexProvider{}
		case "Z.ai":
			p = &providers.ZAIProvider{}
		case "MiniMax":
			p = &providers.MiniMaxProvider{}
		default:
			continue
		}
		result = append(result, p)
	}
	return result
}

func FetchAllProviders(ctx context.Context, providerList []providers.Provider) []ProviderResult {
	var wg sync.WaitGroup
	results := make([]ProviderResult, len(providerList))

	for i, p := range providerList {
		wg.Add(1)
		go func(idx int, prov providers.Provider) {
			defer wg.Done()

			balance, balanceErr := prov.CheckBalance()
			usage, usageErr := prov.CheckUsage()

			result := ProviderResult{
				Name:    prov.Name(),
				Balance: balance,
				Usage:   usage,
				Error:   nil,
			}

			if balanceErr != nil {
				result.Error = balanceErr
			} else if usageErr != nil {
				result.Error = usageErr
			}

			results[idx] = result
		}(i, p)
	}

	wg.Wait()
	return results
}

func ProviderResultsToStates(results []ProviderResult) []ProviderState {
	states := make([]ProviderState, 0, len(results))
	for _, r := range results {
		state := ProviderState{
			Name:         r.Name,
			IsLoading:    false,
			IsConfigured: true,
		}

		if r.Error != nil {
			state.Status = "error"
			state.ErrorMsg = r.Error.Error()
		} else {
			if r.Balance != nil {
				state.Type = ProviderTypeBalance
				state.Balance = r.Balance.Amount
				if r.Balance.Amount > 0 {
					state.Status = "healthy"
				} else {
					state.Status = "critical"
				}
			}
			if r.Usage != nil {
				state.Type = ProviderTypeSubscription
				state.Usage = r.Usage.Used
				state.Remaining = r.Usage.Remaining
				state.Limit = r.Usage.Limit
				state.Status = CalculateStatus(state.Remaining, state.Limit)
			}
			if state.Type == "" {
				state.Type = ProviderTypeBalance
			}
		}

		states = append(states, state)
	}
	return states
}

func GetDefaultProviderStates() []ProviderState {
	providerNames := []string{"OpenAI", "Anthropic", "OpenRouter", "TogetherAI", "Claude Code", "Codex", "Z.ai", "MiniMax"}
	providerTypes := []string{ProviderTypeBalance, ProviderTypeBalance, ProviderTypeBalance, ProviderTypeBalance, ProviderTypeSubscription, ProviderTypeSubscription, ProviderTypeSubscription, ProviderTypeSubscription}
	states := make([]ProviderState, len(providerNames))
	for i, name := range providerNames {
		states[i] = ProviderState{
			Name:         name,
			Type:         providerTypes[i],
			Status:       "loading",
			IsLoading:    true,
			IsConfigured: false,
		}
	}
	return states
}

func CalculateStatus(remaining, limit int) string {
	if remaining <= 0 || limit <= 0 {
		return "critical"
	}
	percent := float64(remaining*100) / float64(limit)
	switch {
	case percent >= 50:
		return "healthy"
	case percent >= 20:
		return "warning"
	default:
		return "critical"
	}
}

func RefreshWithTimeout(ctx context.Context, providerList []providers.Provider, timeout time.Duration) []ProviderResult {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	resultChan := make(chan []ProviderResult, 1)

	go func() {
		resultChan <- FetchAllProviders(ctx, providerList)
	}()

	select {
	case results := <-resultChan:
		return results
	case <-ctx.Done():
		return []ProviderResult{}
	}
}
