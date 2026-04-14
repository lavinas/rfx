package driven

import (
	"log"
)

// SimpleLogger is a simple wrapper around the standard log.Logger
type SimpleLogger struct {
	*log.Logger
}

// NewSimpleLogger creates a new instance of SimpleLogger
func NewSimpleLogger() *SimpleLogger {
	return &SimpleLogger{Logger: log.Default()}
}
