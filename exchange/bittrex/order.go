package bittrex

import (
	"errors"

	"github.com/lysrt/cryptomarkets"
)

func (e *Bittrex) BuyLimit(from, to string, amount, price float64) (int, error) {
	return 0, errors.New("unimplemented")
}

func (e *Bittrex) BuyMarket(from, to string, amount float64) (int, error) {
	return 0, errors.New("unimplemented")
}

func (e *Bittrex) SellLimit(from, to string, amount, price float64) (int, error) {
	return 0, errors.New("unimplemented")
}

func (e *Bittrex) SellMarket(from, to string, amount float64) (int, error) {
	return 0, errors.New("unimplemented")
}

func (e *Bittrex) OrderStatus(orderID int) (cryptomarkets.Order, error) {
	return cryptomarkets.Order{}, errors.New("unimplemented")
}

func (e *Bittrex) CancelOrder(orderID int, from, to string) error {
	return errors.New("unimplemented")
}

func (e *Bittrex) CancelAllOrders() error {
	return errors.New("unimplemented")
}

func (e *Bittrex) ListOrders() ([]cryptomarkets.Order, error) {
	return nil, errors.New("unimplemented")
}
