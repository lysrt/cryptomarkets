package binance

import (
	"errors"
	"fmt"

	"github.com/lysrt/cryptomarkets"
)

func (e *Binance) DepositAddress(currency string) (string, error) {
	ccy := cryptomarkets.NewCurrency(currency)
	switch ccy.Upper() {
	case "BTC":
		return e.bitcoinDepositAddress()
	default:
		return "", fmt.Errorf("cannot get %s deposit address", currency)
	}
}

func (e *Binance) bitcoinDepositAddress() (string, error) {
	return "", errors.New("unimplemented")
}

func (q *Binance) Withdrawal(currency, destination string, amount, fee float64) (int, error) {
	return 0, errors.New("unimplemented")
}

func (e *Binance) bitcoinWithdrawal(destination, amount string) (string, error) {
	return "", errors.New("unimplemented")
}
