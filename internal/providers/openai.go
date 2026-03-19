package providers

type OpenAIProvider struct{}

func (p *OpenAIProvider) Name() string { return "OpenAI" }

func (p *OpenAIProvider) CheckBalance() (*Balance, error) {
	return nil, nil
}

func (p *OpenAIProvider) CheckUsage() (*Usage, error) {
	return nil, nil
}
