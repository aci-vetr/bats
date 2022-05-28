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
