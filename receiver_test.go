package logo

import (
	"bytes"
	"strings"
	"testing"
)

// TestActive tests if a receiver logs only if it's active
func TestActive(t *testing.T) {
	// Create buffer and receiver
	buf := &bytes.Buffer{}
	r := NewReceiver(buf, "")
	// Set inactive and log
	r.Active = false
	r.log(optError, "Just a test")
	if buf.String() != "" {
		t.Error("Buffer should be empty")
	}
	// Set active and log
	r.Active = true
	r.log(optError, "Just a test")
	if !strings.HasSuffix(buf.String(), "Just a test\n") {
		t.Error("Buffer should be filled")
	}
}

// TestDefaults checks if the default values are set correctly in the constructor
func TestDefaults(t *testing.T) {
	// Defaults: format, color, active
	r := NewReceiver(&bytes.Buffer{}, "")
	if r.Format != "[%s] ▶ %s" {
		t.Error("Default format is wrong")
	}
	if r.Color {
		t.Error("Color should be off")
	}
	if !r.Active {
		t.Error("Receiver should be activated")
	}
}

// TestFormat checks if a format is parsed correctly
func TestFormat(t *testing.T) {
	// Create buffer and receiver
	buf := &bytes.Buffer{}
	r := NewReceiver(buf, "")
	r.Format = "%s::%s"
	r.log(optInfo, "msg")
	if !strings.HasSuffix(buf.String(), "INFO::msg\n") {
		t.Error("Custom format should be active")
	}
}

// TestColor checks if the colors are printed only if wanted
func TestColor(t *testing.T) {
	// Create buffer and receiver
	buf := &bytes.Buffer{}
	r := NewReceiver(buf, "")
	// No color
	r.Color = false
	r.log(optInfo, "msg")
	if !strings.HasSuffix(buf.String(), "msg\n") {
		t.Error("Color shouldn't be active")
	}
	// Color
	r.Color = true
	r.log(optInfo, "msg")
	if !strings.HasSuffix(buf.String(), "\x1b[0;32m[INFO] ▶ msg\x1b[0m\n") {
		t.Error("Color should be active")
	}
}

// TestDebugLevels tests if a receiver logs only if the debug level is less or equal
func TestDebugLevels(t *testing.T) {
	// Create buffer and receiver
	buf := &bytes.Buffer{}
	r := NewReceiver(buf, "")

	// Set log level to debug
	r.Level = DEBUG
	r.log(optDebug, "Just a test")
	if buf.String() == "" {
		t.Errorf("Buffer shouldn't be empty")
	}
	buf.Reset()

	// Set log level to info
	r.Level = INFO
	r.log(optInfo, "Just a test")
	if buf.String() == "" {
		t.Errorf("Buffer shouldn't be empty")
	}
	buf.Reset()

	// Set log level to warn
	r.Level = WARN
	r.log(optWarn, "Just a test")
	if buf.String() == "" {
		t.Errorf("Buffer shouldn't be empty")
	}
	buf.Reset()

	// Set log level to error
	r.Level = ERROR
	r.log(optError, "Just a test")
	if buf.String() == "" {
		t.Errorf("Buffer shouldn't be empty")
	}
	buf.Reset()

	// Set log level to fatal
	r.Level = FATAL
	r.log(optFatal, "Just a test")
	if buf.String() == "" {
		t.Errorf("Buffer shouldn't be empty")
	}
	buf.Reset()

	// Set log level to debug, this should show an warning
	r.Level = DEBUG
	r.log(optWarn, "Just a test")
	if !strings.HasSuffix(buf.String(), "Just a test\n") {
		t.Error("Buffer should be filled")
	}
}
