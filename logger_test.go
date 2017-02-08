package logo

import (
	"bytes"
	"testing"
)

// TestNew checks if the logger is created correctly
func TestNew(t *testing.T) {
	// Create buffers
	rec1 := NewReceiver(bytes.NewBufferString(""))
    rec2 := NewReceiver(bytes.NewBufferString(""))
    rec3 := NewReceiver(bytes.NewBufferString(""))
    // Create the logger
    l := NewLogger(rec1, rec2, rec3)
    if len(l.Receivers) != 3 {
        t.Error("Logger not created correctly")
    }
}

// TestNewSimple checks if a simple logger is created correctly
func TestNewSimple(t *testing.T) {
    // Create the logger
    l := NewSimpleLogger(bytes.NewBufferString(""), WARN, true)
    r := l.Receivers[0]
    // Test stuff
    if r.Level != WARN {
        t.Error("Level wasn't assigned properly")
    }
    if !r.Color {
        t.Error("Color wasn't assigned properly")
    }
}
