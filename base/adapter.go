package base

import (
	"github.com/depocket/go-client"
	"math/big"
)

type ProjectBalance struct {
	Chain     string        `json:"chain"`
	Name      string        `json:"name"`
	LogoUrl   string        `json:"logo_url"`
	TagIds    []string      `json:"tags_id"`
	Balances  []interface{} `json:"balances"`
	UpdatedAt int64         `json:"updated_at,omitempty"`
	Version   int           `json:"version"`
}

type BalanceGroup struct {
	Type  string `json:"type"`
	Label string `json:"label"`
}

type Balance struct {
	Chain  string         `json:"chain"`
	Groups []BalanceGroup `json:"groups"`
	Pools  []Pool         `json:"pools"`
}

type Pool struct {
	Name         string         `json:"name"`
	Apy          *big.Float     `json:"apy"`
	RewardTokens []TokenBalance `json:"reward_tokens"`
	SupplyTokens []TokenBalance `json:"supply_tokens"`
	BorrowTokens []TokenBalance `json:"borrow_tokens"`
}

type TokenBalance struct {
	Address  string    `json:"address"`
	Balance  big.Float `json:"balance"`
	Symbol   string    `json:"symbol"`
	IconUrl  string    `json:"icon_url"`
	Price    float64   `json:"price"`
	Chain    string    `json:"chain"`
	Decimals uint64    `json:"decimals"`
	Name     string    `json:"name"`
}

type Manifest struct {
	Code              string            `json:"code"`
	Name              string            `json:"name"`
	Tags              []string          `json:"tags"`
	Description       string            `json:"description"`
	Url               string            `json:"url"`
	Links             map[string]string `json:"links"`
	SupportedNetworks []string          `json:"supported_networks"`
	Groups            []BalanceGroup    `json:"groups"`
}

type Adapter interface {
	ApplicationManifest() Manifest
	InjectDependencies(chain string, client *client.Client) error
	Balances(chain string, projectCode string, ad string) (ProjectBalance, error)
}
