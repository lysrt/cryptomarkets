package bitstamp

import (
	"encoding/json"
	"fmt"

	"github.com/lysrt/cryptomarkets/common"
	"github.com/lysrt/cryptomarkets/currency"
	"github.com/lysrt/cryptomarkets/ticker"
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

func (e *Bitstamp) Ticker(from, to string) (*ticker.Ticker, error) {
	currencyPair := currency.Pair{
		First:  currency.New(from),
		Second: currency.New(to),
	}

	url := fmt.Sprintf("https://www.bitstamp.net/api/v2/ticker/%s/", currencyPair.Lower(""))

	body, err := common.RunRequest(url)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(body))

	var t bitstampTicker

	err = json.Unmarshal(body, &t)
	if err != nil {
		return nil, err
	}

	return &ticker.Ticker{
		Timestamp:    t.Timestamp,
		LastPrice:    t.Last,
		LastQuantity: 0,
		High:         t.High,
		Low:          t.Low,
		Open:         t.Open,
		Close:        0,
		Bid:          t.Bid,
		Ask:          t.Ask,
		VWAP:         t.VWAP,
		Volume:       t.Volume,
		Pair:         currencyPair,
	}, nil
}
