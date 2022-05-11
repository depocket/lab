package ethereum

import (
	"github.com/depocket/studio/base"
	"github.com/depocket/studio/sdk"
)

type Service struct {
	Client sdk.Client
}

func (s *Service) Code() string {
	return "ethereum_lido"
}

func (s *Service) Name() string {
	return "Lido"
}

func (s *Service) Tags() []string {
	return []string{"staking"}
}

func (s *Service) Description() string {
	return "Simplified and secure staking for digital assets."
}

func (s *Service) Url() string {
	return "https://lido.fi/"
}

func (s *Service) Links() map[string]string {
	return map[string]string{
		"site": "https://lido.fi/",
	}
}

func (s *Service) SupportedNetworks() []string {
	return []string{
		"ethereum",
	}
}

func (s *Service) Groups() []base.BalanceGroup {
	return []base.BalanceGroup{
		{
			Type:  "staking",
			Label: "Staking",
		},
	}
}

func (s *Service) FetchBalances(tokens []sdk.Token) ([]base.Balance, error) {
	return []base.Balance{}, nil
}

func (s *Service) FetchStakingBalance(tokens []sdk.Token) (base.Balance, error) {
	return base.Balance{}, nil
}

func (s *Service) FetchFarmingBalance(tokens []sdk.Token) (base.Balance, error) {
	return base.Balance{}, nil
}

func (s *Service) FetchLendingBalance(tokens []sdk.Token) (base.Balance, error) {
	return base.Balance{}, nil
}

func (s *Service) FetchBorrowBalance(tokens []sdk.Token) (base.Balance, error) {
	return base.Balance{}, nil
}
