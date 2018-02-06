package quoinex

import (
	"errors"

	"github.com/lysrt/cryptomarkets"
)

func (e *Quoinex) BuyLimit(from, to string, amount, price float64) (int, error) {
	return 0, errors.New("unimplemented")
}

func (e *Quoinex) SellLimit(from, to string, amount, price float64) (int, error) {
	return 0, errors.New("unimplemented")
}

func (e *Quoinex) BuyMarket(from, to string, amount float64) (int, error) {
	return 0, errors.New("unimplemented")
}

func (e *Quoinex) SellMarket(from, to string, amount float64) (int, error) {
	return 0, errors.New("unimplemented")
}

func (e *Quoinex) OrderStatus(orderID int) (cryptomarkets.Order, error) {
	return cryptomarkets.Order{}, errors.New("unimplemented")
}

func (e *Quoinex) CancelOrder(orderID int, from, to string) error {
	return errors.New("unimplemented")
}

func (e *Quoinex) CancelAllOrders() error {
	return errors.New("unimplemented")
}

func (e *Quoinex) ListOrders() ([]cryptomarkets.Order, error) {
	return nil, errors.New("unimplemented")
}
