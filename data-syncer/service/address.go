package service

import (
	log "github.com/sirupsen/logrus"
	"mdu/explorer/common/model"
	"mdu/explorer/data-syncer/mdu/rpc"
)

func UpsertAddress(addr string) {
	a := &model.Address{}
	a.Address = addr

	txs, err := rpc.MDUTxsByMsgSender(a.Address)
	if nil != err {
		log.Errorln(err)
		return
	}
	a.TransactionNum = txs.TotalCount

	coins, err := rpc.MDUBalance(a.Address)
	if nil != err {
		log.Errorln(err)
		return
	}
	a.Balance = coins.AmountOf("umdu").String() //TODO

	stakingValidator, err := rpc.MDUStakingValidator(a.Address)
	if nil != err {
		log.Errorln(err)
		return
	}
	a.Delegated = stakingValidator.DelegatorShares //TODO one share == one token??

	rewardAddr, err := rpc.MDURewardAddress(a.Address)
	if nil != err {
		log.Errorln(err)
		return
	}
	a.RewardAddress = rewardAddr

	model.NewAddressesModel().Upsert(a)
}
