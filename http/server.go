package http

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

func RunServer(ctx context.Context, e *echo.Echo) {

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

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
		<-exited
	case <-exited:
	}
}
