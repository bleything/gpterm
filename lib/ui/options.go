package ui

import (
	"io"

	"github.com/collinvandyck/gpterm/lib/log"
)

type Option func(*console)

func WithLogger(logger log.Logger) Option {
	return func(c *console) {
		c.Logger = logger
	}
}

func WithInput(reader io.Reader) Option {
	return func(c *console) {
		c.input = reader
	}
}
func WithOutput(writer io.Writer) Option {
	return func(c *console) {
		c.output = writer
	}
}
