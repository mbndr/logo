package logo

import (
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
)

// Receiver holds all receiver options
type Receiver struct {
	// Logger object from the builtin log package
	logger *log.Logger
	Level  Level
	Color  bool
	Active bool
	Format string
}

// SetPrefix sets the prefix of the logger.
// If a prefix is set and no trailing space is written, write one
func (r *Receiver) SetPrefix(prefix string) {
	if prefix != "" && !strings.HasSuffix(prefix, " ") {
		prefix += " "
	}
	r.logger.SetPrefix(prefix)
}

// Logs to the logger
func (r *Receiver) log(opt *levelOptions, s string) {
	// Don't do anything if not wanted
	if !r.Active || opt.Level < r.Level {
		return
	}

	// Pre- and suffix
	prefix := ""
	suffix := "\n"

	// Add colors if wanted
	if r.Color {
		prefix += fmt.Sprintf("\x1b[0;%sm", strconv.Itoa(opt.Color))
		suffix = "\x1b[0m" + suffix
	}

	// Print to the logger
	r.logger.Printf(prefix+r.Format+suffix, opt.Key, s)
}

// NewReceiver returns a new Receiver object with a given Writer
// and sets default values
func NewReceiver(w io.Writer, prefix string) *Receiver {

	logger := log.New(w, "", log.LstdFlags)
	r := &Receiver{
		logger: logger,
	}

	// Default options
	r.Active = true
	r.Level = INFO
	r.Format = "[%s] ▶ %s"

	r.SetPrefix(prefix)

	return r
}
