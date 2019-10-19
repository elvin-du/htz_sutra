package util

import (
	"encoding/base64"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	distr "github.com/cosmos/cosmos-sdk/x/distribution"
	"github.com/cosmos/cosmos-sdk/x/genaccounts"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	"github.com/cosmos/cosmos-sdk/x/gov"
	"github.com/cosmos/cosmos-sdk/x/mint"
	"github.com/cosmos/cosmos-sdk/x/params"
	paramsclient "github.com/cosmos/cosmos-sdk/x/params/client"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	"github.com/cosmos/cosmos-sdk/x/staking"
	"github.com/cosmos/cosmos-sdk/x/supply"
	log "github.com/sirupsen/logrus"
)

var (
	// The module BasicManager is in charge of setting up basic,
	// non-dependant module elements, such as codec registration
	// and genesis verification.
	ModuleBasics = module.NewBasicManager(
		genaccounts.AppModuleBasic{},
		genutil.AppModuleBasic{},
		auth.AppModuleBasic{},
		bank.AppModuleBasic{},
		staking.AppModuleBasic{},
		mint.AppModuleBasic{},
		distr.AppModuleBasic{},
		gov.NewAppModuleBasic(paramsclient.ProposalHandler, distr.ProposalHandler),
		params.AppModuleBasic{},
		crisis.AppModuleBasic{},
		slashing.AppModuleBasic{},
		supply.AppModuleBasic{},
	)

	Cdc *codec.Codec
)

func init() {
	Cdc = MakeCodec()
}

// custom tx codec
func MakeCodec() *codec.Codec {
	var cdc = codec.New()

	ModuleBasics.RegisterCodec(cdc)
	sdk.RegisterCodec(cdc)
	codec.RegisterCrypto(cdc)
	codec.RegisterEvidences(cdc)

	return cdc
}

func Decode(data string) (*types.StdTx, error) {
	//var str = "tAEoKBapCjyoo2GaChTWyOydo2o0L7ppCCTIF955OWDkARIUHhmYb9UNfgNv/09+Uc9SZweiGa4aCgoEYXRvbRICMTASBBDAmgwaagom61rphyEC0gcbk7DutS3eI/0TZXPHcI7zHlH9e2ZMgcKlrLoaTKwSQIs3v6JLWPfnP6qnxOf4rQG5+aypdde3SBrxgnoptl8mIc7B5eR70/bTLV1dgDmpUT3hCQIScN7vMtehBYp+HUk="
	bin, err := base64.StdEncoding.DecodeString(data)
	if nil != err {
		log.Errorln(err)
		return nil, err
	}

	var tx types.StdTx
	err = Cdc.UnmarshalBinaryLengthPrefixed(bin, &tx)
	if nil != err {
		log.Errorln(err)
		return nil, err
	}

	return &tx, nil
	//log.Printf("%+v", tx)
	//for _, m := range tx.GetMsgs() {
	//	log.Printf("type:%s", m.Type())
	//}
}
