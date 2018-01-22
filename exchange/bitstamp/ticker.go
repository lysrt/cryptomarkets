package bitstamp

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/lysrt/cryptomarkets/common"
	"github.com/lysrt/cryptomarkets/entity"
)

type bitstampTicker struct {
	High      float64 `json:"high,string"`
	Last      float64 `json:"last,string"`
	Timestamp int64   `json:"timestamp,string"`
	Bid       float64 `json:"bid,string"`
	VWAP      float64 `json:"vwap,string"`
	Volume    float64 `json:"volume,string"`
	Low       float64 `json:"low,string"`
	Ask       float64 `json:"ask,string"`
	Open      float64 `json:"open,string"`
}

/*
	Supported currency pairs (20180105):
	(https://www.bitstamp.net/api/v2/trading-pairs-info/)
	btcusd, btceur,
	eurusd,
	xrpusd, xrpeur, xrpbtc,
	ltcusd, ltceur, ltcbtc,
	ethusd, etheur, ethbtc,
	bchusd, bcheur, bchbtc
*/

// https://www.bitstamp.net/api/v2/ticker_hour/{currency_pair}/

func (e *Bitstamp) GetTicker(from, to string) (*entity.Ticker, error) {
	currencyPair := entity.Pair{
		First:  entity.NewCurrency(from),
		Second: entity.NewCurrency(to),
	}

	url := fmt.Sprintf("https://www.bitstamp.net/api/v2/ticker/%s/", currencyPair.Lower(""))

	body, err := common.Get(url, nil)
	if err != nil {
		return nil, fmt.Errorf("bad HTTP response: %q", err.Error())
	}

	var t bitstampTicker

	err = json.Unmarshal(body, &t)
	if err != nil {
		return nil, err
	}

	return &entity.Ticker{
		Timestamp:     t.Timestamp,
		LastPrice:     t.Last,
		LastQuantity:  0,
		High:          t.High,
		Low:           t.Low,
		Open:          t.Open,
		Close:         0,
		Ask:           t.Ask,
		AskQuantity:   0,
		Bid:           t.Bid,
		BidQuantity:   0,
		VWAP:          t.VWAP,
		Volume:        t.Volume,
		QuoteVolume:   0,
		PriceChange:   0,
		PercentChange: 0,
		Pair:          currencyPair,
	}, nil
}

type bitstampOrderBook struct {
	Timestamp int64      `json:"timestamp,string"`
	Bids      [][]string `json:"bids, string"`
	Asks      [][]string `json:"asks, string"`
}

func (e *Bitstamp) OrderBook(from, to string) (*entity.OrderBook, error) {
	currencyPair := entity.Pair{
		First:  entity.NewCurrency(from),
		Second: entity.NewCurrency(to),
	}

	url := fmt.Sprintf("https://www.bitstamp.net/api/v2/order_book/%s/", currencyPair.Lower(""))

	body, err := common.Get(url, nil)
	if err != nil {
		return nil, fmt.Errorf("bad HTTP response: %q", err.Error())
	}

	var o bitstampOrderBook

	err = json.Unmarshal(body, &o)
	if err != nil {
		return nil, err
	}

	bids := []entity.Order{}
	for _, b := range o.Bids {
		price, err := strconv.ParseFloat(b[0], 64)
		if err != nil {
			return nil, err
		}
		quantity, err := strconv.ParseFloat(b[1], 64)
		if err != nil {
			return nil, err
		}
		bids = append(bids, entity.Order{
			Price:    price,
			Quantity: quantity,
		})
	}

	asks := []entity.Order{}
	for _, a := range o.Asks {
		price, err := strconv.ParseFloat(a[0], 64)
		if err != nil {
			return nil, err
		}
		quantity, err := strconv.ParseFloat(a[1], 64)
		if err != nil {
			return nil, err
		}
		asks = append(asks, entity.Order{
			Price:    price,
			Quantity: quantity,
		})
	}

	return &entity.OrderBook{
		Timestamp: o.Timestamp,
		Asks:      asks,
		Bids:      bids,
	}, nil
}
