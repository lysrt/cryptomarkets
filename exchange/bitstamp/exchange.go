package bitstamp

import "github.com/lysrt/cryptomarkets"

type Bitstamp struct {
	CustomerID string
	ApiKey     string
	Secret     string
	private    bool
}

func New(c cryptomarkets.ExchangeConfig) *Bitstamp {
	private := c.CustomerID != "" && c.ApiKey != "" && c.Secret != ""

	return &Bitstamp{c.CustomerID, c.ApiKey, c.Secret, private}
}
