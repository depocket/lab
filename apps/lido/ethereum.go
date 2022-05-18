package lido

import (
	"context"
	"fmt"
	"github.com/depocket/lab/base"
	"github.com/depocket/lab/sdk"
)

type EthereumService struct {
	Tokens []*sdk.Token
	Pools  []*sdk.Pool
}

func (s *EthereumService) Code() string {
	return "ethereum_lido"
}

func (s *EthereumService) Name() string {
	return "Lido"
}

func (s *EthereumService) Tags() []string {
	return []string{"staking"}
}

func (s *EthereumService) Description() string {
	return "Simplified and secure staking for digital assets."
}

func (s *EthereumService) Url() string {
	return "https://lido.fi/"
}

func (s *EthereumService) Links() map[string]string {
	return map[string]string{
		"site": "https://lido.fi/",
	}
}

func (s *EthereumService) BeforeFetchData(client *sdk.Client) error {
	ctx := context.Background()
	list, _, err := client.Tokens.List(ctx, &sdk.TokenListOptions{
		ListOptions: sdk.ListOptions{
			Chain: "ethereum",
		},
	})
	if err != nil {
		return err
	}
	s.Tokens = list
	pools, _, err := client.Pools.ListByProjectCode(ctx, s.Code(), &sdk.PoolListOptions{
		ListOptions: sdk.ListOptions{
			Chain: "ethereum",
		},
	})
	s.Pools = pools
	fmt.Println(s.Pools)
	return nil
}

func (s *EthereumService) SupportedNetworks() []string {
	return []string{
		"ethereum",
	}
}

func (s *EthereumService) Groups() []base.BalanceGroup {
	return []base.BalanceGroup{
		{
			Type:  "staking",
			Label: "Staking",
		},
	}
}

func (s *EthereumService) FetchBalances() ([]base.Balance, error) {
	return []base.Balance{}, nil
}

func (s *EthereumService) FetchStakingBalance() (base.Balance, error) {
	return base.Balance{}, nil
}

func (s *EthereumService) FetchFarmingBalance() (base.Balance, error) {
	return base.Balance{}, nil
}

func (s *EthereumService) FetchLendingBalance() (base.Balance, error) {
	return base.Balance{}, nil
}

func (s *EthereumService) FetchBorrowBalance() (base.Balance, error) {
	return base.Balance{}, nil
}
