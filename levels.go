package logo

type level int

// Available log levels
const (
	DEBUG level = iota
	INFO
	WARN
	ERROR
	FATAL
)

// Options to store the key, level and color code of a log
type levelOptions struct {
	Key   string
	Level level
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
