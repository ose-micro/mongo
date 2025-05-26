package mongodb

import (
	"time"

	"github.com/ose-micro/core/dto"
	"go.mongodb.org/mongo-driver/bson"
)

func BuildFilter(params []dto.Filter) bson.M {
	query := bson.M{}

	for _, filter := range params {
		field := filter.Field
		value := filter.Value

		switch filter.Operator {
		case dto.EQUAL:
			query[field] = bson.M{"$eq": value}
		case dto.LIKE:
			query[field] = bson.M{"$regex": value, "$options": "i"}
		case dto.BETWEEN:
			if values, ok := value.([]interface{}); ok && len(values) == 2 {
				query[field] = bson.M{"$gte": values[0], "$lte": values[1]}
			}
		case dto.GREATER_THAN:
			query[field] = bson.M{"$gt": value}
		case dto.LESS_THAN:
			query[field] = bson.M{"$lt": value}
		case dto.GREATER_THAN_EQUAL:
			query[field] = bson.M{"$gte": value}
		case dto.LESS_THAN_EQUAL:
			query[field] = bson.M{"$lte": value}
		case dto.DATE_EQUAL:
			query[field] = bson.M{"$gt": value}
		case dto.BEFORE:
			query[field] = bson.M{"$lt": value}
		case dto.AFTER:
			query[field] = bson.M{"$eq": value}
		case dto.DATE_BETWEEN:
			if values, ok := value.([]time.Time); ok && len(values) == 2 {
				query[field] = bson.M{"$gte": values[0], "$lte": values[1]}
			}
		}

	}

	return query
}
