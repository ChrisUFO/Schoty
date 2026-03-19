package providers

type ZAIProvider struct{}

func (p *ZAIProvider) Name() string { return "Z.ai" }

func (p *ZAIProvider) CheckBalance() (*Balance, error) {
	return nil, nil
}

func (p *ZAIProvider) CheckUsage() (*Usage, error) {
	return nil, nil
}
