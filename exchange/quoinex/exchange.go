package quoinex

import "github.com/lysrt/cryptomarkets"

type Quoinex struct {
	CustomerID string
	ApiKey     string
	Secret     string
	private    bool
}

func New(c cryptomarkets.ExchangeConfig) *Quoinex {
	private := c.CustomerID != "" && c.ApiKey != "" && c.Secret != ""

	return &Quoinex{c.CustomerID, c.ApiKey, c.Secret, private}
}
