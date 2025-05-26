package main

import (
	"context"
	"time"

	"github.com/ose-micro/core/logger"
	"github.com/ose-micro/mongo"
	"github.com/ose-micro/mongo/example/collection"
)

func main() {
	logger, _ := logger.NewZap(logger.Config{
		Level: "info",
	})

	mdb, err := mongodb.New(mongodb.Config{
		Host:     "localhost",
		Port:     27019,
		User:     "fundraising",
		Password: "bcRqCvuAwPsbvriGXrIgSOdiuYbiGUyW",
		Database: "ose_mongo",
		Timeout:  3 * time.Second,
	}, logger)
	if err != nil {
		logger.Fatal(err.Error())
	}

	userCol := collection.UserCollection{
		Collection: mdb.Collection("users"),
	}

	userCol.Find(context.Background())
}
