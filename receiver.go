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

	// if a prefix is set and no trailing space is written, write one
	if prefix != "" && !strings.HasSuffix(prefix, " ") {
		prefix += " "
	}

	r := &Receiver{
		logger: log.New(w, prefix, log.LstdFlags),
	}
	// Default options
	r.Active = true
	r.Level = INFO
	r.Format = "[%s] ▶ %s"

	return r
}
