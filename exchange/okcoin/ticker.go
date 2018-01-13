package okcoin

import (
	"encoding/json"
	"fmt"

	"github.com/lysrt/cryptomarkets/common"
	"github.com/lysrt/cryptomarkets/entity"
)

type okcoinTickerResponse struct {
	Date      int64        `json:"date,string"`
	Ticker    okcoinTicker `json:"ticker"`
	ErrorCode int64        `json:"error_code"`
}

type okcoinTicker struct {
	High float64 `json:"high,string"`
	Vol  float64 `json:"vol,string"`
	Last float64 `json:"last,string"`
	Low  float64 `json:"low,string"`
	Buy  float64 `json:"buy,string"`
	Sell float64 `json:"sell,string"`
}

func (e *Okcoin) GetTicker(from, to string) (*entity.Ticker, error) {
	currencyPair := entity.Pair{
		First:  entity.NewCurrency(from),
		Second: entity.NewCurrency(to),
	}

	url := fmt.Sprintf("https://www.okcoin.com/api/v1/ticker.do?symbol=%s", currencyPair.Lower("_"))

	body, err := common.Get(url, map[string]string{"Content-Type": "application/x-www-form-urlencoded"})
	if err != nil {
		return nil, fmt.Errorf("bad HTTP response: %q", err.Error())
	}

	var resp okcoinTickerResponse

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}

	if resp.ErrorCode != 0 {
		return nil, fmt.Errorf("okcoin error code: %d", resp.ErrorCode)
	}

	return &entity.Ticker{
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
