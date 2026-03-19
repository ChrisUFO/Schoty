package providers

type CodexProvider struct{}

func (p *CodexProvider) Name() string { return "Codex" }

func (p *CodexProvider) CheckBalance() (*Balance, error) {
	return nil, nil
}

func (p *CodexProvider) CheckUsage() (*Usage, error) {
	return nil, nil
}
