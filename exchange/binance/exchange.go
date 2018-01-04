package binance

import "github.com/lysrt/cryptomarkets"

type Binance struct {
	ApiKey  string
	Secret  string
	private bool
}

func New(c cryptomarkets.ExchangeConfig) *Binance {
	private := c.ApiKey != "" && c.Secret != ""

	return &Binance{c.ApiKey, c.Secret, private}
}
