package chain

import (
	"fmt"
	"github.com/angelorc/cosmos-tracker/client/chain/cosmos"
	"github.com/angelorc/cosmos-tracker/client/chain/osmosis"
	"github.com/angelorc/cosmos-tracker/client/codec"
	"github.com/angelorc/cosmos-tracker/client/query"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	gammtypes "github.com/osmosis-labs/osmosis/x/gamm/types"
	"google.golang.org/grpc"
)

type ChainClient struct {
	*grpc.ClientConn
	Query query.QueryClient
	Codec codec.Codec
	Denom string `json:"denom"`
}

type Chains struct {
	Bitsong *ChainClient
	Osmosis *ChainClient
}

func NewClient(grpcURL string, denom string) (*ChainClient, error) {
	chainClient := &ChainClient{}

	grpcConn, err := grpc.Dial(
		grpcURL,             // your gRPC server address.
		grpc.WithInsecure(), // The SDK doesn't support any transport security mechanism.
	)
	if err != nil {
		return chainClient, fmt.Errorf("failed to connect GRPC client: %s", err)
	}

	return &ChainClient{
		ClientConn: grpcConn,
		Denom:      denom,
		Codec:      codec.MakeCodecConfig(),
		Query: query.QueryClient{
			Cosmos: cosmos.QueryClient{
				Bank:         banktypes.NewQueryClient(grpcConn),
				Distribution: distrtypes.NewQueryClient(grpcConn),
				Staking:      stakingtypes.NewQueryClient(grpcConn),
			},
			Osmosis: osmosis.QueryClient{
				Gamm: gammtypes.NewQueryClient(grpcConn),
			},
		},
	}, nil
}
