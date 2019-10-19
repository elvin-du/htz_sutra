package model

import (
	"context"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"mdu/explorer/common/database"
	"mdu/explorer/data-syncer/config"
	"time"
)

const (
	DB_COLLECTION_VALIDATORS = "validators"
)

func NewValidatorModel() *ValidatorModel {
	vm := &ValidatorModel{database.DefaultDB.Client.Database(config.DefaultConfig.DBName).Collection(DB_COLLECTION_VALIDATORS)}
	return vm
}

type ValidatorModel struct {
	collection *mongo.Collection
}

type ValidatorQuery struct {
	PageIndex       int32
	PageSize        int32
	OperatorAddress string
}

func (vm *ValidatorModel) Upsert(val *Validator) error {
	exsit, err := vm.IsExist(val.OperatorAddress)
	if nil != err {
		log.Errorln(err)
		return err
	}

	if exsit {
		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		filter := bson.M{"operator_address": val.OperatorAddress}
		update := bson.M{"$set": val}
		_, err = vm.collection.UpdateOne(ctx, filter, update)
		if nil != err {
			log.Errorln(err)
			return err
		}

		return nil
	}

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	res, err := vm.collection.InsertOne(ctx, val)
	if nil != err {
		log.Errorln(err)
		return err
	}

	log.Debugln("insertid:", res.InsertedID)
	return nil
}

func (vm *ValidatorModel) IsExist(addr string) (bool, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	filter := bson.M{"operator_address": addr}
	result := vm.collection.FindOne(ctx, filter)
	if result.Err() == nil {
		return true, nil
	} else if result.Err() == mongo.ErrNoDocuments {
		return false, nil
	}

	return false, result.Err()
}

func (vm *ValidatorModel) Info(addr string) (*Validator, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	filter := bson.M{"operator_address": addr}
	result := vm.collection.FindOne(ctx, filter)
	log.Debugf("result:%+v", result)
	if result.Err() != nil {
		log.Errorln(result.Err())
		return nil, result.Err()
	}

	val := &Validator{}
	err := result.Decode(val)
	if err != nil {
		log.Errorln(err)
		return nil, err
	}

	return val, nil
}

func (vm *ValidatorModel) Count(query ValidatorQuery) (int64, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	filter := vm.buildFilterByQuery(query)
	totalCount, err := vm.collection.CountDocuments(ctx, filter)
	if err != nil {
		log.Errorln(err)
		return -1, err
	}
	return totalCount, err
}

func (vm *ValidatorModel) List(query ValidatorQuery) ([]Validator, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	filter := vm.buildFilterByQuery(query)
	findOps := vm.buildOptionsByQuery(query)

	cur, err := vm.collection.Find(ctx, filter, findOps)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	results := make([]Validator, 20)
	err = cur.All(ctx, &results)
	if err != nil {
		log.Errorln(err)
		return nil, err
	}
	return results, nil
}

func (vm *ValidatorModel) buildFilterByQuery(query ValidatorQuery) *bson.M {
	filter := bson.M{}
	if query.OperatorAddress != "" {
		filter["operator_address"] = query.OperatorAddress
	}
	return &filter
}

func (vm *ValidatorModel) buildOptionsByQuery(query ValidatorQuery) *options.FindOptions {
	findOps := options.Find()
	min := query.PageIndex * query.PageSize
	max := min + query.PageSize
	findOps.Min, findOps.Max = min, max
	return findOps
}
