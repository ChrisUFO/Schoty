package models

type AppState struct {
	RefreshInterval int
	LastRefresh     string
	ErrorCount      int
}
