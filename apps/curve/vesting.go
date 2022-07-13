package curve

import (
	"errors"
	"math/big"
	"time"

	"github.com/depocket/go-client"
	"github.com/depocket/lab/base"
	"github.com/depocket/lab/constants"
	"github.com/depocket/lab/constants/balancegroup"
	"github.com/depocket/lab/utils"
	"github.com/depocket/multicall-go/call"
	"github.com/ethereum/go-ethereum/common"
	"github.com/samber/lo"
)

func (s *Adapter) VestingBalances(address common.Address) (*base.Balance, error) {
	balances := &base.Balance{
		Chain:  constants.ChainEthereum,
		Groups: []base.BalanceGroup{balancegroup.Vesting},
		Pools:  []base.Pool{},
	}

	vestingType, err := s.checkBalanceOf(address)
	if err != nil || vestingType == "" {
		return balances, err
	}
	contract, ok := VESTING_CONTRACTS[vestingType]
	if !ok {
		return balances, errors.New("unknown vesting contract")
	}

	crvToken, exist := s.getCrvToken()
	if !exist {
		return balances, errors.New("CRV token not found")
	}

	onChainData, err := s.getOnChainData(address, contract.Address)
	if err != nil {
		return nil, err
	}

	unlockTime := time.Unix(onChainData["end_time"][0].(*big.Int).Int64(), 0)
	// TODO:
	// Change SupplyTokens to LockedTokens and RewardTokens to ClaimableTokens
	// After UI is supported for display LockedTokens and ClaimableTokens
	balances.Pools = []base.Pool{{
		Name:          "CRV",
		Apy:           new(big.Float),
		SupplyTokens:  s.getPoolCRVBalances(onChainData["lockedOf"], crvToken),
		RewardTokens:  s.getPoolCRVBalances(onChainData["balanceOf"], crvToken),
		ClaimedTokens: s.getPoolCRVBalances(onChainData["total_claimed"], crvToken),
		UnlockAt:      &unlockTime,
	}}
	// balances.TypeName = contract.Name + " vesting"

	return balances, nil
}

func (s *Adapter) getCrvToken() (client.Token, bool) {
	token, exist := lo.Find(s.Tokens, func(token *client.Token) bool {
		return token.Address == CRV_TOKEN_ADDRESS
	})

	return *token, exist
}

func (s *Adapter) getOnChainData(address common.Address, contractAddr string) (map[string][]interface{}, error) {
	methodNames := []string{
		"balanceOf",
		"lockedOf",
		"total_claimed",
	}
	caller := call.NewContractBuilder().
		AddMethod("end_time()(uint256)").
		AddCall("end_time", contractAddr, "end_time")
	for _, methodName := range methodNames {
		caller.
			AddMethod(methodName+"(address)(uint256)").
			AddCall(methodName, contractAddr, methodName, address)
	}

	_, results, err := caller.Call(nil)
	return results, err
}

func (s *Adapter) getPoolCRVBalances(
	onChainBalances []interface{},
	crvToken client.Token,
) []base.TokenBalance {
	if len(onChainBalances) == 0 {
		return nil
	}

	balance := onChainBalances[0].(*big.Int)
	return []base.TokenBalance{
		utils.GetPoolTokenBalance(CRV_TOKEN_ADDRESS, balance, crvToken),
	}
}

func (s *Adapter) checkBalanceOf(address common.Address) (string, error) {
	caller := call.NewContractBuilder().AddMethod("balanceOf(address)(uint256)")
	for key, contract := range VESTING_CONTRACTS {
		caller.AddCall(key, contract.Address, "balanceOf", address)
	}

	_, res, err := caller.Call(nil)
	if err != nil {
		return "", err
	}

	for key := range VESTING_CONTRACTS {
		if res[key][0].(*big.Int).Cmp(big.NewInt(0)) > 0 {
			return key, nil
		}
	}

	return "", nil
}
