package mongo

import (
	"context"
	"fmt"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var Config = viper.New()

const (
	URIConfig    = "uri"
	DBNameConfig = "db_name"
)

func init() {
	Config.AutomaticEnv()
	Config.SetEnvPrefix("mongo")
}

type Client struct {
	*mongo.Client
	dbName string
}

func NewClient(ctx context.Context) (*Client, error) {
	uri := Config.GetString(URIConfig)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, fmt.Errorf("mongo connect: %w", err)
	}
	// Ping the primary
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, fmt.Errorf("mongo ping: %w", err)
	}
	dbName := Config.GetString(DBNameConfig)
	return &Client{Client: client, dbName: dbName}, nil
}

func (c *Client) Invoices() *mongo.Collection {
	return c.Client.Database(c.dbName).Collection("invoices")
}
