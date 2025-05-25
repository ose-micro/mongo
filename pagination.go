package mongo

import "go.mongodb.org/mongo-driver/mongo/options"

func WithLimit(limit int64) Option {
	return func(opts *options.FindOptions) {
		opts.Limit = &limit
	}
}

func WithSkip(skip int64) Option {
	return func(opts *options.FindOptions) {
		opts.Skip = &skip
	}
}
