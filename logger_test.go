package logo

import (
	"bytes"
	"testing"
)

// TestNew checks if the logger is created correctly
func TestNew(t *testing.T) {
	// Create buffers
	rec1 := NewReceiver(&bytes.Buffer{}, "")
	rec2 := NewReceiver(&bytes.Buffer{}, "")
	rec3 := NewReceiver(&bytes.Buffer{}, "")
	// Create the logger
	l := NewLogger(rec1, rec2, rec3)
	if len(l.Receivers) != 3 {
		t.Error("Logger not created correctly")
	}
	if !l.Active {
		t.Error("Logger wasn't set to active")
	}
}

// TestNewSimple checks if a simple logger is created correctly
func TestNewSimple(t *testing.T) {
	// Create the logger
	l := NewSimpleLogger(&bytes.Buffer{}, WARN, "", true)
	r := l.Receivers[0]
	// Test stuff
	if r.Level != WARN {
		t.Error("Level wasn't assigned properly")
	}
	if !r.Color {
		t.Error("Color wasn't assigned properly")
	}
	if !l.Active {
		t.Error("Logger wasn't set to active")
	}
}

// TestOnlyIfActive checks if it's only logged if logger is active
func TestOnlyIfActive(t *testing.T) {
	buf := &bytes.Buffer{}
	rec := NewReceiver(buf, "")
	rec.Level = DEBUG
	l := NewLogger(rec)
	l.Active = false

	l.Debug("This shouldn't be logged")
	if buf.String() != "" {
		t.Error("Buffer should be empty")
	}
}

// TestSetLevel checks if the levels are set on all receivers
func TestSetLevel(t *testing.T) {
	buf := &bytes.Buffer{}

	rec1 := NewReceiver(buf, "")
	rec1.Level = DEBUG

	rec2 := NewReceiver(buf, "")
	rec2.Level = ERROR

	l := NewLogger(rec1, rec2)
	l.SetLevel(INFO)

	if rec1.Level != INFO || rec2.Level != INFO {
		t.Error("Level is not set correctly")
	}
}

// TestLogAll test if to all receivers is written
// TODO test Fatal level (os.Exit()-problem)
func TestLogLevels(t *testing.T) {
	buf := &bytes.Buffer{}
	rec := NewReceiver(buf, "")
	rec.Level = DEBUG
	l := NewLogger(rec)

	// Test Debug level
	l.Debug("Test")
	if buf.String() == "" {
		t.Errorf("Buffer shouldn't be empty")
	}
	buf.Reset()

	// Test Info level
	l.Info("Test")
	if buf.String() == "" {
		t.Errorf("Buffer shouldn't be empty")
	}
	buf.Reset()

	// Test Warn level
	l.Warn("Test")
	if buf.String() == "" {
		t.Errorf("Buffer shouldn't be empty")
	}
	buf.Reset()

	// Test Error level
	l.Error("Test")
	if buf.String() == "" {
		t.Errorf("Buffer shouldn't be empty")
	}
	buf.Reset()

	// Test Debug level with format
	l.Debugf("Test")
	if buf.String() == "" {
		t.Errorf("Buffer shouldn't be empty")
	}
	buf.Reset()

	// Test Info level with format
	l.Infof("Test")
	if buf.String() == "" {
		t.Errorf("Buffer shouldn't be empty")
	}
	buf.Reset()

	// Test Warn level with format
	l.Warnf("Test")
	if buf.String() == "" {
		t.Errorf("Buffer shouldn't be empty")
	}
	buf.Reset()

	// Test Error level with format
	l.Errorf("Test")
	if buf.String() == "" {
		t.Errorf("Buffer shouldn't be empty")
	}
}
