package gdax

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/lysrt/cryptomarkets"
	"github.com/lysrt/cryptomarkets/internal"
)

type gdaxTicker struct {
	TradeId int64     `json:"trade_id"`
	Price   float64   `json:"price,string"`
	Size    float64   `json:"size,string"`
	Bid     float64   `json:"bid,string"`
	Ask     float64   `json:"ask,string"`
	Volume  float64   `json:"volume,string"`
	Time    time.Time `json:"time"`
}

func (e *Gdax) GetTicker(from, to string) (*cryptomarkets.Ticker, error) {
	currencyPair := cryptomarkets.Pair{
		First:  cryptomarkets.NewCurrency(from),
		Second: cryptomarkets.NewCurrency(to),
	}

	url := fmt.Sprintf("https://api.gdax.com/products/%s/ticker", currencyPair.Upper("-"))

	body, err := internal.Get(url, map[string]string{"Content-Type": "application/json"})
	if err != nil {
		return nil, fmt.Errorf("%q: %s", err, string(body))
	}

	var t gdaxTicker
	err = json.Unmarshal(body, &t)
	if err != nil {
		return nil, err
	}

	// TODO can get more data using stats: /products/<product-id>/stats
	// s.Low, s.High, s.Open, s.Volume (s.Last, s.Volume30Day)

	return &cryptomarkets.Ticker{
		Timestamp:     t.Time.Unix(),
		LastPrice:     t.Price,
		LastQuantity:  t.Size,
		High:          0,
		Low:           0,
		Open:          0,
		Close:         0,
		Ask:           t.Ask,
		AskQuantity:   0,
		Bid:           t.Bid,
		BidQuantity:   0,
		VWAP:          0,
		Volume:        t.Volume,
		QuoteVolume:   0,
		PriceChange:   0,
		PercentChange: 0,
		Pair:          currencyPair,
	}, nil
}

// TODO Could use /products to list all available currency pairs

func (e *Gdax) OrderBook(from, to string) (*cryptomarkets.OrderBook, error) {
	return nil, errors.New("unimplemented")
}
