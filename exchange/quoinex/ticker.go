package quoinex

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type ticker struct {
	ID                 string      `json:"id"`
	ProductType        string      `json:"product_type"`
	MarketAsk          float64     `json:"market_ask"`
	MarketBid          float64     `json:"market_bid"`
	CurrencyPairCode   string      `json:"currency_pair_code"`
	LowMarketBid       json.Number `json:"low_market_bid"`
	HighMarketAsk      json.Number `json:"high_market_ask"`
	Volume24h          json.Number `json:"volume_24h"`
	LastPrice24h       json.Number `json:"last_price_24h"`
	LastTradedPrice    json.Number `json:"last_traded_price"`
	LastTradedQuantity json.Number `json:"last_traded_quantity"`
	QuotedCurrency     string      `json:"quoted_currency"`
	BaseCurrency       string      `json:"base_currency"`
}

func (e *Quoinex) LastPrice(from, to string) float64 {
	url := "https://api.quoine.com/products"

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var tickers []ticker
	err = json.Unmarshal(body, &tickers)
	// err = json.NewDecoder(resp.Body).Decode(&tickers)
	if err != nil {
		panic(err)
	}

	id := ""
	for _, t := range tickers {
		if t.CurrencyPairCode == "BTCUSD" {
			id = t.ID
			break
		}
	}
	if id == "" {
		return -1
	}

	// ---------------------
	url = "https://api.quoine.com/products/" + id

	resp, err = http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var ticker ticker
	err = json.Unmarshal(body, &ticker)
	if err != nil {
		panic(err)
	}

	result, err := ticker.LastTradedPrice.Float64()
	if err != nil {
		panic(err)
	}
	return result
}
