package bittrex

import "github.com/lysrt/cryptomarkets"

type Bittrex struct {
	ApiKey  string
	Secret  string
	private bool
}

func New(c cryptomarkets.ExchangeConfig) *Bittrex {
	private := c.ApiKey != "" && c.Secret != ""

	return &Bittrex{c.ApiKey, c.Secret, private}
}
