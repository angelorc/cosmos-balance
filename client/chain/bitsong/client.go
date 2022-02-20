package bitsong

import (
	"fmt"
	"github.com/angelorc/cosmos-tracker/client/chain"
	"github.com/angelorc/cosmos-tracker/client/chain/cosmos"
	"github.com/angelorc/cosmos-tracker/client/codec"
	"github.com/angelorc/cosmos-tracker/client/query"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/bank"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/distribution"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	"github.com/cosmos/cosmos-sdk/x/staking"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"google.golang.org/grpc"
)

var ModuleBasics = []module.AppModuleBasic{
	bank.AppModuleBasic{},
	distribution.AppModuleBasic{},
	staking.AppModuleBasic{},
}

func NewClient(grpcURL string, denom string) (*chain.ChainClient, error) {
	grpcConn, err := grpc.Dial(
		grpcURL,             // your gRPC server address.
		grpc.WithInsecure(), // The SDK doesn't support any transport security mechanism.
	)
	if err != nil {
		return &chain.ChainClient{}, fmt.Errorf("failed to connect GRPC client: %s", err)
	}

	return &chain.ChainClient{
		ClientConn: grpcConn,
		Denom:      denom,
		Codec:      codec.MakeCodec(ModuleBasics),
		Query: query.QueryClient{
			Cosmos: cosmos.QueryClient{
				Bank:         banktypes.NewQueryClient(grpcConn),
				Distribution: distrtypes.NewQueryClient(grpcConn),
				Staking:      stakingtypes.NewQueryClient(grpcConn),
			},
		},
	}, nil
}
