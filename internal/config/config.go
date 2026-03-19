package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Providers []ProviderConfig `mapstructure:"providers"`
}

type ProviderConfig struct {
	Name    string `mapstructure:"name"`
	APIKey  string `mapstructure:"api_key"`
	Enabled bool   `mapstructure:"enabled"`
	BaseURL string `mapstructure:"base_url"`
	Model   string `mapstructure:"model"`
}

var AppConfig *Config

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME/.schoty")

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return &Config{Providers: []ProviderConfig{}}, nil
		}
		return nil, fmt.Errorf("failed to read config: %w", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	applyEnvOverrides(&cfg)
	AppConfig = &cfg
	return &cfg, nil
}

func applyEnvOverrides(cfg *Config) {
	for i := range cfg.Providers {
		envKey := fmt.Sprintf("SCHOTY_%s_API_KEY", normalizeEnvName(cfg.Providers[i].Name))
		if apiKey := os.Getenv(envKey); apiKey != "" {
			cfg.Providers[i].APIKey = apiKey
		}
	}
}

func normalizeEnvName(name string) string {
	return strings.ToUpper(strings.ReplaceAll(name, "-", "_"))
}

func GetProviderConfig(name string) *ProviderConfig {
	if AppConfig == nil {
		return nil
	}
	for i := range AppConfig.Providers {
		if AppConfig.Providers[i].Name == name {
			return &AppConfig.Providers[i]
		}
	}
	return nil
}

func GetEnabledProviders() []ProviderConfig {
	if AppConfig == nil {
		return []ProviderConfig{}
	}
	var enabled []ProviderConfig
	for _, p := range AppConfig.Providers {
		if p.Enabled {
			enabled = append(enabled, p)
		}
	}
	return enabled
}
