package providers

type OpenRouterProvider struct{}

func (p *OpenRouterProvider) Name() string { return "OpenRouter" }

func (p *OpenRouterProvider) CheckBalance() (*Balance, error) {
	return nil, nil
}

func (p *OpenRouterProvider) CheckUsage() (*Usage, error) {
	return nil, nil
}
