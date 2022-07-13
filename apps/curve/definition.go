package curve

import (
	"context"
	"fmt"
	"sync"

	"github.com/depocket/go-client"
	"github.com/depocket/lab/base"
	"github.com/depocket/lab/constants"
	"github.com/depocket/lab/constants/balancegroup"
	"github.com/depocket/lab/constants/tags"
	"github.com/depocket/lab/core"
)

func init() {
	core.ServiceRegisters.RegisterService("curve", &Adapter{})
}

type Adapter struct {
	Tokens []*client.Token
	Pools  []*client.Pool
}

func (s *Adapter) ApplicationManifest() base.Manifest {
	return base.Manifest{
		Code: "curve",
		Name: "Curve",
		Tags: []string{tags.Vesting},
		Description: `Curve is an exchange liquidity pool on Ethereum.
			Curve is designed for extremely efficient stablecoin trading
			and low risk, supplemental fee income for liquidity providers,
			without an opportunity cost.`,
		Url: "https://curve.fi",
		Links: map[string]string{
			"github":   "https://github.com/curvefi",
			"twitter":  "https://twitter.com/CurveFinance",
			"discord":  "https://discord.gg/9uEHakc",
			"telegram": "https://t.me/curvefi",
		},
		SupportedNetworks: []string{
			constants.ChainEthereum,
		},
		Groups: []base.BalanceGroup{
			balancegroup.Vesting,
		},
	}
}

func (s *Adapter) InjectDependencies(chain string, cl *client.Client) error {
	ctx := context.Background()
	var (
		wg  sync.WaitGroup
		err error
	)

	wg.Add(2)
	go func() {
		defer wg.Done()
		s.Tokens, _, err = cl.Tokens.List(ctx, &client.TokenListOptions{
			ListOptions: client.ListOptions{
				Chain: chain,
			},
		})
	}()
	go func() {
		defer wg.Done()
		s.Pools, _, err = cl.Pools.ListByProjectCode(ctx, fmt.Sprintf("%s_%s", chain, s.ApplicationManifest().Code), &client.PoolListOptions{
			ListOptions: client.ListOptions{
				Chain: chain,
			},
		})
	}()
	wg.Wait()

	return err
}
