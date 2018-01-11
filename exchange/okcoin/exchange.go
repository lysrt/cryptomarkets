package okcoin

import "github.com/lysrt/cryptomarkets"

type Okcoin struct {
	CustomerID string
	ApiKey     string
	Secret     string
	private    bool
}

func New(c cryptomarkets.ExchangeConfig) *Okcoin {
	private := c.CustomerID != "" && c.ApiKey != "" && c.Secret != ""

	return &Okcoin{c.CustomerID, c.ApiKey, c.Secret, private}
}
