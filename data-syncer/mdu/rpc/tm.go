package rpc

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/tendermint/tendermint/rpc/client"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
	"mdu/explorer/data-syncer/config"
)

func TMLastestBlock() (*ctypes.ResultBlock, error) {
	return TMBlock(nil)
}

//get lastest block while height is nil
func TMBlock(height *int64) (*ctypes.ResultBlock, error) {
	cli := client.NewHTTP(fmt.Sprintf("tcp://%s", config.DefaultConfig.NodeAddress), "/websocket")
	err := cli.Start()
	if err != nil {
		log.Errorln(err)
		return nil, err
	}
	defer cli.Stop()

	blockInfo, err := cli.Block(height)
	if err != nil {
		log.Errorln(err)
		return nil, err
	}

	return blockInfo, nil
}

func TMStatus() (*ctypes.ResultStatus, error) {
	cli := client.NewHTTP(fmt.Sprintf("tcp://%s", config.DefaultConfig.NodeAddress), "/websocket")
	err := cli.Start()
	if err != nil {
		log.Errorln(err)
		return nil, err
	}
	defer cli.Stop()

	info, err := cli.Status()
	if err != nil {
		log.Errorln(err)
		return nil, err
	}

	return info, nil
}

func TMLastestValidators() (*ctypes.ResultValidators, error) {
	return TMValidators(nil)
}

//get lastest validators while height is nil
func TMValidators(height *int64) (*ctypes.ResultValidators, error) {
	cli := client.NewHTTP(fmt.Sprintf("tcp://%s", config.DefaultConfig.NodeAddress), "/websocket")
	err := cli.Start()
	if err != nil {
		log.Errorln(err)
		return nil, err
	}
	defer cli.Stop()

	info, err := cli.Validators(height)
	if err != nil {
		log.Errorln(err)
		return nil, err
	}

	return info, nil
}

func TMLastestBlockResults() (*ctypes.ResultBlockResults, error) {
	return TMBlockResults(nil)
}
func TMBlockResults(height *int64) (*ctypes.ResultBlockResults, error) {
	cli := client.NewHTTP(fmt.Sprintf("tcp://%s", config.DefaultConfig.NodeAddress), "/websocket")
	err := cli.Start()
	if err != nil {
		log.Errorln(err)
		return nil, err
	}
	defer cli.Stop()

	info, err := cli.BlockResults(height)
	if err != nil {
		log.Errorln(err)
		return nil, err
	}

	return info, nil
}

func TMLastestBlockCommit() (*ctypes.ResultCommit, error) {
	return TMCommit(nil)
}

func TMCommit(height *int64) (*ctypes.ResultCommit, error) {
	cli := client.NewHTTP(fmt.Sprintf("tcp://%s", config.DefaultConfig.NodeAddress), "/websocket")
	err := cli.Start()
	if err != nil {
		log.Errorln(err)
		return nil, err
	}
	defer cli.Stop()

	info, err := cli.Commit(height)
	if err != nil {
		log.Errorln(err)
		return nil, err
	}

	return info, nil
}

func TMTx(hash []byte) (*ctypes.ResultTx, error) {
	cli := client.NewHTTP(fmt.Sprintf("tcp://%s", config.DefaultConfig.NodeAddress), "/websocket")
	err := cli.Start()
	if err != nil {
		log.Errorln(err)
		return nil, err
	}
	defer cli.Stop()

	info, err := cli.Tx(hash, false)
	if err != nil {
		log.Errorln(err)
		return nil, err
	}

	return info, nil
}
