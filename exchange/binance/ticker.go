package binance

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/lysrt/cryptomarkets/common"

	"github.com/lysrt/cryptomarkets/entity"
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

func (e *Binance) GetTicker(from, to string) (*entity.Ticker, error) {
	currencyPair := entity.Pair{
		First:  entity.NewCurrency(from),
		Second: entity.NewCurrency(to),
	}

	url := fmt.Sprintf("https://api.binance.com/api/v1/ticker/24hr?symbol=%s", currencyPair.Upper(""))

	body, err := common.Get(url, nil)
	if err != nil {
		return nil, err
	}

	var t binanceTicker
	err = json.Unmarshal(body, &t)
	if err != nil {
		return nil, err
	}

	return &entity.Ticker{
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

	body, err := common.Get(url, nil)
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

func (e *Binance) OrderBook(from, to string) (*entity.OrderBook, error) {
	return nil, errors.New("unimplemented")
}
