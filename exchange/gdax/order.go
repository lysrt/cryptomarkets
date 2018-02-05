package gdax

import (
	"errors"

	"github.com/lysrt/cryptomarkets"
)

func (e *Gdax) BuyLimit(from, to string, amount, price float64) (int, error) {
	return 0, errors.New("unimplemented")
}

func (e *Gdax) SellLimit(from, to string, amount, price float64) (int, error) {
	return 0, errors.New("unimplemented")
}

func (e *Gdax) BuyMarket(from, to string, amount float64) (int, error) {
	return 0, errors.New("unimplemented")
}

func (e *Gdax) SellMarket(from, to string, amount float64) (int, error) {
	return 0, errors.New("unimplemented")
}

func (e *Gdax) OrderStatus(orderID int) (cryptomarkets.Order, error) {
	return cryptomarkets.Order{}, errors.New("unimplemented")
}

func (e *Gdax) CancelOrder(orderID int) error {
	return errors.New("unimplemented")
}

func (e *Gdax) CancelAllOrders() error {
	return errors.New("unimplemented")
}

func (e *Gdax) ListOrders() ([]cryptomarkets.Order, error) {
	return nil, errors.New("unimplemented")
}
