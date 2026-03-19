package providers

type AnthropicProvider struct{}

func (p *AnthropicProvider) Name() string { return "Anthropic" }

func (p *AnthropicProvider) CheckBalance() (*Balance, error) {
	return nil, nil
}

func (p *AnthropicProvider) CheckUsage() (*Usage, error) {
	return nil, nil
}
