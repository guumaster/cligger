// nolint:goprintffuncname
package cligger

import (
	"fmt"
	"io"
	"os"

	"github.com/guumaster/logsymbols"
)

// NewLogger new logger with stdout writer
func NewLogger() Logger {
	s := logsymbols.CurrentSymbols()

	return &Log{
		w: os.Stdout,
		s: &s,
	}
}

// NewLoggerWithWriter new logger with a specific writer
func NewLoggerWithWriter(w io.Writer) Logger {
	s := logsymbols.CurrentSymbols()

	return &Log{
		s: &s,
		w: w,
	}
}

// Info logs a string with an info symbol
func (p Log) Info(format string, args ...interface{}) {
	_, _ = fmt.Fprintf(p.w, p.prefixer(p.s.Info, format), args...)
}

// Success logs a string with a success symbol
func (p Log) Success(format string, args ...interface{}) {
	_, _ = fmt.Fprintf(p.w, p.prefixer(p.s.Success, format), args...)
}

// Warning logs a string with a warning symbol
func (p Log) Warning(format string, args ...interface{}) {
	_, _ = fmt.Fprintf(p.w, p.prefixer(p.s.Warning, format), args...)
}

// Error logs a string with an error symbol
func (p Log) Error(format string, args ...interface{}) {
	_, _ = fmt.Fprintf(p.w, p.prefixer(p.s.Error, format), args...)
}

// Errorf returns an error with a message containing an error symbol
func (p Log) Errorf(format string, args ...interface{}) error {
	return fmt.Errorf(fmt.Sprintf("[%s] ", p.s.Error)+format, args...)
}

// Fatal logs an error  containing an error symbol and exit the program
func (p Log) Fatal(format string, args ...interface{}) {
	_, _ = fmt.Fprintf(p.w, fmt.Sprintf("[%s] ", p.s.Error)+format, args...)

	os.Exit(1)
}

// DisableColor strip TTY colors from symbols
func (p *Log) DisableColor() {
	s := logsymbols.BaseSymbols()
	p.s = &s
}

// EnableColor add TTY colors from symbols
func (p *Log) EnableColor() {
	s := logsymbols.Colorize(logsymbols.BaseSymbols())
	p.s = &s
}

func (p Log) prefixer(s logsymbols.Symbol, f string) string {
	return fmt.Sprintf("[%s] ", s) + f + "\n"
}
