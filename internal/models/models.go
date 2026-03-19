package models

type AppState struct {
	RefreshInterval int
	LastRefresh     string
	ErrorCount      int
}

type ProviderInfo struct {
	Name         string
	Type         string
	Status       string
	Balance      float64
	Usage        int
	Remaining    int
	Limit        int
	IsLoading    bool
	IsConfigured bool
	ErrorMsg     string
}
