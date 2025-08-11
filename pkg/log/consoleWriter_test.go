package log

import (
	"bytes"
	"encoding/json"
	"errors"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/rs/zerolog"
)

func TestConsoleWriter_Write_BasicFields(t *testing.T) {
	var buf bytes.Buffer
	writer := ConsoleWriter{
		Out:        &buf,
		NoColor:    true,
		TimeFormat: time.Kitchen,
	}

	event := map[string]interface{}{
		zerolog.TimestampFieldName: time.Now().Format(zerolog.TimeFieldFormat),
		zerolog.LevelFieldName:     zerolog.LevelInfoValue,
		zerolog.MessageFieldName:   "hello world",
		zerolog.CallerFieldName:    "/path/to/file.go",
	}

	b, err := json.Marshal(event)
	if err != nil {
		t.Fatalf("failed to marshal event: %v", err)
	}

	n, err := writer.Write(b)
	if err != nil {
		t.Fatalf("Write failed: %v", err)
	}
	if n != len(b) {
		t.Errorf("expected n=%d, got %d", len(b), n)
	}

	out := buf.String()
	if !strings.Contains(out, "hello world") {
		t.Errorf("output missing message: %s", out)
	}
	if !strings.Contains(out, "INF") {
		t.Errorf("output missing level: %s", out)
	}
}

func TestConsoleWriter_Write_FieldsOrderAndExclude(t *testing.T) {
	var buf bytes.Buffer
	writer := ConsoleWriter{
		Out:          &buf,
		NoColor:      true,
		PartsOrder:   []string{zerolog.LevelFieldName, zerolog.MessageFieldName},
		PartsExclude: []string{zerolog.LevelFieldName},
	}

	event := map[string]interface{}{
		zerolog.LevelFieldName:   zerolog.LevelDebugValue,
		zerolog.MessageFieldName: "test message",
	}

	b, _ := json.Marshal(event)
	_, err := writer.Write(b)
	if err != nil {
		t.Fatalf("Write failed: %v", err)
	}

	out := buf.String()
	if strings.Contains(out, "DBG") {
		t.Errorf("excluded part still present: %s", out)
	}
	if !strings.Contains(out, "test message") {
		t.Errorf("message missing: %s", out)
	}
}

func TestConsoleWriter_Write_FieldsExclude(t *testing.T) {
	var buf bytes.Buffer
	writer := ConsoleWriter{
		Out:           &buf,
		NoColor:       true,
		FieldsExclude: []string{"foo"},
	}

	event := map[string]interface{}{
		zerolog.MessageFieldName: "msg",
		"foo":                    "bar",
		"baz":                    "qux",
	}

	b, _ := json.Marshal(event)
	_, err := writer.Write(b)
	if err != nil {
		t.Fatalf("Write failed: %v", err)
	}

	out := buf.String()
	if strings.Contains(out, "foo=") {
		t.Errorf("excluded field still present: %s", out)
	}
	if !strings.Contains(out, "baz=") {
		t.Errorf("baz field missing: %s", out)
	}
}

func TestConsoleWriter_Write_ErrorField(t *testing.T) {
	var buf bytes.Buffer
	writer := ConsoleWriter{
		Out:     &buf,
		NoColor: true,
	}

	event := map[string]interface{}{
		zerolog.MessageFieldName: "msg",
		zerolog.ErrorFieldName:   "something went wrong",
	}

	b, _ := json.Marshal(event)
	_, err := writer.Write(b)
	if err != nil {
		t.Fatalf("Write failed: %v", err)
	}

	out := buf.String()
	if !strings.Contains(out, "error=") {
		t.Errorf("error field missing: %s", out)
	}
	if !strings.Contains(out, "something went wrong") {
		t.Errorf("error value missing: %s", out)
	}
}

func TestConsoleWriter_Write_FormatExtra(t *testing.T) {
	var buf bytes.Buffer
	writer := ConsoleWriter{
		Out:     &buf,
		NoColor: true,
		FormatExtra: func(evt map[string]interface{}, buf *bytes.Buffer) error {
			buf.WriteString(" EXTRA")
			return nil
		},
	}

	event := map[string]interface{}{
		zerolog.MessageFieldName: "msg",
	}

	b, _ := json.Marshal(event)
	_, err := writer.Write(b)
	if err != nil {
		t.Fatalf("Write failed: %v", err)
	}

	out := buf.String()
	if !strings.Contains(out, "EXTRA") {
		t.Errorf("FormatExtra output missing: %s", out)
	}
}

func TestConsoleWriter_Write_FormatExtraError(t *testing.T) {
	var buf bytes.Buffer
	writer := ConsoleWriter{
		Out:     &buf,
		NoColor: true,
		FormatExtra: func(evt map[string]interface{}, buf *bytes.Buffer) error {
			return errors.New("extra error")
		},
	}

	event := map[string]interface{}{
		zerolog.MessageFieldName: "msg",
	}

	b, _ := json.Marshal(event)
	_, err := writer.Write(b)
	if err == nil || !strings.Contains(err.Error(), "extra error") {
		t.Errorf("expected error from FormatExtra, got: %v", err)
	}
}

func TestConsoleWriter_Write_InvalidJSON(t *testing.T) {
	var buf bytes.Buffer
	writer := ConsoleWriter{
		Out:     &buf,
		NoColor: true,
	}

	invalid := []byte("{invalid json}")
	_, err := writer.Write(invalid)
	if err == nil || !strings.Contains(err.Error(), "cannot decode event") {
		t.Errorf("expected decode error, got: %v", err)
	}
}

func TestConsoleWriter_Write_OutIsStdoutOrStderr(t *testing.T) {
	// This test just ensures no panic when Out is os.Stdout or os.Stderr
	writer := ConsoleWriter{
		Out:     os.Stdout,
		NoColor: true,
	}

	event := map[string]interface{}{
		zerolog.MessageFieldName: "msg",
	}

	b, _ := json.Marshal(event)
	_, err := writer.Write(b)
	if err != nil {
		t.Errorf("Write failed: %v", err)
	}
}

func TestConsoleWriter_Write_QuotedFieldValue(t *testing.T) {
	var buf bytes.Buffer
	writer := ConsoleWriter{
		Out:     &buf,
		NoColor: true,
	}

	event := map[string]interface{}{
		zerolog.MessageFieldName: "msg",
		"foo":                    "bar baz",
	}

	b, _ := json.Marshal(event)
	_, err := writer.Write(b)
	if err != nil {
		t.Fatalf("Write failed: %v", err)
	}

	out := buf.String()
	if !strings.Contains(out, `"bar baz"`) {
		t.Errorf("quoted field value missing: %s", out)
	}
}
