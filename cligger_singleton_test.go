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

func TestErrorf(t *testing.T) {
	e := Errorf("wrapper: %w", errors.New("internal"))

	assert.EqualError(t, e, "[✖] wrapper: internal")
}

func TestEnableColor(t *testing.T) {
	b := bytes.NewBufferString("")
	SetWriter(b)
	EnableColor()
	Info("Info color test")

	s, err := ioutil.ReadAll(b)
	assert.NoError(t, err)

	c := logsymbols.Colorize(logsymbols.BaseSymbols())

	assert.Equal(t, fmt.Sprintf("[%s] Info color test\n", c.Info), string(s))
}

func TestNewLoggerWithWriter(t *testing.T) {
	t.Run("Info", func(t *testing.T) {
		b := bytes.NewBufferString("")
		SetWriter(b)
		DisableColor()

		Info("Info Test")

		s, err := ioutil.ReadAll(b)
		assert.NoError(t, err)

		assert.Equal(t, string(s), "[ℹ] Info Test\n")
	})
	t.Run("Success", func(t *testing.T) {
		b := bytes.NewBufferString("")
		SetWriter(b)
		DisableColor()

		Success("Success Test")

		s, err := ioutil.ReadAll(b)
		assert.NoError(t, err)

		assert.Equal(t, string(s), "[✔] Success Test\n")
	})

	t.Run("Warning", func(t *testing.T) {
		b := bytes.NewBufferString("")
		SetWriter(b)
		DisableColor()

		Warning("Warning Test")

		s, err := ioutil.ReadAll(b)
		assert.NoError(t, err)

		assert.Equal(t, string(s), "[⚠] Warning Test\n")
	})

	t.Run("Error", func(t *testing.T) {
		b := bytes.NewBufferString("")
		SetWriter(b)
		DisableColor()

		Error("Error Test")

		s, err := ioutil.ReadAll(b)
		assert.NoError(t, err)

		assert.Equal(t, string(s), "[✖] Error Test\n")
	})
}
