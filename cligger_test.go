package cligger

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/guumaster/logsymbols"
	"github.com/stretchr/testify/assert"
)

func TestNewLogger(t *testing.T) {
	l := NewLogger()
	i := l.(*Log) // Cast as Log
	assert.Equal(t, *i.s, logsymbols.CurrentSymbols())
}

func TestLogErrorf(t *testing.T) {
	l := NewLogger()
	e := l.Errorf("wrapper: %w", errors.New("internal"))

	assert.EqualError(t, e, "[✖] wrapper: internal")
}

func TestLogEnableColor(t *testing.T) {
	b := bytes.NewBufferString("")
	l := NewLoggerWithWriter(b)
	l.EnableColor()
	l.Info("Info color test")

	s, err := ioutil.ReadAll(b)
	assert.NoError(t, err)

	c := logsymbols.Colorize(logsymbols.BaseSymbols())

	assert.Equal(t, fmt.Sprintf("[%s] Info color test\n", c.Info), string(s))
}

func TestLogNewLoggerWithWriter(t *testing.T) {
	t.Run("Info", func(t *testing.T) {
		b := bytes.NewBufferString("")
		l := NewLoggerWithWriter(b)
		l.DisableColor()

		l.Info("Info Test")

		s, err := ioutil.ReadAll(b)
		assert.NoError(t, err)

		assert.Equal(t, string(s), "[ℹ] Info Test\n")
	})
	t.Run("Success", func(t *testing.T) {
		b := bytes.NewBufferString("")
		l := NewLoggerWithWriter(b)
		l.DisableColor()

		l.Success("Success Test")

		s, err := ioutil.ReadAll(b)
		assert.NoError(t, err)

		assert.Equal(t, string(s), "[✔] Success Test\n")
	})

	t.Run("Warning", func(t *testing.T) {
		b := bytes.NewBufferString("")
		l := NewLoggerWithWriter(b)
		l.DisableColor()

		l.Warning("Warning Test")

		s, err := ioutil.ReadAll(b)
		assert.NoError(t, err)

		assert.Equal(t, string(s), "[⚠] Warning Test\n")
	})

	t.Run("Error", func(t *testing.T) {
		b := bytes.NewBufferString("")
		l := NewLoggerWithWriter(b)
		l.DisableColor()

		l.Error("Error Test")

		s, err := ioutil.ReadAll(b)
		assert.NoError(t, err)

		assert.Equal(t, string(s), "[✖] Error Test\n")
	})
}
