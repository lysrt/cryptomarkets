package bitstamp

import (
	"fmt"
	"strconv"

	"github.com/lysrt/cryptomarkets/common"
	"github.com/lysrt/cryptomarkets/entity"
)

func (e *Bitstamp) BuyMarketOrder(from, to string, amount float64) (int, error) {
	pair := entity.Pair{
		First:  entity.NewCurrency(from),
		Second: entity.NewCurrency(to),
	}
	urlString := "https://www.bitstamp.net/api/v2/buy/market/%s/"
	urlString = fmt.Sprintf(urlString, pair.Lower(""))

	values := e.getAuthValues()
	values.Add("amount", strconv.FormatFloat(amount, 'f', -1, 64))
	body, err := common.Post(urlString, values)
	if err != nil {
		return 0, err
	}

	fmt.Println("body", string(body))

	return 0, nil
}

func (e *Bitstamp) SellMarketOrder(from, to string, amount float64) (int, error) {
	pair := entity.Pair{
		First:  entity.NewCurrency(from),
		Second: entity.NewCurrency(to),
	}
	urlString := "https://www.bitstamp.net/api/v2/sell/market/%s/"
	urlString = fmt.Sprintf(urlString, pair.Lower(""))

	values := e.getAuthValues()
	values.Add("amount", strconv.FormatFloat(amount, 'f', -1, 64))
	body, err := common.Post(urlString, values)
	if err != nil {
		return 0, err
	}

	fmt.Println("body", string(body))

	// body {"price": "0.01701001", "amount": "0.56600000", "type": "1", "id": "771366314", "datetime": "2018-01-15 22:19:27.670222"}

	// TODO Return correct value

	return 0, nil
}
