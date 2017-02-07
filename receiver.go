package logo

import (
    "io"
    "log"
)

type Receiver struct {
    // Logger object from the builtin log package
    logger *log.Logger
    Level level
    Color bool
    Active bool
}

// Logs to the logger
func (r *Receiver) log(l level, key string, color int, a interface{}) {
    if !r.Active {
        return
    }

    if l < r.Level {
        return
    }

    // Add colors if wanted
    if r.Color {
        r.logger.Printf("\x1b[0;%dm[%s] ▶ %v\x1b[0m\n", color, key, a)
        return
    }

    // Print to the logger
    r.logger.Printf("[%s] ▶ %s\n", key, a)
}

func NewReceiver(w io.Writer) *Receiver {
    r := &Receiver{
        logger: log.New(w, "", log.LstdFlags),
    }
    // Default options
    r.Active = true
    r.Level = INFO

    return r
}
