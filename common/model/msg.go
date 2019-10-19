package model
//
//import (
//	"context"
//	"encoding/json"
//	"go.mongodb.org/mongo-driver/mongo"
//	log "github.com/sirupsen/logrus"
//	"mdu/explorer/common/database"
//	"mdu/explorer/data-syncer/config"
//	"time"
//)
//
//const (
//	DB_COLLECTION_MESSAGES = "messages"
//)
//
//func NewMsgModel() *MsgModel {
//	mm := &MsgModel{database.DefaultDB.Client.Database(config.DefaultConfig.DBName).Collection(DB_COLLECTION_MESSAGES)}
//	return mm
//}
//
//type MsgModel struct {
//	collection *mongo.Collection
//}
//
//func (mm *MsgModel) InsertMsg(hash, typ string, doc interface{}) error {
//	bin, err := json.Marshal(doc)
//	if nil != err {
//		log.Errorln(err)
//		return err
//	}
//
//	tmpMap := map[string]interface{}{}
//	err = json.Unmarshal(bin, &tmpMap)
//	if nil != err {
//		log.Errorln(err)
//		return err
//	}
//
//	tmpMap["tx_hash"] = hash
//	tmpMap["tx_type"] = typ
//	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
//	_, err = mm.collection.InsertOne(ctx, tmpMap)
//	if nil != err {
//		log.Errorln(err)
//		return err
//	}
//
//	log.Debugf("tmMap:%+v", tmpMap)
//
//	return nil
//
//}
