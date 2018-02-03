package cryptomarkets

import (
	"strings"
)

// Exchange serves as a utility interface, whose methods are implemented by all exchanges
type Exchange interface {
	GetTicker(from, to string) (*Ticker, error)
	OrderBook(from, to string) (*OrderBook, error)
	GetBalance() (*Balance, error)
	DepositAddress(currency string) (string, error)
	Withdrawal(currency, destination string, amount float64) (int, error)
}

// Ticker holds the common ticker information of an exchange price
type Ticker struct {
	Timestamp     int64
	LastPrice     float64
	LastQuantity  float64
	High          float64
	Low           float64
	Open          float64
	Close         float64
	Ask           float64
	AskQuantity   float64
	Bid           float64
	BidQuantity   float64
	VWAP          float64
	Volume        float64
	QuoteVolume   float64
	PriceChange   float64
	PercentChange float64
	Pair          Pair
}

// OrderBook holds the open orders of an exchange
type OrderBook struct {
	Timestamp int64
	Asks      []Order
	Bids      []Order
}

// Order represents en entry int the order book
type Order struct {
	Price    float64
	Quantity float64
}

type OrderType string

const (
	Limit  OrderType = "limit"
	Market OrderType = "market"
)

// Balance holds the list of the user balances on an exchange
type Balance map[Currency]float64

// Currency represents a currency, usable by all exchanges
type Currency string

func NewCurrency(currency string) Currency {
	return Currency(strings.ToUpper(currency))
}

func (c Currency) Lower() string {
	return strings.ToLower(string(c))
}

func (c Currency) Upper() string {
	return string(c)
}

type Pair struct {
	First, Second Currency
}

func (p Pair) Lower(sep string) string {
	return p.First.Lower() + sep + p.Second.Lower()
}

func (p Pair) Upper(sep string) string {
	return p.First.Upper() + sep + p.Second.Upper()
}
