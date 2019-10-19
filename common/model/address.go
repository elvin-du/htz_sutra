package model

import (
	"context"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"mdu/explorer/common/database"
	"mdu/explorer/data-syncer/config"
	"time"
)

var (
	DB_COLLECTION_ADDRESS = "addresses"
)

func NewAddressesModel() *AddressesModel {
	am := &AddressesModel{database.DefaultDB.Client.Database(config.DefaultConfig.DBName).Collection(DB_COLLECTION_ADDRESS)}
	return am
}

type AddressesModel struct {
	collection *mongo.Collection
}

func (bm *AddressesModel) Find(address string) (*Address, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result := bm.collection.FindOne(ctx, bson.M{"address": address})
	if result.Err() != nil {
		log.Errorln(result.Err())
		return nil, result.Err()
	}
	addr := &Address{}
	err := result.Decode(addr)
	if err != nil {
		log.Errorln(result.Err())
		return nil, err
	}
	return addr, nil
}

func (bm *AddressesModel) Upsert(addr *Address) error {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	filter := bson.M{"address": addr.Address}
	update := bson.M{"$set": addr}
	_, err := bm.collection.UpdateOne(ctx, filter, update)
	if nil != err {
		log.Errorln(err)
		return err
	}

	return nil
}
