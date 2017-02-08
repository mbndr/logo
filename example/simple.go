package main

import (
	"github.com/probandula/logo"
	"os"
)

func simple() {
	// Create a simple cli logger with activated colors
	log := logo.NewSimpleLogger(os.Stderr, logo.DEBUG, true)

	// Test some messages
	log.Debug("Debug info")
}
