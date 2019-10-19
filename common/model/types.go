package model

type Block struct {
	Height       int64  `json:"height"  bson:"height"`
	Transactions int64  `json:"transactions"  bson:"transactions"`
	Timestamp    int64  `json:"timestamp" bson:"timestamp"`
	BlockHash    string `json:"block_hash" bson:"block_hash"`
	Proposer     string `json:"proposer" bson:"proposer"`
	Validators   int    `json:"validators"`
	Rewards      string `json:"rewards"`
}

type Tx struct {
	TxHash      string                   `json:"tx_hash" bson:"tx_hash"`
	Timestamp   int64                    `json:"timestamp" bson:"timestamp"`
	Status      string                   `json:"status"`
	BlockHeight int64                    `json:"block_height" bson:"block_height"`
	Fee         string                   `json:"fee"`
	UsedGas     int64                    `json:"used_gas" bson:"used_gas"`
	Memo        string                   `json:"memo"`
	From        string                   `json:"from"`
	Messages    []map[string]interface{} `json:"messages"`
}

type Validator struct {
	Moniker         string `json:"moniker"`
	OperatorAddress string `json:"operator_address" bson:"operator_address"`
	OwnerAddress    string `json:"owner_address" bson:"owner_address"`
	RewardAddress   string `json:"reward_address" bson:"reward_address"`
	Details         string `json:"details"`
	VotingPower     string `json:"voting_power" bson:"voting_power"`
	Website         string `json:"website"`
}

type Address struct {
	Address        string `json:"address" bson:"address"`
	Balance        string `json:"balance" bson:"balance"`
	Delegated      string `json:"delegated" bson:"delegated"`
	RewardAddress  string `json:"reward_address" bson:"reward_address"`
	TransactionNum uint64 `json:"transaction_num" bson:"transaction_num"`
}
