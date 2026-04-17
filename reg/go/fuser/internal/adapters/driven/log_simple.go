package driven

import (
	"log"
	"strings"
)

// SimpleLogger is a simple wrapper around the standard log.Logger
type SimpleLogger struct {
	*log.Logger
}

// NewSimpleLogger creates a new instance of SimpleLogger
func NewSimpleLogger() *SimpleLogger {
	return &SimpleLogger{Logger: log.Default()}
}

// IPrintf is a identity function for printf, allowing it to be used as a ports.Logger
func (s *SimpleLogger) IPrintf(tabs int, format string, v ...interface{}) {
	format = strings.Repeat("\t", tabs) + format
	s.Printf(format, v...)
}
