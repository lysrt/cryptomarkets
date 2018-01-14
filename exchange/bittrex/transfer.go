package bittrex

import (
	"errors"
	"fmt"

	"github.com/lysrt/cryptomarkets/entity"
)

func (e *Bittrex) DepositAddress(currency string) (string, error) {
	ccy := entity.NewCurrency(currency)
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

func (q *Bittrex) Withdrawal(currency, destination string, amount float64) error {
	return errors.New("unimplemented")
}

func (e *Bittrex) bitcoinWithdrawal(destination, amount string) (string, error) {
	return "", errors.New("unimplemented")
}
