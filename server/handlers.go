package server

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"net/http"
)

func (s *Server) registerRoutes() {
	s.GET("/bitsong/:address", s.GetBitsongBalances)
	s.GET("/osmosis/:address", s.GetOsmosisBalances)
	s.GET("/swagger/*", echoSwagger.WrapHandler)
}

type BalancesResponse struct {
	Totals      map[string]interface{} `json:"totals"`
	Available   map[string]interface{} `json:"available"`
	Delegations map[string]interface{} `json:"delegations"`
	Rewards     map[string]interface{} `json:"rewards"`
}

// GetBitsongBalances godoc
// @Summary Get bitsong balances by address.
// @Description Get bitsong balances by address.
// @Tags bitsong
// @Accept */*
// @Produce json
// @Success 200 {object} BalancesResponse
// @Param address path string true "Bitsong address to query"
// @Router /bitsong/{address} [get]
func (s *Server) GetBitsongBalances(c echo.Context) error {
	addr := c.Param("address")
	balances := s.chains.Bitsong.Query.Cosmos.GetBalances(addr)

	return c.JSON(http.StatusOK, balances)
}

// GetOsmosisBalances godoc
// @Summary Get osmosis balances by address.
// @Description Get osmosis balances by address.
// @Tags osmosis
// @Accept */*
// @Produce json
// @Success 200 {object} BalancesResponse
// @Param address path string true "Osmosis address to query"
// @Router /osmosis/{address} [get]
func (s *Server) GetOsmosisBalances(c echo.Context) error {
	addr := c.Param("address")
	balances := s.chains.Osmosis.Query.Cosmos.GetBalances(addr)

	return c.JSON(http.StatusOK, balances)
}
