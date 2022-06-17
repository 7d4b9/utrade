package data

import (
	"context"

	"github.com/7d4b9/utrade/internal/api"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
)

// Data provide high level data.
type Data struct {
	db *mongo.Database
}

func NewData(db *mongo.Database) (*Data, error) {
	return &Data{db: db}, nil
}

// GetInvoice implements http.Data.
func (d *Data) GetInvoice(ctx context.Context, id string) (out *api.Invoice, err error) {
	return nil, errors.WithMessage(api.ErrNotFound, "get invoice from db")
}

// CreateInvoice implements http.Data.
func (d *Data) CreateInvoice(ctx context.Context, in *api.Invoice) (out *api.Invoice, err error) {
	return &api.Invoice{}, nil
}
