package mongo

import (
	"context"
	"fmt"

	"github.com/7d4b9/utrade/internal/api"
	utmongo "github.com/7d4b9/utrade/mongo"
	"github.com/google/uuid"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var Config = viper.New()

const (
	DBNameConfig = "db_name"
)

func init() {
	Config.AutomaticEnv()
	Config.SetEnvPrefix("mongo")
}

type Client struct {
	client   *utmongo.Client
	Invoices *mongo.Collection
}

func NewClient(ctx context.Context) (*Client, error) {
	client, err := utmongo.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("mongo new: %w", err)
	}
	dbName := Config.GetString(DBNameConfig)
	return &Client{
		client:   client,
		Invoices: client.Database(dbName).Collection("invoices"),
	}, nil
}

func (c *Client) Close(ctx context.Context) error {
	if err := c.client.Disconnect(ctx); err != nil {
		return fmt.Errorf("disconnect data client: %w", err)
	}
	return nil
}

// GetInvoice implements api.Data
func (d *Client) GetInvoice(ctx context.Context, id string) (*api.Invoice, error) {
	res := d.Invoices.FindOne(ctx, bson.D{
		{Key: "_id", Value: id},
	})
	if err := res.Err(); err == mongo.ErrNoDocuments {
		// generates api.Date expected error
		return nil, fmt.Errorf("data get invoice, no documents: %w", api.ErrNotFound)
	} else if err != nil {
		return nil, fmt.Errorf("data get invoice: %w", err)
	}
	var output api.Invoice
	if err := res.Decode(&output); err == mongo.ErrNoDocuments {
		// generates api expected error
		return nil, fmt.Errorf("data decode invoice, no documents: %w", api.ErrNotFound)
	} else if err != nil {
		return nil, fmt.Errorf("data decode invoice: %w", err)
	}
	return &output, nil
}

// CreateInvoice implements api.Data
func (d *Client) CreateInvoice(ctx context.Context, in *api.Invoice) (out *api.Invoice, err error) {
	if in.ID == "" {
		in.ID = uuid.New().String()
	}
	if _, err := d.Invoices.InsertOne(ctx, in); err != nil {
		if mongo.IsDuplicateKeyError(err) {
			// generates api expected error
			return nil, fmt.Errorf("data create invoice, duplicate key: %w", api.ErrAlreadyExists)
		} else if err != nil {
			return nil, fmt.Errorf("data create invoice: %w", err)
		}
	}
	return in, nil
}
