package logging

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	Init("debug")
	assert.NotNil(t, defaultLogger)
}

func TestInitWithWriter(t *testing.T) {
	var buf bytes.Buffer
	InitWithWriter(&buf, "info")
	assert.NotNil(t, defaultLogger)

	Info("test message", "key", "value")
	assert.Contains(t, buf.String(), "test message")
	assert.Contains(t, buf.String(), "key")
	assert.Contains(t, buf.String(), "value")
}

func TestLogLevels(t *testing.T) {
	var buf bytes.Buffer
	InitWithWriter(&buf, "error")

	Debug("debug message", "level", "debug")
	Info("info message", "level", "info")
	Warn("warn message", "level", "warn")
	Error("error message", "level", "error")

	output := buf.String()
	assert.Contains(t, output, "error message")
	assert.NotContains(t, output, "debug message")
	assert.NotContains(t, output, "info message")
	assert.NotContains(t, output, "warn message")
}

func TestWith(t *testing.T) {
	var buf bytes.Buffer
	InitWithWriter(&buf, "info")

	logger := With("component", "test")
	logger.Info("test message")

	assert.Contains(t, buf.String(), "component")
	assert.Contains(t, buf.String(), "test")
	assert.Contains(t, buf.String(), "test message")
}

func TestInvalidLogLevel(t *testing.T) {
	var buf bytes.Buffer
	InitWithWriter(&buf, "invalid")
	assert.NotNil(t, defaultLogger)

	Info("test after invalid level")
	assert.Contains(t, buf.String(), "test after invalid level")
}

func TestInfoLogging(t *testing.T) {
	var buf bytes.Buffer
	InitWithWriter(&buf, "info")

	Info("starting", "action", "test")

	output := buf.String()
	assert.True(t, strings.Contains(output, "starting") || strings.Contains(output, "action"))
}
