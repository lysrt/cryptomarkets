package main

import (
	"fmt"
	"log"

	"github.com/lysrt/cryptomarkets/entity"

	"github.com/lysrt/cryptomarkets/exchange/bitstamp"
	"github.com/lysrt/cryptomarkets/exchange/quoinex"
)

type pricer interface {
	GetTicker(from, to string) (*entity.Ticker, error)
}

func main() {
	providers := map[string]pricer{
		"bitstamp": &bitstamp.Bitstamp{},
		"quoinex":  &quoinex.Quoinex{},
	}

	bp, err := providers["bitstamp"].GetTicker("BTC", "USD")
	if err != nil {
		log.Fatal(err)
	}

	qp, err := providers["quoinex"].GetTicker("BTC", "USD")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Bitstamp: ", bp.LastPrice)
	fmt.Println("Quoinex: ", qp.LastPrice)
}
