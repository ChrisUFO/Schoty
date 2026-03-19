package providers

type ClaudeCodeProvider struct{}

func (p *ClaudeCodeProvider) Name() string { return "Claude Code" }

func (p *ClaudeCodeProvider) CheckBalance() (*Balance, error) {
	return nil, nil
}

func (p *ClaudeCodeProvider) CheckUsage() (*Usage, error) {
	return nil, nil
}
