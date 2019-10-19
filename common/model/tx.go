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
	DB_COLLECTION_TRANSACTIONS = "transactions"
)

func NewTxModel() *TxModel {
	tm := &TxModel{database.DefaultDB.Client.Database(config.DefaultConfig.DBName).Collection(DB_COLLECTION_TRANSACTIONS)}
	return tm
}

type TxModel struct {
	collection *mongo.Collection
}

type TxQuery struct {
	PageIndex int32
	PageSize  int32
	BlockHash string
	From      string
	To        string
}

func (tm *TxModel) Insert(tx *Tx) error {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	res, err := tm.collection.InsertOne(ctx, tx)
	if nil != err {
		log.Errorln(err)
		return err
	}

	log.Debugln("insertid:", res.InsertedID)
	return nil
}

func (bm *TxModel) FindTxByHash(hash string) (*Tx, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result := bm.collection.FindOne(ctx, bson.M{"tx_hash": hash})
	if result.Err() != nil {
		log.Errorln(result.Err())
		return nil, result.Err()
	}
	tx := &Tx{}
	err := result.Decode(tx)
	if err != nil {
		log.Errorln(result.Err())
		return nil, err
	}
	return tx, nil
}

func (bm *TxModel) CountTx(query TxQuery) (int64, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	filter := bm.buildFilterByQuery(query)
	totalCount, err := bm.collection.CountDocuments(ctx, filter)
	if err != nil {
		log.Errorln(err)
		return -1, err
	}
	return totalCount, err
}

func (bm *TxModel) List(query TxQuery) ([]*Tx, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	filter := bm.buildFilterByQuery(query)
	findOps := bm.buildOptionsByQuery(query)

	cur, err := bm.collection.Find(ctx, filter, findOps)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	results := make([]*Tx, 20)
	err = cur.All(ctx, &results)
	if err != nil {
		log.Errorln(err)
		return nil, err
	}
	return results, nil
}

func (bm *TxModel) buildFilterByQuery(query TxQuery) *bson.M {
	filter := bson.M{}
	if query.BlockHash != "" {
		filter["block_hash"] = query.BlockHash
	}
	if query.From != "" {
		filter["from"] = query.From
	}
	if query.To != "" {
		filter["to"] = query.To
	}

	return &filter
}

func (bm *TxModel) buildOptionsByQuery(query TxQuery) *options.FindOptions {
	findOps := options.Find()
	min := query.PageIndex * query.PageSize
	max := min + query.PageSize
	findOps.Min, findOps.Max = min, max
	return findOps
}
