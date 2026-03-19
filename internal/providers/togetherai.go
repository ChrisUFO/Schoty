package providers

type TogetherAIProvider struct{}

func (p *TogetherAIProvider) Name() string { return "Together.ai" }

func (p *TogetherAIProvider) CheckBalance() (*Balance, error) {
	return nil, nil
}

func (p *TogetherAIProvider) CheckUsage() (*Usage, error) {
	return nil, nil
}
