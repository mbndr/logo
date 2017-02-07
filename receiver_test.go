package logo

import (
	"bytes"
	"strings"
	"testing"
)

// TODO color test

// TestActive tests if a receiver logs only if it's active
func TestActive(t *testing.T) {
	// Create buffer and receiver
	buf := bytes.NewBufferString("")
	r := NewReceiver(buf)
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

// TestDebugLevels tests if a receiver logs only if the debug level is less or equal
func TestDebugLevels(t *testing.T) {
	// Create buffer and receiver
	buf := bytes.NewBufferString("")
	r := NewReceiver(buf)
	// Set log level to debug
	r.Level = DEBUG
	r.log(optDebug, "Just a test")
	if buf.String() == "" {
		t.Errorf("Buffer shouldn't be empty")
	}
	// Set log level to info
	r.Level = INFO
	r.log(optInfo, "Just a test")
	if buf.String() == "" {
		t.Errorf("Buffer shouldn't be empty")
	}
	// Set log level to warn
	r.Level = WARN
	r.log(optWarn, "Just a test")
	if buf.String() == "" {
		t.Errorf("Buffer shouldn't be empty")
	}
	// Set log level to error
	r.Level = ERROR
	r.log(optError, "Just a test")
	if buf.String() == "" {
		t.Errorf("Buffer shouldn't be empty")
	}
	// Set log level to fatal
	r.Level = FATAL
	r.log(optFatal, "Just a test")
	if buf.String() == "" {
		t.Errorf("Buffer shouldn't be empty")
	}
	// Set log level to debug, this should show an warning
	r.Level = DEBUG
	r.log(optWarn, "Just a test")
	if !strings.HasSuffix(buf.String(), "Just a test\n") {
		t.Error("Buffer should be filled")
	}
}
