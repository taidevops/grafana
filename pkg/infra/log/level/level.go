package level

import "github.com/go-kit/log"

// Error returns a logger that includes a Key/ErrorValue pair
func Error(logger log.Logger) log.Logger {
	return log.WritePrefix(logger, Key(), ErrorValue())
}

// NewFilter wraps next and implements level filtering. See the commentary on
// the Option functions for a detailed description of how to configure levels.
// If no options are provided, all leveled log events created with Debug,
// Info, Warn or Error helper methods are squelched and non-leveled log
// events are passed to next unmodified.
func NewFilter(next log.Logger, options ...Option) log.Logger {
	l := &logger{
		next: next,
	}
	for _, option := range options {
		option(l)
	}
	return l
}

func (l *logger) Log(keyvals ...interface{}) error {
	var hasLevel, levelAllowed bool
	for i := 1; i < len(keyvals); i += 2 {
		if v, ok := keyvals[i].(*levelValue); ok {
			hasLevel = true
			levelAllowed = l.allowed&v.level != 0
			break
		}
	}
	if !hasLevel && l.squelchNoLevel {
		return l.errNoLevel
	}
	if hasLevel && !levelAllowed {
		return l.errNotAllowed
	}
	return l.next.Log(keyvals...)
}

type logger struct {
	next           log.Logger
	allowed        level
	squelchNoLevel bool
	errNotAllowed  error
	errNoLevel     error
}

// Option sets a parameter for the leveled logger.
type Option func(*logger)

// Value is the interface that each of the canonical level values implement.
// It contains unexported methods that prevent types from other packages from
// implementing it and guaranteeing that NewFilter can distinguish the levels
// defined in this package from all other values.
type Value interface {
	String() string
	levelVal()
}

// Key returns the unique key added to log events by the loggers in this
// package.
func Key() interface{} { return key }

// ErrorValue returns the unique value added to log events by Error.
func ErrorValue() Value { return errorValue }

// WarnValue returns the unique value added to log events by Warn.
func WarnValue() Value { return warnValue }

var (
	// key is of type interface{} so that it allocates once during package
	// initialization and avoids allocating every time the value is added to a
	// []interface{} later.
	key interface{} = "lvl"

	errorValue = &levelValue{level: levelError, name: "eror"}
	warnValue  = &levelValue{level: levelWarn, name: "warn"}
)

type level byte

const (
	levelDebug level = 1 << iota
	levelInfo
	levelWarn
	levelError
)

type levelValue struct {
	name string
	level
}
