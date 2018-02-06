package bittrex

import (
	"errors"
	"fmt"

	"github.com/lysrt/cryptomarkets"
)

func (e *Bittrex) DepositAddress(currency string) (string, error) {
	ccy := cryptomarkets.NewCurrency(currency)
	switch ccy.Upper() {
	case "BTC":
		return e.bitcoinDepositAddress()
	default:
		return "", fmt.Errorf("cannot get %s deposit address", currency)
	}
}

func (e *Bittrex) bitcoinDepositAddress() (string, error) {
	return "", errors.New("unimplemented")
}

func (q *Bittrex) Withdrawal(currency, destination string, amount, fee float64) (int, error) {
	return 0, errors.New("unimplemented")
}

func (e *Bittrex) bitcoinWithdrawal(destination, amount string) (string, error) {
	return "", errors.New("unimplemented")
}
