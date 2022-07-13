package curve

import (
	"time"

	"github.com/depocket/lab/base"
	"github.com/depocket/lab/constants"
	"github.com/ethereum/go-ethereum/common"
)

func (s *Adapter) Balances(chain string, project string, ad string) (base.ProjectBalance, error) {
	var balance = base.ProjectBalance{
		Chain:     chain,
		Name:      s.ApplicationManifest().Name,
		LogoUrl:   constants.GetTokenIcon(chain, s.ApplicationManifest().Code),
		TagIds:    s.ApplicationManifest().Tags,
		UpdatedAt: time.Now().Unix(),
		Version:   constants.ResponseSchemaVersion,
	}

	balances := []base.Balance{}

	switch chain {
	case constants.ChainEthereum:
		val, err := s.VestingBalances(common.HexToAddress(ad))
		if err != nil {
			return balance, err
		}
		balances = append(balances, *val)
	}

	balance.Balances = balances
	return balance, nil
}
