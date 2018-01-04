package example

import (
	"github.com/lysrt/cryptomarkets/exchange/bitstamp"
)

func getBitstampPrice() float64 {
	bitstamp := bitstamp.Bitstamp{}
	return bitstamp.LastPrice("BTC", "USD")
}
