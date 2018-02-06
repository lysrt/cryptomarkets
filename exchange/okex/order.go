package okex

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/lysrt/cryptomarkets"
	"github.com/lysrt/cryptomarkets/internal"
)

type orderType string

const (
	limitBuy   orderType = "buy"
	limitSell  orderType = "sell"
	marketBuy  orderType = "buy_market"
	marketSell orderType = "sell_market"
)

/*
	ltc_btc
	eth_btc
	etc_btc
	bch_btc
	btc_usdt
	eth_usdt
	ltc_usdt
	etc_usdt
	bch_usdt
	etc_eth
	bt1_btc
	bt2_btc
	btg_btc
	qtum_btc
	hsr_btc
	neo_btc
	gas_btc
	qtum_usdt
	hsr_usdt
	neo_usdt
	gas_usdt
*/
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
	buy  Side = "buy"
	sell Side = "sell"
)

type okexOrderResponse struct {
	Result    bool `json:"result"`
	OrderID   int  `json:"order_id"`
	ErrorCode int  `json:"error_code"`
}

func (e *Okex) BuyLimit(from, to string, amount, price float64) (int, error) {
	return e.limitOrder(from, to, buy, amount, price)
}

func (e *Okex) SellLimit(from, to string, amount, price float64) (int, error) {
	return e.limitOrder(from, to, sell, amount, price)
}

func (e *Okex) limitOrder(from, to string, side Side, amount, price float64) (int, error) {
	pair := cryptomarkets.Pair{
		First:  cryptomarkets.NewCurrency(from),
		Second: cryptomarkets.NewCurrency(to),
	}

	amnt := strconv.FormatFloat(amount, 'f', -1, 64)
	prce := strconv.FormatFloat(price, 'f', -1, 64)

	var values url.Values

	if side == buy {
		values = url.Values{
			"symbol": {pair.Lower("_")},
			"type":   {string(limitBuy)},
			"price":  {prce},
			"amount": {amnt},
		}
	} else if side == sell {
		values = url.Values{
			"symbol": {pair.Lower("_")},
			"type":   {string(limitSell)},
			"price":  {prce},
			"amount": {amnt},
		}
	} else {
		return 0, errors.New("unknown okex OrderMarket side")
	}

	urlString := "https://www.okex.com/api/v1/trade.do"

	body, err := internal.Post(urlString, e.getSignedValues(values))
	if err != nil {
		return 0, err
	}

	var resp okexOrderResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return 0, err
	}

	if resp.ErrorCode != 0 {
		return 0, fmt.Errorf("okex API error (%d): %s", resp.ErrorCode, errorCodes[resp.ErrorCode])
	}

	if !resp.Result {
		return 0, errors.New("okex cannot execute order (no reason given)")
	}

	return resp.OrderID, nil
}

func (e *Okex) BuyMarket(from, to string, amount float64) (int, error) {
	return e.marketOrder(from, to, buy, amount)
}

func (e *Okex) SellMarket(from, to string, amount float64) (int, error) {
	return e.marketOrder(from, to, sell, amount)
}

func (e *Okex) marketOrder(from, to string, side Side, amount float64) (int, error) {
	pair := cryptomarkets.Pair{
		First:  cryptomarkets.NewCurrency(from),
		Second: cryptomarkets.NewCurrency(to),
	}

	amnt := strconv.FormatFloat(amount, 'f', -1, 64)

	var values url.Values

	if side == buy {
		values = url.Values{
			"symbol": {pair.Lower("_")},
			"type":   {string(marketBuy)},
			"price":  {amnt}, // Needed on market buy
			"amount": {},     // Ignore on market buy
		}
	} else if side == sell {
		values = url.Values{
			"symbol": {pair.Lower("_")},
			"type":   {string(marketSell)},
			"price":  {},     // Ignore on Market sell
			"amount": {amnt}, // Needed on market sell
		}
	} else {
		return 0, errors.New("unknown okex OrderMarket side")
	}

	urlString := "https://www.okex.com/api/v1/trade.do"

	body, err := internal.Post(urlString, e.getSignedValues(values))
	if err != nil {
		return 0, err
	}

	var resp okexOrderResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return 0, err
	}

	if resp.ErrorCode != 0 {
		return 0, fmt.Errorf("okex API error (%d): %s", resp.ErrorCode, errorCodes[resp.ErrorCode])
	}

	if !resp.Result {
		return 0, errors.New("okex cannot execute order (no reason given)")
	}

	return resp.OrderID, nil
}

func (e *Okex) OrderStatus(orderID int) (cryptomarkets.Order, error) {
	return cryptomarkets.Order{}, errors.New("unimplemented")
}

type okexCancelResponse struct {
	Result    bool `json:"result"`
	OrderID   int  `json:"order_id"`
	ErrorCode int  `json:"error_code"`
}

func (e *Okex) CancelOrder(orderID int, from, to string) error {
	pair := cryptomarkets.Pair{
		First:  cryptomarkets.NewCurrency(from),
		Second: cryptomarkets.NewCurrency(to),
	}

	urlString := "https://www.okex.com/api/v1/cancel_order.do"

	body, err := internal.Post(urlString, e.getSignedValues(url.Values{
		"order_id": {strconv.Itoa(orderID)},
		"symbol":   {pair.Lower("_")},
	}))
	if err != nil {
		return err
	}

	var resp okexCancelResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return err
	}

	if resp.ErrorCode != 0 {
		return fmt.Errorf("okex API error (%d): %s", resp.ErrorCode, errorCodes[resp.ErrorCode])
	}

	return errors.New("unimplemented")
}

func (e *Okex) CancelAllOrders() error {
	return errors.New("unimplemented")
}

type okexListResponse struct {
	Result    bool            `json:"result"`
	ErrorCode int             `json:"error_code"`
	Orders    []okexListOrder `json:"orders"`
}

type okexListOrder struct {
	Amount     float64 `json:"amount"`
	AvgPrice   float64 `json:"avg_price"`
	CreateDate int     `json:"create_date"`
	DealAmount float64 `json:"deal_amount"`
	OrderID    int     `json:"order_id"`
	// orders_id int `json:"orders_id"` // Deprecated
	Price  float64   `json:"price"`
	Status int       `json:"status"`
	Symbol string    `json:"symbol"`
	Type   orderType `json:"type"`
}

func (e *Okex) ListOrders() ([]cryptomarkets.Order, error) {
	urlString := "https://www.okex.com/api/v1/order_info.do"

	body, err := internal.Post(urlString, e.getSignedValues(url.Values{
		"symbol":   {"ltc_btc"},
		"order_id": {"-1"}, // All unfilled orders
	}))
	if err != nil {
		return nil, err
	}

	var resp okexListResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}

	if resp.ErrorCode != 0 {
		return nil, fmt.Errorf("okex API error (%d): %s", resp.ErrorCode, errorCodes[resp.ErrorCode])
	}

	orders := []cryptomarkets.Order{}

	for _, o := range resp.Orders {
		p := strings.Split(o.Symbol, "_")

		order := cryptomarkets.Order{
			ID: o.OrderID,
			Pair: cryptomarkets.Pair{
				First:  cryptomarkets.NewCurrency(p[0]),
				Second: cryptomarkets.NewCurrency(p[1]),
			},
			Amount:       o.Amount,
			Price:        o.Price,
			Type:         cryptomarkets.OrderType(o.Type),
			Status:       cryptomarkets.OrderStatus(o.Status),
			CreationTime: o.CreateDate,
			AveragePrice: o.AvgPrice,
			DealAmount:   o.DealAmount,
		}
		orders = append(orders, order)
	}

	return orders, nil
}
