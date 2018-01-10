package main

import (
	"fmt"
	"log"

	"github.com/lysrt/cryptomarkets/entity"
	"github.com/lysrt/cryptomarkets/exchange/binance"
	"github.com/lysrt/cryptomarkets/exchange/bitstamp"
	"github.com/lysrt/cryptomarkets/exchange/bittrex"
	"github.com/lysrt/cryptomarkets/exchange/gdax"
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
	}

	bp, err := providers["bitstamp"].GetTicker("BTC", "USD")
	if err != nil {
		log.Fatal(err)
	}

	qp, err := providers["quoinex"].GetTicker("BTC", "USD")
	if err != nil {
		log.Fatal(err)
	}

	bip, err := providers["binance"].GetTicker("BTC", "USDT")
	if err != nil {
		log.Fatal(err)
	}

	btp, err := providers["bittrex"].GetTicker("USDT", "BTC")
	if err != nil {
		log.Fatal(err)
	}

	gp, err := providers["gdax"].GetTicker("BTC", "USD")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Bitstamp:", bp.LastPrice)
	fmt.Println("Quoinex: ", qp.LastPrice)
	fmt.Println("Binance: ", bip.LastPrice)
	fmt.Println("Bittrex: ", btp.LastPrice)
	fmt.Println("GDAX:    ", gp.LastPrice)
}
