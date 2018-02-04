package bitstamp

import (
	"fmt"
	"strconv"

	"github.com/lysrt/cryptomarkets"
	"github.com/lysrt/cryptomarkets/internal"
)

func (e *Bitstamp) BuyLimit(from, to string, amount, price float64) (int, error) {
	pair := cryptomarkets.Pair{
		First:  cryptomarkets.NewCurrency(from),
		Second: cryptomarkets.NewCurrency(to),
	}
	urlString := "https://www.bitstamp.net/api/v2/buy/%s/"
	urlString = fmt.Sprintf(urlString, pair.Lower(""))

	values := e.getAuthValues()
	values.Add("amount", strconv.FormatFloat(amount, 'f', -1, 64))
	values.Add("price", strconv.FormatFloat(price, 'f', -1, 64))

	body, err := internal.Post(urlString, values)
	if err != nil {
		return 0, err
	}

	fmt.Println("body", string(body))

	return 0, nil
}

func (e *Bitstamp) BuyMarket(from, to string, amount float64) (int, error) {
	pair := cryptomarkets.Pair{
		First:  cryptomarkets.NewCurrency(from),
		Second: cryptomarkets.NewCurrency(to),
	}
	urlString := "https://www.bitstamp.net/api/v2/buy/market/%s/"
	urlString = fmt.Sprintf(urlString, pair.Lower(""))

	values := e.getAuthValues()
	values.Add("amount", strconv.FormatFloat(amount, 'f', -1, 64))
	body, err := internal.Post(urlString, values)
	if err != nil {
		return 0, err
	}

	fmt.Println("body", string(body))

	return 0, nil
}

func (e *Bitstamp) SellLimit(from, to string, amount, price float64) (int, error) {
	pair := cryptomarkets.Pair{
		First:  cryptomarkets.NewCurrency(from),
		Second: cryptomarkets.NewCurrency(to),
	}
	urlString := "https://www.bitstamp.net/api/v2/sell/%s/"
	urlString = fmt.Sprintf(urlString, pair.Lower(""))

	values := e.getAuthValues()
	values.Add("amount", strconv.FormatFloat(amount, 'f', -1, 64))
	values.Add("price", strconv.FormatFloat(price, 'f', -1, 64))
	body, err := internal.Post(urlString, values)
	if err != nil {
		return 0, err
	}

	fmt.Println("body", string(body))

	return 0, nil
}

func (e *Bitstamp) SellMarket(from, to string, amount float64) (int, error) {
	pair := cryptomarkets.Pair{
		First:  cryptomarkets.NewCurrency(from),
		Second: cryptomarkets.NewCurrency(to),
	}
	urlString := "https://www.bitstamp.net/api/v2/sell/market/%s/"
	urlString = fmt.Sprintf(urlString, pair.Lower(""))

	values := e.getAuthValues()
	values.Add("amount", strconv.FormatFloat(amount, 'f', -1, 64))
	body, err := internal.Post(urlString, values)
	if err != nil {
		return 0, err
	}

	fmt.Println("body", string(body))

	// body {"price": "0.01701001", "amount": "0.56600000", "type": "1", "id": "771366314", "datetime": "2018-01-15 22:19:27.670222"}

	// TODO Return correct value

	return 0, nil
}

func (e *Bitstamp) OrderStatus(orderID int) error {
	urlString := "https://www.bitstamp.net/api/order_status/"

	values := e.getAuthValues()
	values.Add("id", strconv.Itoa(orderID))
	body, err := internal.Post(urlString, values)
	if err != nil {
		return err
	}

	err = e.checkResponse(body)
	if err != nil {
		return fmt.Errorf("order status error: %q", err)
	}

	fmt.Println("orderStatus: ", string(body))

	return nil
}

func (e *Bitstamp) CancelOrder(orderID int) error {
	urlString := "https://www.bitstamp.net/api/v2/cancel_order/"

	values := e.getAuthValues()
	values.Add("id", strconv.Itoa(orderID))
	body, err := internal.Post(urlString, values)
	if err != nil {
		return err
	}

	err = e.checkResponse(body)
	if err != nil {
		return fmt.Errorf("cancel order error: %q", err)
	}

	fmt.Println("cancellOrder: ", string(body))

	return nil
}

func (e *Bitstamp) CancelAllOrders() error {
	urlString := "https://www.bitstamp.net/api/cancel_all_orders/"

	values := e.getAuthValues()
	body, err := internal.Post(urlString, values)
	if err != nil {
		return err
	}

	err = e.checkResponse(body)
	if err != nil {
		return fmt.Errorf("cancel all orders error: %q", err)
	}

	fmt.Println("cancellAllOrders: ", string(body))

	return nil
}

func (e *Bitstamp) ListOrders() {

}
