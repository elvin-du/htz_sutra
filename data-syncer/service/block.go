package service

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	distrTypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	log "github.com/sirupsen/logrus"
	"mdu/explorer/common/model"
	"mdu/explorer/data-syncer/mdu/rpc"
)

func UpdateBlocks() error {
	tmLatestBlock, err := rpc.TMLastestBlock()
	if nil != err {
		log.Errorln(err)
		return err
	}

	latestBlock, err := model.NewBlocksModel().FindLatestBlock()
	if nil != err {
		log.Errorln(err)
		return err
	}

	diff := tmLatestBlock.Block.Height - latestBlock.Height

	for i := int64(0); i < diff; i++ {
		h := latestBlock.Height + i + 1
		InsertBlocks(h)

		//TODO
		//blockchain system donot save too old data of validator info
		if (tmLatestBlock.Block.Height - h) < 99 {
			UpsertValidators(h)
		}
	}

	return nil
}

func InsertBlocks(height int64) error {
	block, err := rpc.TMBlock(&height)
	if nil != err {
		log.Errorln(err)
		return err
	}
	log.Debugf("block:%+v", block)

	b := model.Block{}
	b.Height = block.BlockMeta.Header.Height
	b.Proposer = block.BlockMeta.Header.ProposerAddress.String()
	b.BlockHash = block.BlockMeta.BlockID.Hash.String()
	b.Timestamp = block.BlockMeta.Header.Time.Unix()
	b.Transactions = block.BlockMeta.Header.NumTxs

	commit, err := rpc.TMCommit(&block.BlockMeta.Header.Height)
	if nil != err {
		log.Errorln(err)
		return err
	}
	b.Validators = len(commit.Commit.Precommits)

	b.Rewards, err = Rewards(height)
	if nil != err {
		log.Errorln(err)
		return err
	}

	model.NewBlocksModel().InsertBlock(&b)

	InsertTxs(b.Timestamp, b.Height, block.Block.Txs)

	return nil
}

//block rewards
func Rewards(height int64) (string, error) {
	bResult, err := rpc.TMBlockResults(&height)
	if nil != err {
		panic(err)
	}

	//TODO stake denom will be change to mdu??
	rewards := sdk.NewDecCoin("umdu", sdk.NewInt(0))

	evts := bResult.Results.BeginBlock.Events
	for _, e := range evts {
		if distrTypes.EventTypeRewards == e.Type ||
			e.Type == distrTypes.EventTypeProposerReward {

			for _, v2 := range e.Attributes {
				if string(v2.Key) == sdk.AttributeKeyAmount {
					decCoin, err := sdk.ParseDecCoin(string(v2.Value))
					if nil != err {
						log.Errorln(err)
						return "", err
					}

					rewards = rewards.Add(decCoin)
				}
			}
		}
	}

	return rewards.String(), nil
}
