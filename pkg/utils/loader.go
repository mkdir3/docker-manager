package utils

import (
	"time"

	"github.com/briandowns/spinner"
)

type Loader struct {
	spinner *spinner.Spinner
}

func NewLoader() *Loader {
	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	return &Loader{spinner: s}
}

func (l *Loader) Start(message string) {
	l.spinner.Prefix = message
	l.spinner.Start()
	// Below pending only for test purposes
	time.Sleep(2 * time.Second)
}

func (l *Loader) Stop() {
	l.spinner.Stop()
}
