package context

import (
	"context"
	"os"
	"os/signal"
)

func New() (context.Context, func()) {
	s := make(chan os.Signal, 1)
	signal.Notify(s, os.Interrupt)
	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan struct{})
	go func() {
		defer close(done)
		select {
		case <-s: // catched os signals
		case <-ctx.Done(): // exit from cancel func
		}
	}()
	return ctx, func() {
		cancel()
		<-done
	}
}
