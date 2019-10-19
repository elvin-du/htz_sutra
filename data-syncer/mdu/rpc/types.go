package rpc

type Commission struct {
	Rate          string `json:"rate"`
	MaxRate       string `json:"max_rate"`
	MaxChangeRate string `json:"max_change_rate"`
	UpdateTime    string `json:"update_time"`
}

type Description struct {
	Website  string `json:"website"`
	Details  string `json:"details"`
	Moniker  string `json:"moniker"`
	Identity string `json:"identity"`
}

type StakingValidator struct {
	Commission        *Commission  `json:"commission"`
	Description       *Description `json:"description"`
	ConsensusPubkey   string       `json:"consensus_pubkey"`
	Jailed            bool         `json:"jailed"`
	Status            int          `json:"status"`
	Tokens            string       `json:"tokens"`
	DelegatorShares   string       `json:"delegator_shares"`
	UnbondingHeight   string       `json:"unbonding_height"`
	UnbondingTime     string       `json:"unbonding_time"`
	OperatorAddress   string       `json:"operator_address"`
	MinSelfDelegation string       `json:"min_self_delegation"`
}

type StakingValidators struct {
	Height string              `json:"height"`
	Result []*StakingValidator `json:"result"`
}

type Txs struct {
	TotalCount uint64 `json:"total_count"`
	Count      uint64 `json:"count"`
	PageNumber uint64 `json:"page_number"`
	PageTotal  uint64 `json:"page_total"`
	Limit      uint64 `json:"limit"`
	//Txs        interface{}   `json:"txs"`
}
