package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/lysrt/cryptomarkets/entity"
	"github.com/lysrt/cryptomarkets/exchange/binance"
	"github.com/lysrt/cryptomarkets/exchange/bitstamp"
	"github.com/lysrt/cryptomarkets/exchange/bittrex"
	"github.com/lysrt/cryptomarkets/exchange/gdax"
	"github.com/lysrt/cryptomarkets/exchange/okex"
	"github.com/lysrt/cryptomarkets/exchange/quoinex"
)

type pricer interface {
	GetTicker(from, to string) (*entity.Ticker, error)
}

func main() {
	providers := map[string]pricer{
		"bitstamp": &bitstamp.Bitstamp{},
		"quoinex":  &quoinex.Quoinex{},
		"bittrex":  &bittrex.Bittrex{},
		"binance":  &binance.Binance{},
		"gdax":     &gdax.Gdax{},
		"okex":     &okex.Okex{},
	}

	wg := &sync.WaitGroup{}
	wg.Add(6)

	go getPrice(wg, providers["bitstamp"], "bitstamp", "BTC", "USD")
	go getPrice(wg, providers["quoinex"], "quoinex", "BTC", "USD")
	go getPrice(wg, providers["binance"], "binance", "BTC", "USDT")
	go getPrice(wg, providers["bittrex"], "bittrex", "USDT", "BTC")
	go getPrice(wg, providers["gdax"], "gdax", "BTC", "USD")
	go getPrice(wg, providers["okex"], "okex", "BTC", "USD")

	wg.Wait()
}

func getPrice(wg *sync.WaitGroup, exchange pricer, name, ccy1, ccy2 string) {
	ticker, err := exchange.GetTicker(ccy1, ccy2)
	if err != nil {
		log.Print(err)
	} else {
		fmt.Println(name, ":", ticker.LastPrice)
	}

	wg.Done()
}
