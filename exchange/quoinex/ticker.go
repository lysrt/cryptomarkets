package quoinex

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/lysrt/cryptomarkets/common"

	"github.com/lysrt/cryptomarkets/currency"
	"github.com/lysrt/cryptomarkets/ticker"
)

type quoinexTicker struct {
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

func (e *Quoinex) Ticker(from, to string) (*ticker.Ticker, error) {
	currencyPair := currency.Pair{
		First:  currency.New(from),
		Second: currency.New(to),
	}

	url := "https://api.quoine.com/products"

	body, err := common.RunRequest(url)
	if err != nil {
		return nil, err
	}

	var tickers []quoinexTicker
	err = json.Unmarshal(body, &tickers)
	if err != nil {
		return nil, err
	}

	id := ""
	for _, t := range tickers {
		if t.CurrencyPairCode == "BTCUSD" {
			id = t.ID
			break
		}
	}
	if id == "" {
		return nil, errors.New("did not find currency pair")
	}

	// ---------------------
	url = "https://api.quoine.com/products/" + id

	body, err = common.RunRequest(url)
	if err != nil {
		return nil, err
	}

	var t quoinexTicker
	err = json.Unmarshal(body, &t)
	if err != nil {
		return nil, err
	}

	last, err := t.LastTradedPrice.Float64()
	quantity, err := t.LastTradedQuantity.Float64()
	high, err := t.HighMarketAsk.Float64()
	low, err := t.LowMarketBid.Float64()
	volume, err := t.Volume24h.Float64()
	if err != nil {
		return nil, err
	}

	return &ticker.Ticker{
		Timestamp:    time.Now().Unix(),
		LastPrice:    last,
		LastQuantity: quantity,
		High:         high,
		Low:          low,
		Open:         0,
		Close:        0,
		Bid:          t.MarketBid,
		Ask:          t.MarketAsk,
		VWAP:         0,
		Volume:       volume,
		Pair:         currencyPair,
	}, nil
}
