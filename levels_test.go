package logo

import (
	"testing"
)

// TestItol checks the correct conversion from an int to a Level
func TestItol(t *testing.T) {
	if Itol(0) != DEBUG {
		t.Error("Conversion failure")
	}
	if Itol(1) != INFO {
		t.Error("Conversion failure")
	}
	if Itol(2) != WARN {
		t.Error("Conversion failure")
	}
	if Itol(3) != ERROR {
		t.Error("Conversion failure")
	}
	if Itol(4) != FATAL {
		t.Error("Conversion failure")
	}
	// Default value is DEBUG
	if Itol(100) != DEBUG {
		t.Error("Conversion failure")
	}
}
