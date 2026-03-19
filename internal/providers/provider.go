package providers

type Balance struct {
	Amount   float64
	Currency string
}

type Usage struct {
	Used      int
	Remaining int
	Limit     int
}

type Provider interface {
	Name() string
	CheckBalance() (*Balance, error)
	CheckUsage() (*Usage, error)
}
