package context

import (
	"context"
	"os"
	"os/signal"
)

// NewContext provides a context automatically canceled on os signal SIGINT.
// Canceling this context releases resources associated with it, so code should
// call cancel as soon as the operations running in this Context complete.
func NewContext() (ctx context.Context, cancel context.CancelFunc) {
	ctx, cancel = context.WithCancel(context.Background())
	s := make(chan os.Signal, 1)
	signal.Notify(s, os.Interrupt)
	done := make(chan struct{})
	go func() {
		defer close(done)
		select {
		case <-s: // Interrupt
			// cancel returned context
			cancel()
		case <-ctx.Done(): // primary exit triggered from returned stop function
		}
	}()
	// overload primary cancel to release local resources.
	cancel = func() {
		cancel()
		<-done
	}
	return ctx, cancel
}
