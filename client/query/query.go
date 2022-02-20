package query

import (
	"github.com/angelorc/cosmos-tracker/client/chain/cosmos"
	"github.com/angelorc/cosmos-tracker/client/chain/osmosis"
)

type QueryClient struct {
	Cosmos  cosmos.QueryClient
	Osmosis osmosis.QueryClient
}
