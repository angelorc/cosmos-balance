package main

import (
	"encoding/json"
	"fmt"
	grpccli "github.com/angelorc/cosmos-balance/client"
	"github.com/angelorc/cosmos-balance/server"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"go.uber.org/zap"
	"log"
)

const (
	GRPCUrl     = "78.47.82.185:9090"
	BitsongAddr = "bitsong1uaw9csdydp4kpzxh5yphp5xzerwtdac8mng6g7"
)

func printJsonResponse(v interface{}) string {
	res, _ := json.MarshalIndent(v, "", " ")
	return string(res)
}

type Balances struct {
	Available sdk.Coins
}

func main() {
	log.Printf("connecting to grpc server...\n")
	client, err := grpccli.NewClient(GRPCUrl)
	if err != nil {
		fmt.Errorf("grpc conn error: %w", err)
	}
	log.Printf("grpc server connected...\n")
	defer client.Close()

	logger, _ := zap.NewProductionConfig().Build()
	defer logger.Sync()

	s := server.NewServer(client, logger)
	s.Start("127.0.0.1:8000")
}
