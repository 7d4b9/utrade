package context

import (
	"context"
	"os"
	"os/signal"
)

// NewContext provide a context cancel on os signal SIGINT.
func NewContext() (context.Context, func()) {
	s := make(chan os.Signal, 1)
	signal.Notify(s, os.Interrupt)
	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan struct{})
	go func() {
		defer close(done)
		select {
		case <-s: // catched os signals
			cancel() // cancel returned context
		case <-ctx.Done(): // exit from returned func
		}
	}()
	return ctx, func() {
		cancel()
		<-done
	}
}
