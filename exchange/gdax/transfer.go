package gdax

import (
	"errors"
	"fmt"

	"github.com/lysrt/cryptomarkets"
)

func (e *Gdax) DepositAddress(currency string) (string, error) {
	ccy := cryptomarkets.NewCurrency(currency)
	switch ccy.Upper() {
	case "BTC":
		return e.bitcoinDepositAddress()
	default:
		return "", fmt.Errorf("cannot get %s deposit address", currency)
	}
}

func (e *Gdax) bitcoinDepositAddress() (string, error) {
	return "", errors.New("unimplemented")
}

func (q *Gdax) Withdrawal(currency, destination string, amount float64) (int, error) {
	return 0, errors.New("unimplemented")
}

func (e *Gdax) bitcoinWithdrawal(destination, amount string) (string, error) {
	return "", errors.New("unimplemented")
}
