package api

import (
	"context"
	"errors"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

var Config = viper.New()

const (
	shutdownTimeoutConfig = "shutdown_timeout"
)

func init() {
	Config.SetDefault(shutdownTimeoutConfig, "10s")
	Config.SetEnvPrefix("http")
	Config.AutomaticEnv()
}

var (
	// ErrAlreadyExists should be returned if the data creation
	// produces a conflict or if the data already exists.
	ErrAlreadyExists = errors.New("already exists")
	// ErrNotFound should be returned if the data iInvoices not found.
	ErrNotFound = errors.New("not found")
)

// Data defines how data are retrieved and stored.
type Data interface {
	// GetInvoice returns an Invoice not nil and a nil error in case of success.
	// If the invoice can not be found ErrNotFound should be returned and a nil Invoice.
	GetInvoice(ctx context.Context, id string) (out *Invoice, err error)
	// CreateInvoice create a new invoice. If the invoice cannot be created
	// due to a conflict with an existing data ErrAlreadyExists is expected.
	CreateInvoice(ctx context.Context, in *Invoice) (out *Invoice, err error)
}

func RunServer(ctx context.Context, repository Data) {

	// Echo instance
	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/api/v1/workspace/:workspace/invoice/:invoice_id", getInvoice(repository))
	e.POST("/api/v1/workspace/:workspace/invoices", createInvoice(repository))

	// Start server
	exited := make(chan struct{})
	go func() {
		defer close(exited)
		e.Logger.Printf("http server exited: %w", e.Start(":1323"))
	}()
	shutdownTimeout := Config.GetDuration(shutdownTimeoutConfig)
	shutdownCtx, cancelShutdown := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancelShutdown()
	select {
	case <-ctx.Done():
		e.Logger.Printf("http server shutdown: %w", e.Shutdown(shutdownCtx))
	case <-exited:
	}
	<-exited
}
