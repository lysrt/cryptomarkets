package binance

import (
	"errors"

	"github.com/lysrt/cryptomarkets"
)

func (e *Binance) BuyLimit(from, to string, amount, price float64) (int, error) {
	return 0, errors.New("unimplemented")
}

func (e *Binance) SellLimit(from, to string, amount, price float64) (int, error) {
	return 0, errors.New("unimplemented")
}

func (e *Binance) BuyMarket(from, to string, amount float64) (int, error) {
	return 0, errors.New("unimplemented")
}

func (e *Binance) SellMarket(from, to string, amount float64) (int, error) {
	return 0, errors.New("unimplemented")
}

func (e *Binance) OrderStatus(orderID int) (cryptomarkets.Order, error) {
	return cryptomarkets.Order{}, errors.New("unimplemented")
}

func (e *Binance) CancelOrder(orderID int, from, to string) error {
	return errors.New("unimplemented")
}

func (e *Binance) CancelAllOrders() error {
	return errors.New("unimplemented")
}

func (e *Binance) ListOrders() ([]cryptomarkets.Order, error) {
	return nil, errors.New("unimplemented")
}
