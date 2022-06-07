package logger

import (
	"io"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	a := assert.New(t)

	log, err := New(Config{
		Filename:   "logger-test",
		ConsoleOut: io.Discard,
	})
	a.NoError(err)
	log.Info().Msg("info")
	log.Debug().Msg("debug")
	content, _ := os.ReadFile("logger-test")
	a.True(strings.Contains(string(content), "info"))
	a.True(strings.Contains(string(content), "debug"))
	os.Remove("logger-test")
}

func TestConsoleLevel(t *testing.T) {
	a := assert.New(t)

	// Default level
	var b strings.Builder
	log, _ := New(Config{
		FileOut:    io.Discard,
		ConsoleOut: &b,
	})
	log.Info().Msg("test")
	a.Contains(b.String(), "test")

	// Higher level
	b.Reset()
	log, _ = New(Config{
		FileOut:      io.Discard,
		ConsoleOut:   &b,
		ConsoleLevel: ErrorLevel,
	})
	log.Info().Msg("test")
	a.NotContains(b.String(), "test")
}
