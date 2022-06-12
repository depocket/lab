package pancakeswap

import (
	"context"
	"fmt"
	"github.com/depocket/go-client"
	"github.com/depocket/lab/base"
	"github.com/depocket/lab/constants"
	"github.com/depocket/lab/constants/balancegroup"
	"github.com/depocket/lab/constants/tags"
	"github.com/depocket/lab/core"
)

func init() {
	core.ServiceRegisters.RegisterService("pancakeswap", &Adapter{})
}

type Adapter struct {
	Tokens []*client.Token
	Pools  []*client.Pool
}

func (s *Adapter) ApplicationManifest() base.Manifest {
	return base.Manifest{
		Code: "pancakeswap",
		Name: "PancakeSwap",
		Tags: []string{tags.Staking},
		Description: "Cheaper and faster than Uniswap? " +
			"Discover PancakeSwap, the leading DEX on BNB Smart Chain (BSC) " +
			"with the best farms in DeFi and a lottery for CAKE.",
		Url: "https://pancakeswap.finance/",
		Links: map[string]string{
			"github":   "https://github.com/pancakeswap/",
			"twitter":  "https://twitter.com/pancakeswap",
			"discord":  "https://discord.gg/pancakeswap",
			"telegram": "https://t.me/pancakeswap",
		},
		SupportedNetworks: []string{
			constants.ChainBinance,
		},
		Groups: []base.BalanceGroup{
			balancegroup.Staking,
		},
	}
}

func (s *Adapter) InjectDependencies(chain string, cl *client.Client) error {
	ctx := context.Background()
	list, _, err := cl.Tokens.List(ctx, &client.TokenListOptions{
		ListOptions: client.ListOptions{
			Chain: chain,
		},
	})
	if err != nil {
		return err
	}
	s.Tokens = list
	pools, _, err := cl.Pools.ListByProjectCode(ctx, fmt.Sprintf("%s_%s", chain, s.ApplicationManifest().Code), &client.PoolListOptions{
		ListOptions: client.ListOptions{
			Chain: chain,
		},
	})
	s.Pools = pools
	return nil
}
