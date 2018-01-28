package okex

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/lysrt/cryptomarkets/internal"
)

type orderType string

const (
	limitBuy   orderType = "buy"
	limitSell            = "sell"
	marketBuy            = "buy_market"
	marketSell           = "sell_market"
)

func okexSymbol(ccy string) string {
	switch strings.ToLower(ccy) {
	case "btc":
		return "btc_usd"
	case "ltc":
		return "ltc_usd"
	case "eth":
		return "eth_usd"
	case "etc":
		return "etc_usd"
	case "bch":
		return "bch_usd"
	default:
		return ""
	}
}

type Side string

const (
	Buy  Side = "buy"
	Sell Side = "sell"
)

type okexOrderResponse struct {
	Result  bool `json:"result"`
	OrderID int  `json:"order_id"`
}

func (e *Okex) MarketOrder(from, to string, side Side, amount float64) (int, error) {
	urlString := "https://www.okcoin.com/api/v1/trade.do"

	fromCurrency := okexSymbol(from)
	if fromCurrency == "" {
		return 0, fmt.Errorf("okcoin cannot withdraw %s, only: btc, ltc, eth, etc, bch", from)
	}

	toCurrency := strings.ToLower(to)
	if toCurrency != "usd" {
		return 0, errors.New("okcoin can only exchange against USD")
	}

	amnt := strconv.FormatFloat(amount, 'f', -1, 64)

	var values url.Values

	if side == Buy {
		values = url.Values{
			"symbol": {fromCurrency},
			"type":   {marketBuy},
			"price":  {amnt}, // Needed on market buy
			"amount": {},     // Ignore on market buy
		}
	} else if side == Sell {
		values = url.Values{
			"symbol": {fromCurrency},
			"type":   {marketSell},
			"price":  {},     // Ignore on Market sell
			"amount": {amnt}, // Needed on market sell
		}
	} else {
		return 0, errors.New("unknown okex OrderMarket side")
	}

	body, err := internal.Post(urlString, e.getSignedValues(values))
	if err != nil {
		return 0, err
	}

	var resp okexOrderResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return 0, err
	}

	fmt.Println("body:", string(body))
	fmt.Println("resp:", resp)

	return resp.OrderID, nil
}
