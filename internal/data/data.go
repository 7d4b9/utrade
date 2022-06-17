package data

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/7d4b9/utrade/internal/api"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Data provide high level data.
type Data struct {
	invoices *mongo.Collection
}

func NewData(db *mongo.Collection) (*Data, error) {
	return &Data{invoices: db}, nil
}

// GetInvoice implements api.Data
func (d *Data) GetInvoice(ctx context.Context, id string) (*api.Invoice, error) {
	res := d.invoices.FindOne(ctx, bson.D{
		{Key: "_id", Value: id},
	})
	if err := res.Err(); errors.Is(err, mongo.ErrNoDocuments) {
		// generates api.Date expected error
		return nil, fmt.Errorf("data get invoice, no documents: %w", api.ErrNotFound)
	} else if err != nil {
		return nil, fmt.Errorf("data get invoice: %w", err)
	}
	var output api.Invoice
	if err := res.Decode(&output); errors.Is(err, mongo.ErrNoDocuments) {
		// generates api.Date expected error
		return nil, fmt.Errorf("data decode invoice, no documents: %w", api.ErrNotFound)
	} else if err != nil {
		return nil, fmt.Errorf("data decode invoice: %w", err)
	}
	return &output, nil
}

// CreateInvoice implements api.Data
func (d *Data) CreateInvoice(ctx context.Context, in *api.Invoice) (out *api.Invoice, err error) {
	if in.ID == "" {
		in.ID = uuid.New().String()
	}
	if _, err := d.invoices.InsertMany(ctx, []any{*in}); err != nil {
		if mongo.IsDuplicateKeyError(err) {
			// generates api.Date expected error
			return nil, fmt.Errorf("data create invoice, duplicate key: %w", api.ErrAlreadyExists)
		} else if err != nil {
			return nil, fmt.Errorf("data create invoice: %w", err)
		}
	}
	return in, nil
}
