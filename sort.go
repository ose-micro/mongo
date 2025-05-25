package mongo

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Direction int

type Sort struct {
	Field     string
	Direction Direction
}

type Option func(*options.FindOptions)

func BuildSort(fields ...Sort) Option {
	return func(opts *options.FindOptions) {
		sort := bson.D{}
		for _, f := range fields {
			sort = append(sort, bson.E{Key: f.Field, Value: f.Direction})
		}
		opts.Sort = sort
	}
}
