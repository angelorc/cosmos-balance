package osmosis

import (
	"context"
	"fmt"
	osmogammtypes "github.com/osmosis-labs/osmosis/x/gamm/types"
)

type QueryClient struct {
	Gamm osmogammtypes.QueryClient
}

func (qc *QueryClient) GetPoolBalance(poolId uint64) error {
	_, err := qc.Gamm.Pool(
		context.Background(),
		&osmogammtypes.QueryPoolRequest{
			PoolId: poolId,
		},
	)
	if err != nil {
		fmt.Errorf("gamm - GetPoolBalance err: %w", err)
	}

	//pool := res.Pool.(osmogammtypes.PoolI)

	return nil
}
