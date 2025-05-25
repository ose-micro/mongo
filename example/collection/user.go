package collection

import (
	"context"
	"log"

	"github.com/ose-micro/core/dto"
	"github.com/ose-micro/mongo"
	mdb "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	Id      string  `bson:"_id"`
	Name    string  `bson:"name"`
	Account float64 `bson:"account"`
}

type UserCollection struct {
	Collection *mdb.Collection
}

func (u UserCollection) Insert(ctx context.Context, user User) error {
	_, err := u.Collection.InsertOne(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (u UserCollection) Find(ctx context.Context) error {
	filter := mongo.BuildFilter([]dto.Filter{
		{Field: "name", Operator: dto.LIKE, Value: "Kargbo"},
	})

	sort := mongo.BuildSort(
		mongo.Sort{Field: "name", Direction: 1},
	)

	limitValue := int64(2)
	limit := mongo.WithLimit(limitValue)
	skipValue := int64(0)
	skip := mongo.WithSkip(skipValue)

	findOpts := options.Find()
	sort(findOpts)
	limit(findOpts)
	skip(findOpts)

	cursor, err := u.Collection.Find(ctx, filter, findOpts)
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var user User
		if err := cursor.Decode(&user); err != nil {
			return err
		}

		log.Println(user)
	}

	if err := cursor.Err(); err != nil {
		return err
	}

	return nil
}
