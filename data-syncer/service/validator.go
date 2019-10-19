package service

import (
	log "github.com/sirupsen/logrus"
	"github.com/tendermint/tendermint/libs/bech32"
	"mdu/explorer/common/model"
	"mdu/explorer/data-syncer/mdu/rpc"
)

func UpsertValidators(height int64) error {
	stakingVals, err := rpc.MDUStakingValidators("", height)
	if nil != err {
		log.Errorln(err)
		return err
	}

	for _, sval := range stakingVals.Result {
		val := &model.Validator{}
		val.Details = sval.Description.Details
		val.Moniker = sval.Description.Moniker
		val.Website = sval.Description.Website
		val.VotingPower = sval.DelegatorShares //TODO check again?
		val.OperatorAddress = sval.OperatorAddress
		val.OwnerAddress = sval.OperatorAddress

		_, bz, err := bech32.DecodeAndConvert(val.OperatorAddress)
		if nil != err {
			log.Errorln(err)
			return err
		}

		addr, err := bech32.ConvertAndEncode("mdu", bz)
		if nil != err {
			log.Errorln(err)
			return err
		}

		rewardAddr, err := rpc.MDUDelegatorWithdrawAddress(addr)
		if nil != err {
			log.Errorln(err)
			return err
		}
		//val.RewardAddress = sval.OperatorAddress
		val.RewardAddress = rewardAddr

		model.NewValidatorModel().Upsert(val)
	}

	return nil
}
