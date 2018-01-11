package okcoin

import (
	"errors"
)

func (e *Okcoin) DepositAddress(currency string) (string, error) {
	return "", errors.New("unavailable in okcoin, check deposit address in your account")
}
