package api

import (
	"context"
	"errors"

	"github.com/7d4b9/utrade/http"
	"github.com/labstack/echo/v4"
)

// RunServer configures and starts the http server.
func RunServer(ctx context.Context, data Data) {
	e := echo.New()
	addRoutes(e, data)
	http.RunServer(ctx, e)
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

const (
	workspaceParam = "workspace"
	invoiceParam   = "invoice"
)

func addRoutes(e *echo.Echo, data Data) {
	e.GET("/api/v1/workspace/:"+workspaceParam+"/invoice/:"+invoiceParam+"", getInvoice(data))
	e.POST("/api/v1/workspace/:"+workspaceParam+"/invoices", createInvoice(data))
}
