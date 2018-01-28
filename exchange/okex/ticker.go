package okex

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/lysrt/cryptomarkets"
	"github.com/lysrt/cryptomarkets/internal"
)

type okexTickerResponse struct {
	Date      int64      `json:"date,string"`
	Ticker    okexTicker `json:"ticker"`
	ErrorCode int64      `json:"error_code"`
}

type okexTicker struct {
	High float64 `json:"high,string"`
	Vol  float64 `json:"vol,string"`
	Last float64 `json:"last,string"`
	Low  float64 `json:"low,string"`
	Buy  float64 `json:"buy,string"`
	Sell float64 `json:"sell,string"`
}

func (e *Okex) GetTicker(from, to string) (*cryptomarkets.Ticker, error) {
	currencyPair := cryptomarkets.Pair{
		First:  cryptomarkets.NewCurrency(from),
		Second: cryptomarkets.NewCurrency(to),
	}

	url := fmt.Sprintf("https://www.okcoin.com/api/v1/ticker.do?symbol=%s", currencyPair.Lower("_"))

	body, err := internal.Get(url, map[string]string{"Content-Type": "application/x-www-form-urlencoded"})
	if err != nil {
		return nil, fmt.Errorf("bad HTTP response: %q", err.Error())
	}

	var resp okexTickerResponse

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}

	if resp.ErrorCode != 0 {
		return nil, fmt.Errorf("okcoin error code: %d", resp.ErrorCode)
	}

	return &cryptomarkets.Ticker{
		Timestamp:     resp.Date,
		LastPrice:     resp.Ticker.Last,
		LastQuantity:  0,
		High:          resp.Ticker.High,
		Low:           resp.Ticker.Low,
		Open:          0,
		Close:         0,
		Ask:           resp.Ticker.Sell,
		AskQuantity:   0,
		Bid:           resp.Ticker.Buy,
		BidQuantity:   0,
		VWAP:          0,
		Volume:        resp.Ticker.Vol,
		QuoteVolume:   0,
		PriceChange:   0,
		PercentChange: 0,
		Pair:          currencyPair,
	}, nil
}

func (e *Okex) OrderBook(from, to string) (*cryptomarkets.OrderBook, error) {
	return nil, errors.New("unimplemented")
}
