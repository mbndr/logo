package main

import (
	"github.com/probandula/logo"
	"os"
)

func advanced() {
	// Open log file (io.Writer)
	logFile, err := logo.Open("./example/logo.log")
	if err != nil {
		panic(err)
	}
	defer logFile.Close()

	// Receiver for the terminal
	cliRec := logo.NewReceiver(os.Stderr)
	cliRec.Color = true
	cliRec.Level = logo.DEBUG

	// Receiver for the log file
	fileRec := logo.NewReceiver(logFile)
	fileRec.Format = "%s: %s"

	// Create the logger
	log := logo.NewLogger(cliRec, fileRec)

	// Test some messages
	log.Info("Information")
}
