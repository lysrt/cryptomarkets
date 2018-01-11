package quoinex

import (
	"fmt"

	"github.com/lysrt/cryptomarkets/entity"
)

func (e *Quoinex) DepositAddress(currency string) (string, error) {
	ccy := entity.NewCurrency(currency)
	switch ccy.Upper() {
	default:
		return "", fmt.Errorf("cannot get %s deposit address", currency)
	}
}
