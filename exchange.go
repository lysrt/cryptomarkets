package cryptomarkets

import "github.com/lysrt/cryptomarkets/entity"

type Exchange interface {
	GetTicker(from, to string) (*entity.Ticker, error)
	OrderBook(from, to string) (*entity.OrderBook, error)
}

type PrivateExchange interface {
	GetBalance() (*entity.Balance, error)
	DepositAddress(currency string) (string, error)
	Withdrawal(currency, destination string, amount float64) (int, error)
}

type ExchangeConfig struct {
	Name       string
	ApiKey     string
	Secret     string
	CustomerID string
}
