package balancegroup

import "github.com/depocket/lab/base"

var (
	Staking = base.BalanceGroup{
		Type:  "staking",
		Label: "Staking",
	}
	Vesting = base.BalanceGroup{
		Type:  "vesting",
		Label: "Vesting",
	}
)
