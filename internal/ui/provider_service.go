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

type providerMeta struct {
	Name    string
	Type    string
	Factory func() providers.Provider
}

var allProviderMetas = []providerMeta{
	{Name: "OpenAI", Type: ProviderTypeBalance, Factory: func() providers.Provider { return &providers.OpenAIProvider{} }},
	{Name: "Anthropic", Type: ProviderTypeBalance, Factory: func() providers.Provider { return &providers.AnthropicProvider{} }},
	{Name: "OpenRouter", Type: ProviderTypeBalance, Factory: func() providers.Provider { return &providers.OpenRouterProvider{} }},
	{Name: "TogetherAI", Type: ProviderTypeBalance, Factory: func() providers.Provider { return &providers.TogetherAIProvider{} }},
	{Name: "Claude Code", Type: ProviderTypeSubscription, Factory: func() providers.Provider { return &providers.ClaudeCodeProvider{} }},
	{Name: "Codex", Type: ProviderTypeSubscription, Factory: func() providers.Provider { return &providers.CodexProvider{} }},
	{Name: "Z.ai", Type: ProviderTypeSubscription, Factory: func() providers.Provider { return &providers.ZAIProvider{} }},
	{Name: "MiniMax", Type: ProviderTypeSubscription, Factory: func() providers.Provider { return &providers.MiniMaxProvider{} }},
}

func GetAllProviderNames() []string {
	names := make([]string, len(allProviderMetas))
	for i, m := range allProviderMetas {
		names[i] = m.Name
	}
	return names
}

func GetProviderMetaByName(name string) *providerMeta {
	for i := range allProviderMetas {
		if allProviderMetas[i].Name == name {
			return &allProviderMetas[i]
		}
	}
	return nil
}

func CreateProvidersFromConfig(cfg *config.Config) []providers.Provider {
	var result []providers.Provider
	for _, pcfg := range cfg.Providers {
		if !pcfg.Enabled {
			continue
		}
		meta := GetProviderMetaByName(pcfg.Name)
		if meta == nil {
			continue
		}
		result = append(result, meta.Factory())
	}
	return result
}

func FetchAllProviders(ctx context.Context, providerList []providers.Provider) []ProviderResult {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	resultChan := make(chan []ProviderResult, 1)

	go func() {
		resultChan <- fetchProvidersInternal(providerList)
	}()

	select {
	case results := <-resultChan:
		return results
	case <-ctx.Done():
		timeoutResults := make([]ProviderResult, len(providerList))
		for i, p := range providerList {
			timeoutResults[i] = ProviderResult{
				Name:  p.Name(),
				Error: ctx.Err(),
			}
		}
		return timeoutResults
	}
}

func fetchProvidersInternal(providerList []providers.Provider) []ProviderResult {
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
				state.Type = "unknown"
			}
		}

		states = append(states, state)
	}
	return states
}

func GetDefaultProviderStates() []ProviderState {
	states := make([]ProviderState, len(allProviderMetas))
	for i, meta := range allProviderMetas {
		states[i] = ProviderState{
			Name:         meta.Name,
			Type:         meta.Type,
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
