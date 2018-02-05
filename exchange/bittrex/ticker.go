package bittrex

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/lysrt/cryptomarkets"
	"github.com/lysrt/cryptomarkets/internal"
)

type bittrexTickerResponse struct {
	Success bool            `json:"success"`
	Message string          `json:"message"`
	Result  []bittrexTicker `json:"result"`
}

type bittrexTicker struct {
	MarketName     string      `json:"MarketName"`
	High           float64     `json:"High"`
	Low            float64     `json:"Low"`
	Volume         float64     `json:"Volume"`
	Last           float64     `json:"Last"`
	BaseVolume     float64     `json:"BaseVolume"`
	TimeStamp      bittrexTime `json:"TimeStamp"` //		"2018-01-10T21:24:18.347",
	Bid            float64     `json:"Bid"`
	Ask            float64     `json:"Ask"`
	OpenBuyOrders  int64       `json:"OpenBuyOrders"`
	OpenSellOrders int64       `json:"OpenSellOrders"`
	PrevDay        float64     `json:"PrevDay"`
	Created        bittrexTime `json:"Created"` //		"2014-02-13T00:00:00"
}

type bittrexTime time.Time

func (bt *bittrexTime) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	format := "2006-01-02T15:04:05.000"
	for len(s) < 23 {
		s = s + "0"
	}

	t, err := time.Parse(format, s)
	if err != nil {
		return err
	}
	*bt = (bittrexTime)(t)
	return nil
}

func (e *Bittrex) GetTicker(from, to string) (*cryptomarkets.Ticker, error) {
	currencyPair := cryptomarkets.Pair{
		First:  cryptomarkets.NewCurrency(from),
		Second: cryptomarkets.NewCurrency(to),
	}

	url := fmt.Sprintf("https://bittrex.com/api/v1.1/public/getmarketsummary?market=%s", currencyPair.Lower("-"))

	body, err := internal.Get(url, nil)
	if err != nil {
		return nil, err
	}

	var r bittrexTickerResponse
	err = json.Unmarshal(body, &r)
	if err != nil {
		return nil, err
	}

	if r.Success == false {
		// TODO Put currency pair in the message?
		return nil, fmt.Errorf("bad API response: %s", r.Message)
	}

	tickers := r.Result
	if len(tickers) != 1 {
		return nil, fmt.Errorf("bad tickers count (%d), 1 expected", len(tickers))
	}

	t := tickers[0]

	// TODO Could use this info
	// t.OpenBuyOrders
	// t.OpenSellOrders

	return &cryptomarkets.Ticker{
		Timestamp:     time.Time(t.TimeStamp).Unix(),
		LastPrice:     t.Last,
		LastQuantity:  0,
		High:          t.High,
		Low:           t.Low,
		Open:          0,
		Close:         t.PrevDay,
		Ask:           t.Ask,
		AskQuantity:   0,
		Bid:           t.Bid,
		BidQuantity:   0,
		VWAP:          0,
		Volume:        t.Volume,
		QuoteVolume:   t.BaseVolume,
		PriceChange:   0,
		PercentChange: 0,
		Pair:          currencyPair,
	}, nil
}

// TODO Could use "getAllMarkets"

type bittrexOrderBookResponse struct {
	Success bool             `json:"success"`
	Message string           `json:"message"`
	Result  bittrexOrderBook `json:"result"`
}

type bittrexOrderBook struct {
	Buy  []bittrexOrder `json:"buy"`
	Sell []bittrexOrder `json:"sell"`
}

type bittrexOrder struct {
	Quantity float64 `json:"Quantity"`
	Rate     float64 `json:"Rate"`
}

func (e *Bittrex) OrderBook(from, to string) (*cryptomarkets.OrderBook, error) {
	currencyPair := cryptomarkets.Pair{
		First:  cryptomarkets.NewCurrency(from),
		Second: cryptomarkets.NewCurrency(to),
	}

	url := fmt.Sprintf("https://bittrex.com/api/v1.1/public/getorderbook?market=%s&type=both", currencyPair.Lower("-"))

	body, err := internal.Get(url, nil)
	if err != nil {
		return nil, err
	}

	var r bittrexOrderBookResponse
	err = json.Unmarshal(body, &r)
	if err != nil {
		return nil, err
	}

	if r.Success == false {
		// TODO Put currency pair in the message?
		return nil, fmt.Errorf("bad API response: %s", r.Message)
	}

	bids := []cryptomarkets.BookOrder{}
	for _, b := range r.Result.Buy {
		bids = append(bids, cryptomarkets.BookOrder{
			Price:    b.Rate,
			Quantity: b.Quantity,
		})
	}

	asks := []cryptomarkets.BookOrder{}
	for _, a := range r.Result.Sell {
		asks = append(asks, cryptomarkets.BookOrder{
			Price:    a.Rate,
			Quantity: a.Quantity,
		})
	}

	return &cryptomarkets.OrderBook{
		Timestamp: time.Now().Unix(),
		Asks:      asks,
		Bids:      bids,
	}, nil
}
