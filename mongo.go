package mongo

import (
	"context"
	"fmt"
	"time"

	"github.com/ose-micro/core/logger"
	"github.com/ose-micro/core/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Client struct {
	db     *mongo.Database
	logger logger.Logger
}

func New(conf Config, logger logger.Logger) (*Client, error) {
	var mClient *mongo.Client

	err := utils.Retry(10*time.Second, 5, func() error {
		uri := fmt.Sprintf(
			"mongodb://%s:%s@%s:%d",
			conf.User,
			conf.Password,
			conf.Host,
			conf.Port,
		)
		// Created a new client and connect to the server
		client, err := mongo.NewClient(options.Client().ApplyURI(uri))
		if err != nil {
			logger.Panic(fmt.Sprintf("Failed to create MongoDB client: %v", err))
			panic(err)
		}

		// Context with timeout for connecting to MongoDB
		ctx, cancel := context.WithTimeout(context.Background(), conf.Timeout)
		defer cancel()

		// Connect to MongoDB
		err = client.Connect(ctx)
		if err != nil {
			logger.Panic(fmt.Sprintf("Failed to connect to MongoDB: %v", err))
		}

		// Ping the database to verify the connection
		err = client.Ping(ctx, nil)
		if err != nil {
			logger.Error(fmt.Sprintf("Failed to ping MongoDB: %v", err))
			panic(err)
		}

		mClient = client
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &Client{
		db:     mClient.Database(conf.Database),
		logger: logger,
	}, nil
}

func (c *Client) Collection(name string) *mongo.Collection {
	return c.db.Collection(name)
}
