package config

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoadConfig_FileNotFound(t *testing.T) {
	viper.Reset()

	cfg, err := LoadConfig()
	require.NoError(t, err)
	assert.NotNil(t, cfg)
	assert.Empty(t, cfg.Providers)
}

func TestLoadConfig_WithFile(t *testing.T) {
	viper.Reset()

	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "config.yaml")
	err := os.WriteFile(configPath, []byte(`
providers:
  - name: OpenAI
    api_key: test-key-123
    enabled: true
  - name: Anthropic
    api_key: ""
    enabled: false
`), 0644)
	require.NoError(t, err)

	viper.AddConfigPath(tmpDir)
	cfg, err := LoadConfig()
	require.NoError(t, err)
	require.Len(t, cfg.Providers, 2)
	assert.Equal(t, "OpenAI", cfg.Providers[0].Name)
	assert.Equal(t, "test-key-123", cfg.Providers[0].APIKey)
	assert.True(t, cfg.Providers[0].Enabled)
	assert.Equal(t, "Anthropic", cfg.Providers[1].Name)
	assert.False(t, cfg.Providers[1].Enabled)
}

func TestGetProviderConfig(t *testing.T) {
	viper.Reset()

	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "config.yaml")
	err := os.WriteFile(configPath, []byte(`
providers:
  - name: OpenAI
    api_key: key-123
    enabled: true
`), 0644)
	require.NoError(t, err)

	viper.AddConfigPath(tmpDir)
	_, err = LoadConfig()
	require.NoError(t, err)

	provider := GetProviderConfig("OpenAI")
	require.NotNil(t, provider)
	assert.Equal(t, "OpenAI", provider.Name)
	assert.Equal(t, "key-123", provider.APIKey)

	notFound := GetProviderConfig("NonExistent")
	assert.Nil(t, notFound)
}

func TestGetEnabledProviders(t *testing.T) {
	viper.Reset()

	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "config.yaml")
	err := os.WriteFile(configPath, []byte(`
providers:
  - name: OpenAI
    enabled: true
  - name: Anthropic
    enabled: false
  - name: TogetherAI
    enabled: true
`), 0644)
	require.NoError(t, err)

	viper.AddConfigPath(tmpDir)
	_, err = LoadConfig()
	require.NoError(t, err)

	enabled := GetEnabledProviders()
	require.Len(t, enabled, 2)
	assert.Equal(t, "OpenAI", enabled[0].Name)
	assert.Equal(t, "TogetherAI", enabled[1].Name)
}

func TestEnvOverride(t *testing.T) {
	viper.Reset()

	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "config.yaml")
	err := os.WriteFile(configPath, []byte(`
providers:
  - name: OpenAI
    api_key: original-key
    enabled: true
`), 0644)
	require.NoError(t, err)

	viper.AddConfigPath(tmpDir)
	os.Setenv("SCHOTY_OPENAI_API_KEY", "env-key-456")
	defer os.Unsetenv("SCHOTY_OPENAI_API_KEY")

	_, err = LoadConfig()
	require.NoError(t, err)

	provider := GetProviderConfig("OpenAI")
	require.NotNil(t, provider)
	assert.Equal(t, "env-key-456", provider.APIKey)
}

func TestEnvOverrideWithDashesAndSpaces(t *testing.T) {
	tests := []struct {
		name         string
		providerName string
		envKey       string
		envValue     string
	}{
		{
			name:         "Z.ai provider with dot",
			providerName: "Z.ai",
			envKey:       "SCHOTY_Z_AI_API_KEY",
			envValue:     "z-ai-env-key",
		},
		{
			name:         "Claude Code provider with space",
			providerName: "Claude Code",
			envKey:       "SCHOTY_CLAUDE_CODE_API_KEY",
			envValue:     "claude-code-env-key",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			viper.Reset()

			tmpDir := t.TempDir()
			configPath := filepath.Join(tmpDir, "config.yaml")
			err := os.WriteFile(configPath, []byte(`
providers:
  - name: "`+tc.providerName+`"
    api_key: original-key
    enabled: true
`), 0644)
			require.NoError(t, err)

			viper.AddConfigPath(tmpDir)
			os.Setenv(tc.envKey, tc.envValue)
			defer os.Unsetenv(tc.envKey)

			_, err = LoadConfig()
			require.NoError(t, err)

			provider := GetProviderConfig(tc.providerName)
			require.NotNil(t, provider, "provider %s should be found", tc.providerName)
			assert.Equal(t, tc.envValue, provider.APIKey)
		})
	}
}
