package cligger

import (
	"io"

	"github.com/guumaster/logsymbols"
)

type Logger interface {
	Info(format string, args ...interface{})
	Success(pattern string, args ...interface{})
	Warning(pattern string, args ...interface{})
	Error(pattern string, args ...interface{})
	Errorf(pattern string, args ...interface{}) error
	Fatal(pattern string, args ...interface{})
	EnableColor()
	DisableColor()
}

type Log struct {
	w io.Writer
	s *logsymbols.Symbols
}
