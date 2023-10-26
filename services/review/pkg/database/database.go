package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/hossein1376/BehKhan/review/pkg/config"
)

func OpenDB(cfg *config.Settings) (*mongo.Collection, func() error, error) {
	dsn := fmt.Sprintf(
		"mongodb://%s:%s@%s:%s",
		cfg.DB.Username,
		cfg.DB.Password,
		cfg.DB.Host,
		cfg.DB.Port,
	)

	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(dsn).SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		return nil, nil, err
	}

	// Send a ping to confirm a successful connection
	if err = client.Database(cfg.DB.Name).RunCommand(context.Background(), bson.D{{"ping", 1}}).Err(); err != nil {
		return nil, nil, err
	}

	// Disconnect function
	disconnect := func() error {
		if err = client.Disconnect(context.Background()); err != nil {
			return err
		}
		return nil
	}

	return client.Database(cfg.DB.Name).Collection(cfg.DB.Collection), disconnect, nil
}
