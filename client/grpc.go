package client

import (
	"fmt"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"google.golang.org/grpc"
)

type QueryClient struct {
	Bank         banktypes.QueryClient
	Distribution distrtypes.QueryClient
	Staking      stakingtypes.QueryClient
}

type ChainClient struct {
	*grpc.ClientConn
	Query QueryClient
	Denom string `json:"denom"`
}

type Chains struct {
	Bitsong *ChainClient
	Osmosis *ChainClient
}

func NewClient(grpcURL string, denom string) (*ChainClient, error) {
	grpcConn, err := grpc.Dial(
		grpcURL,             // your gRPC server address.
		grpc.WithInsecure(), // The SDK doesn't support any transport security mechanism.
	)
	if err != nil {
		return &ChainClient{}, fmt.Errorf("failed to connect GRPC client: %s", err)
	}

	return &ChainClient{
		ClientConn: grpcConn,
		Query: QueryClient{
			Bank:         banktypes.NewQueryClient(grpcConn),
			Distribution: distrtypes.NewQueryClient(grpcConn),
			Staking:      stakingtypes.NewQueryClient(grpcConn),
		},
		Denom: denom,
	}, nil
}
