package types

import sdk "github.com/cosmos/cosmos-sdk/types"

type Balances struct {
	Available   sdk.Coins    `json:"available"`
	Delegations sdk.Coin     `json:"delegations"`
	Rewards     sdk.DecCoins `json:"rewards"`
	Totals      sdk.Coins    `json:"totals"`
}

func NewBalances() *Balances {
	return &Balances{}
}
