package cosmos

import (
	"context"
	"fmt"
	"github.com/angelorc/cosmos-tracker/client/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

type QueryClient struct {
	Bank         banktypes.QueryClient
	Distribution distrtypes.QueryClient
	Staking      stakingtypes.QueryClient
}

func (qc *QueryClient) GetBalances(addr string) *types.Balances {
	balances := types.NewBalances()
	balances.Available, _ = qc.GetAvailableBalances(addr)
	delegations, _ := qc.GetDelegations(addr)

	for _, del := range delegations {
		balances.Delegations = balances.Delegations.Add(del.Balance)

		rewards, _ := qc.GetDelegatorRewards(addr, del.Delegation.ValidatorAddress)
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

func (qc *QueryClient) GetAvailableBalances(addr string) (sdk.Coins, error) {
	res, err := qc.Bank.AllBalances(
		context.Background(),
		&banktypes.QueryAllBalancesRequest{
			Address: addr,
		},
	)
	if err != nil {
		return sdk.Coins{}, fmt.Errorf("GetAvailableBalances err: %w", err)
	}

	fmt.Println(res.GetBalances())

	return res.GetBalances(), nil
}

func (qc *QueryClient) GetDelegations(addr string) ([]stakingtypes.DelegationResponse, error) {
	res, err := qc.Staking.DelegatorDelegations(
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

func (qc *QueryClient) GetDelegatorRewards(addr, valAddr string) (sdk.DecCoins, error) {
	res, err := qc.Distribution.DelegationRewards(
		context.Background(),
		&distrtypes.QueryDelegationRewardsRequest{
			DelegatorAddress: addr,
			ValidatorAddress: valAddr,
		},
	)
	if err != nil {
		fmt.Errorf("distr - delegatorvalidators err: %w", err)
	}

	return res.Rewards, nil
}
