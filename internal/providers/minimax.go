package providers

type MiniMaxProvider struct{}

func (p *MiniMaxProvider) Name() string { return "MiniMax" }

func (p *MiniMaxProvider) CheckBalance() (*Balance, error) {
	return nil, nil
}

func (p *MiniMaxProvider) CheckUsage() (*Usage, error) {
	return nil, nil
}
