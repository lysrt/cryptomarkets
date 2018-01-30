package binance

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/lysrt/cryptomarkets"
	"github.com/lysrt/cryptomarkets/internal"
)

type binanceTicker struct {
	Symbol             string  `json:"symbol"`
	PriceChange        float64 `json:"priceChange,string"`
	PriceChangePercent float64 `json:"priceChangePercent,string"`
	WeightedAvgPrice   float64 `json:"weightedAvgPrice,string"`
	PrevClosePrice     float64 `json:"prevClosePrice,string"`
	LastPrice          float64 `json:"lastPrice,string"`
	LastQty            float64 `json:"lastQty,string"`
	BidPrice           float64 `json:"bidPrice,string"`
	BidQty             float64 `json:"bidQty,string"`
	AskPrice           float64 `json:"askPrice,string"`
	AskQty             float64 `json:"askQty,string"`
	OpenPrice          float64 `json:"openPrice,string"`
	HighPrice          float64 `json:"highPrice,string"`
	LowPrice           float64 `json:"lowPrice,string"`
	Volume             float64 `json:"volume,string"`
	QuoteVolume        float64 `json:"quoteVolume,string"`
	OpenTime           int64   `json:"openTime"`
	CloseTime          int64   `json:"closeTime"`
	FirstId            int64   `json:"firstId"`
	LastId             int64   `json:"lastId"`
	Count              int64   `json:"count"`
}

func (e *Binance) GetTicker(from, to string) (*cryptomarkets.Ticker, error) {
	currencyPair := cryptomarkets.Pair{
		First:  cryptomarkets.NewCurrency(from),
		Second: cryptomarkets.NewCurrency(to),
	}

	url := fmt.Sprintf("https://api.binance.com/api/v1/ticker/24hr?symbol=%s", currencyPair.Upper(""))

	body, err := internal.Get(url, nil)
	if err != nil {
		return nil, err
	}

	var t binanceTicker
	err = json.Unmarshal(body, &t)
	if err != nil {
		return nil, err
	}

	return &cryptomarkets.Ticker{
		Timestamp:     time.Now().Unix(),
		LastPrice:     t.LastPrice,
		LastQuantity:  t.LastQty,
		High:          t.HighPrice,
		Low:           t.LowPrice,
		Open:          t.OpenPrice,
		Close:         t.PrevClosePrice,
		Ask:           t.AskPrice,
		AskQuantity:   t.AskQty,
		Bid:           t.BidPrice,
		BidQuantity:   t.BidQty,
		VWAP:          t.WeightedAvgPrice,
		Volume:        t.Volume,
		QuoteVolume:   t.QuoteVolume,
		PriceChange:   t.PriceChange,
		PercentChange: t.PriceChangePercent,
		Pair:          currencyPair,
	}, nil
}

func (e *Binance) PrintAllPrices() error {
	url := "https://api.binance.com/api/v1/ticker/allPrices"

	body, err := internal.Get(url, nil)
	if err != nil {
		return err
	}

	var prices []struct {
		Symbol string  `json:"symbol"`
		Price  float64 `json:"price,string"`
	}

	err = json.Unmarshal(body, &prices)
	if err != nil {
		return err
	}

	for _, p := range prices {
		fmt.Printf("%s: %f\n", p.Symbol, p.Price)
	}

	return nil
}

type binanceOrderBook struct {
	LastUpdateId int             `json:"lastUpdateId"`
	Bids         [][]interface{} `json:"bids"`
	Asks         [][]interface{} `json:"asks"`
}

/*
{
	"lastUpdateId": 50325974,
	"bids": [
	  [
		"10481.64000000",
		"0.26734600",
		[]
	  ],
	  ...
	],
	"asks": [
	  [
		"10489.71000000",
		"0.15654600",
		[]
	  ],
	  ...
	]
  }
*/
func (e *Binance) OrderBook(from, to string) (*cryptomarkets.OrderBook, error) {
	currencyPair := cryptomarkets.Pair{
		First:  cryptomarkets.NewCurrency(from),
		Second: cryptomarkets.NewCurrency(to),
	}

	// ignore the limit param, Default 100; max 1000. Valid limits:[5, 10, 20, 50, 100, 500, 1000]
	url := fmt.Sprintf("https://api.binance.com/api/v1/depth?symbol=%s", currencyPair.Upper(""))

	body, err := internal.Get(url, nil)
	if err != nil {
		return nil, fmt.Errorf("bad HTTP response: %q", err.Error())
	}

	var o binanceOrderBook

	err = json.Unmarshal(body, &o)
	if err != nil {
		return nil, err
	}

	bids := []cryptomarkets.Order{}
	for _, b := range o.Bids {
		priceStr := b[0].(string)
		qtyStr := b[1].(string)

		price, err := strconv.ParseFloat(priceStr, 64)
		if err != nil {
			return nil, err
		}
		quantity, err := strconv.ParseFloat(qtyStr, 64)
		if err != nil {
			return nil, err
		}
		bids = append(bids, cryptomarkets.Order{
			Price:    price,
			Quantity: quantity,
		})
	}

	asks := []cryptomarkets.Order{}
	for _, a := range o.Asks {
		priceStr := a[0].(string)
		qtyStr := a[1].(string)

		price, err := strconv.ParseFloat(priceStr, 64)
		if err != nil {
			return nil, err
		}
		quantity, err := strconv.ParseFloat(qtyStr, 64)
		if err != nil {
			return nil, err
		}
		asks = append(asks, cryptomarkets.Order{
			Price:    price,
			Quantity: quantity,
		})
	}

	return &cryptomarkets.OrderBook{
		Timestamp: time.Now().Unix(),
		Asks:      asks,
		Bids:      bids,
	}, nil
}
