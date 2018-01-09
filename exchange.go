package cryptomarkets

import "github.com/lysrt/cryptomarkets/entity"

type Exchange interface {
	GetTicker(from, to string) (*entity.Ticker, error)
	GetBalance() (*entity.Balance, error)
	// DepositAddress(currency string) (string, error)
	// Withdrawal(currency, destination string, value float64) error
}

type ExchangeConfig struct {
	Name       string
	ApiKey     string
	Secret     string
	CustomerID string
}
