package model

import (
	"context"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"mdu/explorer/common/database"
	"mdu/explorer/data-syncer/config"
	"time"
)

var (
	DB_COLLECTION_BLOCKS = "blocks"
)

func NewBlocksModel() *BlocksModel {
	bm := &BlocksModel{database.DefaultDB.Client.Database(config.DefaultConfig.DBName).Collection(DB_COLLECTION_BLOCKS)}
	return bm
}

type BlocksModel struct {
	collection *mongo.Collection
}

type BlockQuery struct {
	PageIndex int32
	PageSize  int32
	Proposer  string
}

func (bm *BlocksModel) InsertBlock(b *Block) error {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	res, err := bm.collection.InsertOne(ctx, b)
	if nil != err {
		log.Errorln(err)
		return err
	}

	log.Debugln("insertid:", res.InsertedID)
	return nil
}

func (bm *BlocksModel) FindLatestBlock() (*Block, error) {
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	var b Block
	err := bm.collection.FindOne(
		ctx,
		bsonx.Doc{},
		options.FindOne().SetSort(bsonx.Doc{{"height", bsonx.Int32(-1)}})).Decode(&b)
	if mongo.ErrNoDocuments == err {
		b.Height = 1 //TODO tmp
		return &b, nil
	} else if nil != err {
		log.Errorln(err)
		return nil, err
	}

	log.Debugf("FindLastestBlock:%+v", b)
	return &b, nil
}

func (bm *BlocksModel) FindBlockByHeight(height int64) (*Block, error) {
	return bm.findOneBlock(bson.M{"height": height})
}

func (bm *BlocksModel) FindBlockByHash(hash string) (*Block, error) {
	return bm.findOneBlock(bson.M{"block_hash": hash})
}

func (bm *BlocksModel) CountBlocks(query BlockQuery) (int64, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	filter := bm.buildFilterByQuery(query)
	totalCount, err := bm.collection.CountDocuments(ctx, filter)
	if err != nil {
		log.Errorln(err)
		return -1, err
	}
	return totalCount, err
}

func (bm *BlocksModel) List(query BlockQuery) ([]*Block, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	filter := bm.buildFilterByQuery(query)
	findOps := bm.buildOptionsByQuery(query)

	cur, err := bm.collection.Find(ctx, filter, findOps)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	results := make([]*Block, 20)
	err = cur.All(ctx, &results)
	if err != nil {
		log.Errorln(err)
		return nil, err
	}
	return results, nil
}

func (bm *BlocksModel) findOneBlock(filter interface{}) (*Block, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result := bm.collection.FindOne(ctx, filter)
	if result.Err() != nil {
		log.Errorln(result.Err())
		return nil, result.Err()
	}
	tx := &Block{}
	err := result.Decode(tx)
	if err != nil {
		log.Errorln(result.Err())
		return nil, err
	}
	return tx, nil
}

func (bm *BlocksModel) buildFilterByQuery(query BlockQuery) *bson.M {
	filter := bson.M{}
	if query.Proposer != "" {
		filter["proposer"] = query.Proposer
	}
	return &filter
}

func (bm *BlocksModel) buildOptionsByQuery(query BlockQuery) *options.FindOptions {
	findOps := options.Find()
	min := query.PageIndex * query.PageSize
	max := min + query.PageSize
	findOps.Min, findOps.Max = min, max
	return findOps
}
