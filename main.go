package main

import (
	"github.com/7d4b9/utrade/context"
	api "github.com/7d4b9/utrade/internal/api"
	"github.com/7d4b9/utrade/internal/data"
	"github.com/7d4b9/utrade/internal/mongo"
)

func main() {
	ctx, cancel := context.New()
	defer cancel()
	// Create a new client and connect to the server
	client, err := mongo.NewClient(ctx)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	data, err := data.NewData(client.Database())
	if err != nil {
		panic(err)
	}
	api.RunServer(ctx, data)
}
