package service

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	log "github.com/sirupsen/logrus"
	tmTypes "github.com/tendermint/tendermint/types"
	"mdu/explorer/common/model"
	"mdu/explorer/data-syncer/mdu/rpc"
	"mdu/explorer/data-syncer/util"
)

const TxStatusOK = "Success"

func InsertTxs(timestamp, height int64, tmTxs tmTypes.Txs) error {
	for _, tmTx := range tmTxs {
		var stdTx types.StdTx
		err := util.Cdc.UnmarshalBinaryLengthPrefixed(tmTx, &stdTx)
		if nil != err {
			log.Errorln(err)
			return err
		}

		tx := model.Tx{}
		tx.From = stdTx.GetSigners()[0].String() //TODO one tx have many signers??
		tx.Timestamp = timestamp
		tx.BlockHeight = height
		tx.Memo = stdTx.Memo
		hash := tmTx.Hash()
		tx.TxHash = fmt.Sprintf("%X", hash)
		tx.Fee = stdTx.Fee.Amount.String()

		txResult, err := rpc.TMTx(hash)
		if nil != err {
			log.Errorln(err)
			return err
		}

		if 0 != txResult.TxResult.Code {
			tx.Status = TxStatusOK
		} else {
			tx.Status = sdk.CodeToDefaultMsg(sdk.CodeType(txResult.TxResult.Code))
		}
		tx.UsedGas = txResult.TxResult.GasUsed
		tx.Messages = GetMsgs(&stdTx)

		model.NewTxModel().Insert(&tx)

		//TODO
		UpsertAddress(tx.From)

		//InsertMsgs(tx.TxHash, &stdTx)
	}

	return nil
}
