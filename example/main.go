package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/lysrt/cryptomarkets/loader"
)

func main() {
	loader, err := loader.New("config.json")
	if err != nil {
		log.Fatal(err)
	}

	wg := &sync.WaitGroup{}
	wg.Add(6)

	go getPrice(wg, loader, "bitstamp", "BTC", "USD")
	go getPrice(wg, loader, "quoinex", "BTC", "USD")
	go getPrice(wg, loader, "binance", "BTC", "USDT")
	go getPrice(wg, loader, "bittrex", "USDT", "BTC")
	go getPrice(wg, loader, "gdax", "BTC", "USD")
	go getPrice(wg, loader, "okex", "BTC", "USD")

	wg.Wait()
}

func getPrice(wg *sync.WaitGroup, loader loader.Loader, name, ccy1, ccy2 string) {
	defer wg.Done()

	exchange, err := loader.GetExchange(name)
	if err != nil {
		log.Println(err)
		return
	}

	ticker, err := exchange.GetTicker(ccy1, ccy2)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(name, ":", ticker.LastPrice)
	}
}
