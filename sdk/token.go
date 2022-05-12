package sdk

import (
	"context"
	"fmt"
)

type TokenService service

type Token struct {
	Address  string  `json:"address"`
	Name     string  `json:"name"`
	Symbol   string  `json:"symbol"`
	IconUrl  string  `json:"icon_url"`
	Type     string  `json:"type"`
	Decimals int     `json:"decimals"`
	Price    float64 `json:"price"`
}

type TokensResponse struct {
	Data      []*Token `json:"data"`
	ErrorCode int      `json:"error_code"`
}

func (i Token) String() string {
	return Stringify(i)
}

type TokenListOptions struct {
	ListOptions
}

func (s *TokenService) List(ctx context.Context, opts *TokenListOptions) ([]*Token, *Response, error) {
	var u = "tokens"
	return s.listTokens(ctx, u, opts)
}

func (s *TokenService) listTokens(ctx context.Context, u string, opts *TokenListOptions) ([]*Token, *Response, error) {
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var tokensResponse *TokensResponse
	resp, err := s.client.Do(ctx, req, &tokensResponse)
	if err != nil {
		return nil, resp, err
	}
	if tokensResponse.ErrorCode != 0 {
		return nil, resp, fmt.Errorf("error code %d", tokensResponse.ErrorCode)
	}

	return tokensResponse.Data, resp, nil
}
