package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	log "github.com/sirupsen/logrus"
	"time"
)

var (
	DefaultDB = &database{}
)

type database struct {
	Client *mongo.Client
}

func (db *database) Start(uri string) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if nil != err {
		panic(err)
	}
	db.Client = client

	//check if connection is successful
	ctx, _ = context.WithTimeout(context.Background(), 2*time.Second)
	err = client.Ping(ctx, readpref.Primary())
	if nil != err {
		panic(err)
	}
	log.Infoln("connect mongo success")
}
