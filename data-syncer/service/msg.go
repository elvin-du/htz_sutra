package service

import (
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	"github.com/cosmos/cosmos-sdk/x/distribution"
	"github.com/cosmos/cosmos-sdk/x/gov"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	"github.com/cosmos/cosmos-sdk/x/staking"
)

func GetMsgs(stdTx *types.StdTx) []map[string]interface{} {
	ret := make([]map[string]interface{}, 0, len(stdTx.Msgs))

	insertMsg := func(typ string, data interface{}) {
		tmp := map[string]interface{}{}
		tmp["typ"] = typ
		tmp["data"] = data
		ret = append(ret, tmp)
	}

	for _, stdMsg := range stdTx.Msgs {
		switch msg := stdMsg.(type) {
		case bank.MsgSend:
			insertMsg(msg.Type(), &msg)
		case bank.MsgMultiSend:
			insertMsg(msg.Type(), &msg)

		case crisis.MsgVerifyInvariant:
			insertMsg(msg.Type(), &msg)

		case distribution.MsgSetWithdrawAddress:
			insertMsg(msg.Type(), &msg)

		case distribution.MsgWithdrawDelegatorReward:
			insertMsg(msg.Type(), &msg)

		case distribution.MsgWithdrawValidatorCommission:
			insertMsg(msg.Type(), &msg)

		case gov.MsgSubmitProposal:
			insertMsg(msg.Type(), &msg)

		case gov.MsgDeposit:
			insertMsg(msg.Type(), &msg)

		case gov.MsgVote:
			insertMsg(msg.Type(), &msg)

		case slashing.MsgUnjail:
			insertMsg(msg.Type(), &msg)

		case staking.MsgCreateValidator:
			insertMsg(msg.Type(), &msg)

		case staking.MsgEditValidator:
			insertMsg(msg.Type(), &msg)

		case staking.MsgDelegate:
			insertMsg(msg.Type(), &msg)

		case staking.MsgBeginRedelegate:
			insertMsg(msg.Type(), &msg)

		case staking.MsgUndelegate:
			insertMsg(msg.Type(), &msg)

		default:
			panic("unkown msg type")
		}
	}

	return ret
}

//
//func InsertMsgs(hash string, stdTx *types.StdTx) {
//	for _, stdMsg := range stdTx.Msgs {
//		switch msg := stdMsg.(type) {
//		case bank.MsgSend:
//			model.NewMsgModel().InsertMsg(hash, "send", &msg)
//
//		case bank.MsgMultiSend:
//			model.NewMsgModel().InsertMsg(hash, "multi_send", &msg)
//
//		case crisis.MsgVerifyInvariant:
//			model.NewMsgModel().InsertMsg(hash, "verify_invariant", &msg)
//
//		case distribution.MsgSetWithdrawAddress:
//			model.NewMsgModel().InsertMsg(hash, "set_withdraw_address", &msg)
//
//		case distribution.MsgWithdrawDelegatorReward:
//			model.NewMsgModel().InsertMsg(hash, "withdraw_delegator_reward", &msg)
//
//		case distribution.MsgWithdrawValidatorCommission:
//			model.NewMsgModel().InsertMsg(hash, "withdraw_validator_commission", &msg)
//
//		case gov.MsgSubmitProposal:
//			model.NewMsgModel().InsertMsg(hash, "submit_proposal", &msg)
//
//		case gov.MsgDeposit:
//			model.NewMsgModel().InsertMsg(hash, "deposit", &msg)
//
//		case gov.MsgVote:
//			model.NewMsgModel().InsertMsg(hash, "vote", &msg)
//
//		case slashing.MsgUnjail:
//			model.NewMsgModel().InsertMsg(hash, "unjail", &msg)
//
//		case staking.MsgCreateValidator:
//			model.NewMsgModel().InsertMsg(hash, "create_validator", &msg)
//
//		case staking.MsgEditValidator:
//			model.NewMsgModel().InsertMsg(hash, "edit_validator", &msg)
//
//		case staking.MsgDelegate:
//			model.NewMsgModel().InsertMsg(hash, "delegate", &msg)
//
//		case staking.MsgBeginRedelegate:
//			model.NewMsgModel().InsertMsg(hash, "begin_redelegate", &msg)
//
//		case staking.MsgUndelegate:
//			model.NewMsgModel().InsertMsg(hash, "begin_unbonding", &msg)
//
//		default:
//			panic("unkown msg type")
//		}
//	}
//}
