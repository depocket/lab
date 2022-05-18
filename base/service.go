package base

import (
	"github.com/depocket/lab/sdk"
	"math/big"
)

type BalanceGroup struct {
	Type  string `json:"type"`
	Label string `json:"label"`
}

type Balance struct {
	Group BalanceGroup `json:"group"`
	Pools []Pool       `json:"pools"`
}

type Pool struct {
	Name         string         `json:"name"`
	Apy          *big.Float     `json:"apy"`
	RewardTokens []TokenBalance `json:"reward_tokens"`
	SupplyTokens []TokenBalance `json:"supply_tokens"`
	BorrowTokens []TokenBalance `json:"borrow_tokens"`
}

type TokenBalance struct {
	Address  string     `json:"address"`
	Balance  *big.Float `json:"balance"`
	Symbol   string     `json:"symbol"`
	IconUrl  string     `json:"icon_url"`
	Price    float64    `json:"price"`
	Chain    string     `json:"chain"`
	Decimals uint64     `json:"decimals"`
	Name     string     `json:"name"`
}

type AppManifest struct {
	Code              string         `json:"code"`
	Name              string         `json:"name"`
	Tags              []string       `json:"tags"`
	Description       string         `json:"description"`
	Url               string         `json:"url"`
	Links             []string       `json:"links"`
	SupportedNetworks []string       `json:"supported_networks"`
	Groups            []BalanceGroup `json:"groups"`
}

type AppService interface {
	Code() string
	Name() string
	Tags() []string
	Description() string
	Url() string
	Links() map[string]string
	SupportedNetworks() []string
	Groups() []BalanceGroup

	BeforeFetchData(client *sdk.Client) error
	FetchBalances() ([]Balance, error)
	FetchStakingBalance() (Balance, error)
	FetchFarmingBalance() (Balance, error)
	FetchLendingBalance() (Balance, error)
	FetchBorrowBalance() (Balance, error)
}
