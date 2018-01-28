package quoinex

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/lysrt/cryptomarkets"
	"github.com/lysrt/cryptomarkets/internal"
)

// ids lists all available products on Quoinex (2018-01-28)
var ids = map[string]int{
	"BTCUSD":   1,
	"BTCEUR":   3,
	"BTCJPY":   5,
	"BTCSGD":   7,
	"BTCHKD":   9,
	"BTCIDR":   11,
	"BTCAUD":   13,
	"BTCPHP":   15,
	"BTCCNY":   17,
	"BTCINR":   18,
	"ETHUSD":   27,
	"ETHEUR":   28,
	"ETHJPY":   29,
	"ETHSGD":   30,
	"ETHHKD":   31,
	"ETHIDR":   32,
	"ETHAUD":   33,
	"ETHPHP":   34,
	"ETHCNY":   35,
	"ETHINR":   36,
	"ETHBTC":   37,
	"BCHUSD":   39,
	"BCHSGD":   40,
	"BCHJPY":   41,
	"DASHSGD":  42,
	"DASHUSD":  43,
	"DASHJPY":  44,
	"DASHEUR":  45,
	"QTUMSGD":  46,
	"QTUMUSD":  47,
	"QTUMJPY":  48,
	"QTUMEUR":  49,
	"QASHJPY":  50,
	"QASHETH":  51,
	"QASHBTC":  52,
	"NEOUSD":   53,
	"NEOJPY":   54,
	"NEOSGD":   55,
	"NEOEUR":   56,
	"QASHUSD":  57,
	"QASHEUR":  58,
	"QASHSGD":  59,
	"QASHAUD":  60,
	"QASHIDR":  61,
	"QASHHKD":  62,
	"QASHPHP":  63,
	"QASHCNY":  64,
	"QASHINR":  65,
	"UBTCUSD":  71,
	"UBTCJPY":  72,
	"UBTCSGD":  73,
	"UBTCBTC":  74,
	"UBTCETH":  75,
	"UBTCQASH": 76,
	"XRPJPY":   83,
	"XRPUSD":   84,
	"XRPEUR":   85,
	"XRPSGD":   86,
	"XRPIDR":   87,
	"XRPQASH":  88,
}

type quoinexTicker struct {
	ID                 int         `json:"id,string"`
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

func (e *Quoinex) GetTicker(from, to string) (*cryptomarkets.Ticker, error) {
	currencyPair := cryptomarkets.Pair{
		First:  cryptomarkets.NewCurrency(from),
		Second: cryptomarkets.NewCurrency(to),
	}

	id, ok := ids[currencyPair.Upper("")]
	if !ok {
		return nil, fmt.Errorf("unavailable pair: %s", currencyPair.Upper("/"))
	}
	urlString := fmt.Sprintf("https://api.quoine.com/products/%d", id)

	body, err := internal.Get(urlString, nil)
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

	return &cryptomarkets.Ticker{
		Timestamp:     time.Now().Unix(),
		LastPrice:     last,
		LastQuantity:  quantity,
		High:          high,
		Low:           low,
		Open:          0,
		Close:         0,
		Ask:           t.MarketAsk,
		AskQuantity:   0,
		Bid:           t.MarketBid,
		BidQuantity:   0,
		VWAP:          0,
		Volume:        volume,
		QuoteVolume:   0,
		PriceChange:   0,
		PercentChange: 0,
		Pair:          currencyPair,
	}, nil
}

func (e *Quoinex) OrderBook(from, to string) (*cryptomarkets.OrderBook, error) {
	return nil, errors.New("unimplemented")
}
