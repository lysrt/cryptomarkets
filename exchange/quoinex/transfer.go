package quoinex

import (
	"errors"
	"fmt"

	"github.com/lysrt/cryptomarkets"
)

func (e *Quoinex) DepositAddress(currency string) (string, error) {
	ccy := cryptomarkets.NewCurrency(currency)
	switch ccy.Upper() {
	default:
		return "", fmt.Errorf("cannot get %s deposit address", currency)
	}
}

func (q *Quoinex) Withdrawal(currency, destination string, amount float64) (int, error) {
	return 0, errors.New("unimplemented")
}
