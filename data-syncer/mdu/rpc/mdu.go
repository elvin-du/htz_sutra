package rpc

import (
	"errors"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	log "github.com/sirupsen/logrus"
	"mdu/explorer/data-syncer/config"
	"mdu/explorer/data-syncer/util"
)

//The validator bond status. Must be either 'bonded’, 'unbonded’, or 'unbonding’.
func MDUStakingValidators(status string, height int64) (*StakingValidators, error) {
	url := fmt.Sprintf("http://%s/staking/validators?height=%d", config.DefaultConfig.LcdAddress, height)
	if "" != status {
		url += fmt.Sprintf("&status=%s", status)
	}
	log.Debugln("url", url)

	sVal := &StakingValidators{}
	err := util.HTTPGetJson(url, sVal)
	if nil != err {
		log.Errorln(err)
		return nil, err
	}

	return sVal, nil
}

func MDUStakingValidator(addr string) (*StakingValidator, error) {
	url := fmt.Sprintf("http://%s/staking/validators/%s", config.DefaultConfig.LcdAddress, addr)
	log.Debugln("url", url)

	sVal := &StakingValidator{}
	err := util.HTTPGetJson(url, sVal)
	var httpErr util.HTTPError
	if nil != err {
		if errors.As(err, &httpErr) {
			log.Infoln(httpErr.Error())
			return sVal, nil
		}

		log.Errorln(err)
		return nil, err
	}

	return sVal, nil
}

func MDUDelegatorWithdrawAddress(delegatorAddr string) (string, error) {
	url := fmt.Sprintf("http://%s/distribution/delegators/%s/withdraw_address", config.DefaultConfig.LcdAddress, delegatorAddr)
	log.Debugln("url", url)

	m := map[string]string{}
	err := util.HTTPGetJson(url, &m)
	if nil != err {
		log.Errorln(err)
		return "", err
	}

	return m["result"], nil
}

func MDUTxsByMsgSender(from string) (*Txs, error) {
	url := fmt.Sprintf("http://%s/txs?message.sender=%d", config.DefaultConfig.LcdAddress, from)
	log.Debugln("url", url)

	txs := &Txs{}
	err := util.HTTPGetJson(url, txs)
	if nil != err {
		log.Errorln(err)
		return nil, err
	}

	return txs, nil
}

func MDUBalance(addr string) (*sdk.Coins, error) {
	url := fmt.Sprintf("http://%s/bank/balances/%s", config.DefaultConfig.LcdAddress, addr)
	log.Debugln("url", url)

	coins := &sdk.Coins{}
	err := util.HTTPGetJson(url, coins)
	if nil != err {
		log.Errorln(err)
		return nil, err
	}

	return coins, nil
}

func MDURewardAddress(addr string) (string, error) {
	url := fmt.Sprintf("http://%s/distribution/delegators/%s/withdraw_address", config.DefaultConfig.LcdAddress, addr)
	log.Debugln("url", url)

	ret, err := util.HTTPGetString(url)
	var httpErr util.HTTPError
	if nil != err {
		if errors.As(err, &httpErr) {
			log.Infoln(httpErr.Error())
			return "", nil
		}

		log.Errorln(err)
		return "", err
	}

	return ret, nil
}
