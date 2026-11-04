package log

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestGetLogWriter_PrettyPrint_DefaultSettings(t *testing.T) {
	os.Setenv("KINDE_STRUCTURED_LOG", "false")
	os.Unsetenv("KINDE_LOG_LEVEL")

	settings := &SharedLogSettings{
		ComponentName: "testComponent",
	}

	var buf bytes.Buffer
	origStderr := os.Stderr
	r, w, _ := os.Pipe()
	os.Stderr = w

	logger := GetLogWriter(settings)
	logger.Warn().Msg("test message")

	w.Close()
	os.Stderr = origStderr
	buf.ReadFrom(r)
	output := buf.String()

	if !strings.Contains(output, "test message") {
		t.Errorf("Expected log output to contain 'test message', got: %s", output)
	}
	if !strings.Contains(output, "testComponent") {
		t.Errorf("Expected log output to contain component name, got: %s", output)
	}
}

func TestGetLogWriter_StructuredLog(t *testing.T) {
	os.Setenv("KINDE_STRUCTURED_LOG", "true")
	os.Unsetenv("KINDE_LOG_LEVEL")

	settings := &SharedLogSettings{
		ComponentName: "structuredComponent",
	}

	var buf bytes.Buffer
	origStderr := os.Stderr
	r, w, _ := os.Pipe()
	os.Stderr = w

	logger := GetLogWriter(settings)
	logger.Warn().Msg("structured log message")

	w.Close()
	os.Stderr = origStderr
	buf.ReadFrom(r)
	output := buf.String()

	if !strings.Contains(output, "structured log message") {
		t.Errorf("Expected log output to contain 'structured log message', got: %s", output)
	}
	if !strings.Contains(output, "structuredComponent") {
		t.Errorf("Expected log output to contain component name, got: %s", output)
	}
	// Structured logs should be JSON
	if !strings.Contains(output, "{") {
		t.Errorf("Expected structured log output to be JSON, got: %s", output)
	}
}

func TestGetLogWriter_InvalidLogLevel(t *testing.T) {
	os.Setenv("KINDE_STRUCTURED_LOG", "false")
	os.Setenv("KINDE_LOG_LEVEL", "invalid_level")

	settings := &SharedLogSettings{
		ComponentName: "invalidLevelComponent",
	}

	var buf bytes.Buffer
	origStderr := os.Stderr
	r, w, _ := os.Pipe()
	os.Stderr = w

	logger := GetLogWriter(settings)
	logger.Warn().Msg("info message")

	w.Close()
	os.Stderr = origStderr
	buf.ReadFrom(r)
	output := buf.String()

	if !strings.Contains(output, "info message") {
		t.Errorf("Expected log output to contain 'info message', got: %s", output)
	}
}

func TestGetLogWriter_HidesInfoByDefault(t *testing.T) {
	t.Setenv("KINDE_STRUCTURED_LOG", "false")
	os.Unsetenv("KINDE_LOG_LEVEL")

	var buf bytes.Buffer
	logger := GetLogWriter(&SharedLogSettings{ComponentName: "quiet", Out: &buf})
	logger.Info().Msg("routine detail")
	logger.Warn().Msg("something is off")

	output := buf.String()
	if strings.Contains(output, "routine detail") {
		t.Errorf("Expected info logs to be hidden by default, got: %s", output)
	}
	if !strings.Contains(output, "something is off") {
		t.Errorf("Expected warnings to be shown, got: %s", output)
	}
}
