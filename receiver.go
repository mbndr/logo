package logo

import (
	"io"
	"log"
)

type Receiver struct {
	// Logger object from the builtin log package
	logger *log.Logger
	Level  level
	Color  bool
	Active bool
}

// Logs to the logger
func (r *Receiver) log(opt *levelOptions, s string) {
	// Don't do anything if not wanted
	if !r.Active || opt.Level < r.Level {
		return
	}

	// Add colors if wanted
	// TODO better here (no 2 prints!)
	if r.Color {
		r.logger.Printf("\x1b[0;%dm[%s] ▶ %s\x1b[0m\n", opt.Color, opt.Key, s)
		return
	}

	// Print to the logger
	r.logger.Printf("[%s] ▶ %s\n", opt.Key, s)
}

// NewReceiver returns a new Receiver object with a given Writer
// and sets default values
func NewReceiver(w io.Writer) *Receiver {
	r := &Receiver{
		logger: log.New(w, "", log.LstdFlags),
	}
	// Default options
	r.Active = true
	r.Level = INFO

	return r
}
