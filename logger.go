package logo

// TODO simple function to create a single receiver and return the logger

import (
    "os"
)

type level int

// Available log levels
const (
    DEBUG level = iota
    INFO
    WARN
    ERROR
    FATAL
)

// Logger holds all Receivers
type Logger struct {
    Receivers []*Receiver
}

// Write to all Receivers
func (l *Logger) logAll(lev level, key string, color int, a interface{}) {
    for _, r := range l.Receivers {
        r.log(lev, key, color, a)
    }
}

func (l *Logger) Debug(a interface{}) {
	l.logAll(DEBUG, "DEBU", 34, a)
}

func (l *Logger) Info(a interface{}) {
	l.logAll(INFO, "INFO", 32, a)
}

func (l *Logger) Warn(a interface{}) {
	l.logAll(WARN, "WARN", 33, a)
}

func (l *Logger) Error(a interface{}) {
	l.logAll(ERROR, "ERRO", 31, a)
}

func (l *Logger) Fatal(a interface{}) {
	l.logAll(FATAL, "CRIT", 31, a)
	os.Exit(1)
}

// Short function to open a file with needed options
func Open(path string) (*os.File, error) {
    return os.OpenFile("./example/logo.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0660)
}
