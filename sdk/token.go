package sdk

import (
	"context"
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

	var issues []*Token
	resp, err := s.client.Do(ctx, req, &issues)
	if err != nil {
		return nil, resp, err
	}

	return issues, resp, nil
}
