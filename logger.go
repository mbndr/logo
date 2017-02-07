package logo

import (
	"fmt"
	"io"
	"os"
)

// Logger holds all Receivers
type Logger struct {
	Receivers []*Receiver
}

// NewLogger returns a new Logger filled with given Receivers
func NewLogger(recs ...*Receiver) *Logger {
	l := &Logger{
		Receivers: recs,
	}
	return l
}

// SimpleLogger returns a logger with one simple Receiver
func NewSimpleLogger(w io.Writer, lvl level, color bool) *Logger {
	l := &Logger{}
	r := NewReceiver(w)
	r.Color = color
    r.Level = lvl
	l.Receivers = []*Receiver{r}
	return l
}

// Write to all Receivers
func (l *Logger) logAll(opt *levelOptions, s string) {
	for _, r := range l.Receivers {
		r.log(opt, s)
	}
}

// Debug logs arguments
func (l *Logger) Debug(a ...interface{}) {
	l.logAll(optDebug, fmt.Sprint(a...))
}

// Info logs arguments
func (l *Logger) Info(a ...interface{}) {
	l.logAll(optInfo, fmt.Sprint(a...))
}

// Warn logs arguments
func (l *Logger) Warn(a ...interface{}) {
	l.logAll(optWarn, fmt.Sprint(a...))
}

// Error logs arguments
func (l *Logger) Error(a ...interface{}) {
	l.logAll(optError, fmt.Sprint(a...))
}

// Fatal logs arguments
func (l *Logger) Fatal(a ...interface{}) {
	l.logAll(optFatal, fmt.Sprint(a...))
	os.Exit(1)
}

// Debugf logs formated arguments
func (l *Logger) Debugf(format string, a ...interface{}) {
	l.logAll(optDebug, fmt.Sprintf(format, a...))
}

// Infof logs formated arguments
func (l *Logger) Infof(format string, a ...interface{}) {
	l.logAll(optInfo, fmt.Sprintf(format, a...))
}

// Warnf logs formated arguments
func (l *Logger) Warnf(format string, a ...interface{}) {
	l.logAll(optWarn, fmt.Sprintf(format, a...))
}

// Errorf logs formated arguments
func (l *Logger) Errorf(format string, a ...interface{}) {
	l.logAll(optFatal, fmt.Sprintf(format, a...))
}

// Fatalf logs formated arguments
func (l *Logger) Fatalf(format string, a ...interface{}) {
	l.logAll(optFatal, fmt.Sprintf(format, a...))
}

// Short function to open a file with needed options
func Open(path string) (*os.File, error) {
	return os.OpenFile("./example/logo.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0660)
}
