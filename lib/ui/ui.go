package ui

import (
	"context"
	"io"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/collinvandyck/gpterm/lib/client"
	"github.com/collinvandyck/gpterm/lib/log"
	"github.com/collinvandyck/gpterm/lib/store"
)

type UI interface {
	// blocks until the program exits
	Run(ctx context.Context) error
}

func New(store *store.Store, client client.Client, opts ...Option) UI {
	console := &console{
		uiOpts: uiOpts{
			input:         os.Stdin,
			output:        os.Stdout,
			Logger:        log.Discard,
			store:         store,
			client:        client,
			styles:        newStaticStyles(),
			clientTimeout: time.Minute,
			rhsPadding:    2,
		},
	}
	for _, o := range opts {
		o(console)
	}
	return console
}

type uiOpts struct {
	log.Logger
	input         io.Reader
	output        io.Writer
	store         *store.Store
	client        client.Client
	styles        styles
	clientTimeout time.Duration // how long to wait for a response
	rhsPadding    int           // RHS padding for rendered markdown
}

func (uiOpts uiOpts) NamedLogger(prefix string) uiOpts {
	uiOpts.Logger = uiOpts.Logger.New("name", prefix)
	return uiOpts
}

type console struct {
	uiOpts
}

func (t *console) Run(ctx context.Context) error {
	t.Log("+-------------------+")
	t.Log("| gpterm starting...|")
	t.Log("+-------------------+")
	model := newControlModel(t.uiOpts)
	p := tea.NewProgram(model, tea.WithInput(t.input), tea.WithOutput(t.output))
	_, err := p.Run()
	return err
}
