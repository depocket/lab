package pancakeswap

import (
	"github.com/depocket/lab/base"
	"github.com/depocket/lab/constants"
	"time"
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
	switch chain {
	case constants.ChainEthereum:
	}
	return balance, nil
}
