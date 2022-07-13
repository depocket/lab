package utils

import (
	"math/big"

	"github.com/depocket/go-client"
	"github.com/depocket/lab/base"
)

func ExpX(x int) *big.Float {
	ten := big.NewInt(10)
	return new(big.Float).SetInt(ten.Exp(ten, big.NewInt(int64(x)), nil))
}

func ConvertTokenBalance(wei *big.Int, decimals int) *big.Float {
	return new(big.Float).Quo(new(big.Float).SetInt(wei), ExpX(decimals))
}

func GetPoolTokenBalance(tokenAddress string, tokenBalance *big.Int, token client.Token) base.TokenBalance {
	return base.TokenBalance{
		Address:  tokenAddress,
		Balance:  ConvertTokenBalance(tokenBalance, token.Decimals),
		Name:     token.Name,
		Symbol:   token.Symbol,
		IconUrl:  token.IconUrl,
		Price:    token.Price,
		Chain:    "",
		Decimals: token.Decimals,
	}
}

func GetPoolTokenBalances(tokenBalances map[string]*big.Int, tokens map[string]*client.Token) []base.TokenBalance {
	balances := []base.TokenBalance{}
	for tokenAddress, tokenBalance := range tokenBalances {
		token, ok := tokens[tokenAddress]
		if !ok {
			continue
		}

		balance := ConvertTokenBalance(tokenBalance, token.Decimals)
		balances = append(balances, base.TokenBalance{
			Address:  tokenAddress,
			Balance:  balance,
			Name:     token.Name,
			Symbol:   token.Symbol,
			IconUrl:  token.IconUrl,
			Chain:    "",
			Decimals: token.Decimals,
			Price:    token.Price,
		})
	}

	return balances
}
