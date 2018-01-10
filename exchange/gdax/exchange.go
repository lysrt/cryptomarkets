package gdax

import "github.com/lysrt/cryptomarkets"

type Gdax struct {
	ApiKey  string
	Secret  string
	private bool
}

func New(c cryptomarkets.ExchangeConfig) *Gdax {
	private := c.ApiKey != "" && c.Secret != ""

	return &Gdax{c.ApiKey, c.Secret, private}
}
