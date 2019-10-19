package model

import (
	"context"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func FindOne(collection *mongo.Collection, val interface{}, filter interface{}, opts ...*options.FindOneOptions) error {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result := collection.FindOne(ctx, filter, opts...)
	if result.Err() != nil {
		log.Errorln(result.Err())
		return result.Err()
	}
	err := result.Decode(val)
	if err != nil {
		log.Errorln(result.Err())
		return err
	}
	return nil
}
