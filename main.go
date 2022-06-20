package main

import (
	"github.com/7d4b9/utrade/context"
	"github.com/7d4b9/utrade/internal/api"
	"github.com/7d4b9/utrade/internal/mongo"
)

func main() {
	ctx, cancel := context.NewContext()
	defer cancel()
	// Create a new database and connect to the server
	database, err := mongo.NewClient(ctx)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = database.Close(ctx); err != nil {
			panic(err)
		}
	}()
	// server
	api.RunServer(ctx, database)
}
