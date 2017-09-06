package logo

// Level is a log level
type Level int

// Available log levels
const (
	DEBUG Level = iota
	INFO
	WARN
	ERROR
	FATAL
)

// Options to store the key, level and color code of a log
type levelOptions struct {
	Key   string
	Level Level
	Color int
}

// Log options for all levels
var (
	optDebug = &levelOptions{"DEBU", DEBUG, 34}
	optInfo  = &levelOptions{"INFO", INFO, 32}
	optWarn  = &levelOptions{"WARN", WARN, 33}
	optError = &levelOptions{"ERRO", ERROR, 31}
	optFatal = &levelOptions{"CRIT", FATAL, 31}
)

// Itol converts an integer to a logo.Level
// TODO is it possible to cast somehow?
func Itol(level int) Level {
	switch level {
	case 0:
		return DEBUG
	case 1:
		return INFO
	case 2:
		return WARN
	case 3:
		return ERROR
	case 4:
		return FATAL
	default:
		return DEBUG
	}
}
