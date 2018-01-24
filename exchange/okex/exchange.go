package okex

import "github.com/lysrt/cryptomarkets"

type Okex struct {
	CustomerID string
	ApiKey     string
	Secret     string
	private    bool
}

func New(c cryptomarkets.ExchangeConfig) *Okex {
	private := c.CustomerID != "" && c.ApiKey != "" && c.Secret != ""

	return &Okex{c.CustomerID, c.ApiKey, c.Secret, private}
}
