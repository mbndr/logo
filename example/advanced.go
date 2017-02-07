package main

import (
    "github.com/probandula/logo"
    "os"
)

func main() {
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
    fileRec.Level = logo.DEBUG

    // Create the logger
    log := logo.NewLogger(cliRec, fileRec)

    // Test some messages
    log.Debug("Debug info", "and other info", 230)
}
