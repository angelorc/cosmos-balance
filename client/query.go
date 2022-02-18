package client

import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

type Balances struct {
	Available   sdk.Coins    `json:"available"`
	Delegations sdk.Coin     `json:"delegations"`
	Rewards     sdk.DecCoins `json:"rewards"`
	Totals      sdk.Coins    `json:"totals"`
}

func NewBalances(denom string) *Balances {
	return &Balances{
		Delegations: sdk.Coin{
			Denom:  denom,
			Amount: sdk.NewInt(0),
		},
	}
}

func (c *ChainClient) GetAvailableBalances(addr string) (sdk.Coins, error) {
	res, err := c.Query.Bank.AllBalances(
		context.Background(),
		&banktypes.QueryAllBalancesRequest{
			Address: addr,
		},
	)
	if err != nil {
		return sdk.Coins{}, fmt.Errorf("GetAvailableBalances err: %w", err)
	}

	return res.GetBalances(), nil
}

func (c *ChainClient) GetBalances(addr string) *Balances {
	balances := NewBalances(c.Denom)
	balances.Available, _ = c.GetAvailableBalances(addr)
	delegations, _ := c.GetDelegations(addr)

	for _, del := range delegations {
		balances.Delegations = balances.Delegations.Add(del.Balance)

		rewards, _ := c.GetDelegatorRewards(addr, del.Delegation.ValidatorAddress)
		for _, reward := range rewards {
			balances.Rewards = balances.Rewards.Add(reward)
		}
	}

	for _, coin := range balances.Available {
		balances.Totals = balances.Totals.Add(coin)
	}

	balances.Totals = balances.Totals.Add(balances.Delegations)

	for _, deccoin := range balances.Rewards {
		coin, _ := deccoin.TruncateDecimal()
		balances.Totals = balances.Totals.Add(coin)
	}

	return balances
}

func (c *ChainClient) GetDelegations(addr string) ([]stakingtypes.DelegationResponse, error) {
	res, err := c.Query.Staking.DelegatorDelegations(
		context.Background(),
		&stakingtypes.QueryDelegatorDelegationsRequest{
			DelegatorAddr: addr,
		},
	)
	if err != nil {
		fmt.Errorf("GetDelegations err: %w", err)
	}

	return res.DelegationResponses, nil
}

func (c *ChainClient) GetDelegatorRewards(addr, valAddr string) (sdk.DecCoins, error) {
	res, err := c.Query.Distribution.DelegationRewards(
		context.Background(),
		&distrtypes.QueryDelegationRewardsRequest{
			DelegatorAddress: addr,
			ValidatorAddress: valAddr,
		},
	)
	if err != nil {
		fmt.Errorf("distr - delegatorvalidators err: %w", err)
	}

	return res.Rewards, err
}
