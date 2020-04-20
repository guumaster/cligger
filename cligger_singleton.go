// nolint:goprintffuncname
package cligger

import (
	"io"
	"sync"

	"github.com/guumaster/logsymbols"
)

var instance *Log
var once sync.Once

func init() { // nolint:gochecknoinits
	once.Do(func() {
		l := NewLogger()
		instance = l.(*Log)
	})
}

// SetWriter set a new internal writer to the singleton instance
func SetWriter(w io.Writer) {
	instance.w = w
}

// DisableColor strip colors from the singleton instance symbols
func DisableColor() {
	s := logsymbols.BaseSymbols()
	instance.s = &s
}

// EnableColor adds colors to the singleton instance symbols
func EnableColor() {
	s := logsymbols.Colorize(logsymbols.BaseSymbols())
	instance.s = &s
}

// Info singleton method to log a string with an info symbol
func Info(format string, args ...interface{}) {
	instance.Info(format, args...)
}

// Success singleton method to log a string with an success symbol
func Success(format string, args ...interface{}) {
	instance.Success(format, args...)
}

// Warning singleton method to log a string with an warning symbol
func Warning(format string, args ...interface{}) {
	instance.Warning(format, args...)
}

// Error singleton method to log a string with an error symbol
func Error(format string, args ...interface{}) {
	instance.Error(format, args...)
}

// Errorf singleton method to get a wrapped error with an error symbol
func Errorf(format string, args ...interface{}) error {
	return instance.Errorf(format, args...)
}

// Fatal singleton method to log a message and exit the program
func Fatal(format string, args ...interface{}) {
	instance.Fatal(format, args...)
}
