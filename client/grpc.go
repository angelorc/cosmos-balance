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

type Client struct {
	*grpc.ClientConn
	Query QueryClient
}

func NewClient(grpcURL string) (*Client, error) {
	grpcConn, err := grpc.Dial(
		grpcURL,             // your gRPC server address.
		grpc.WithInsecure(), // The SDK doesn't support any transport security mechanism.
	)
	if err != nil {
		return &Client{}, fmt.Errorf("failed to connect GRPC client: %s", err)
	}

	return &Client{
		ClientConn: grpcConn,
		Query: QueryClient{
			Bank:         banktypes.NewQueryClient(grpcConn),
			Distribution: distrtypes.NewQueryClient(grpcConn),
			Staking:      stakingtypes.NewQueryClient(grpcConn),
		},
	}, nil
}
