package bitstamp

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/lysrt/cryptomarkets/entity"

	"github.com/lysrt/cryptomarkets/common"
)

// https://www.bitstamp.net/api/v2/withdrawal-requests/

func (e *Bitstamp) DepositAddress(currency string) (string, error) {
	ccy := entity.NewCurrency(currency)
	switch ccy.Upper() {
	case "BTC":
		return e.bitcoinDepositAddress()
	default:
		return "", fmt.Errorf("%s deposit address unimplemented in bitstamp", currency)
	}
}

func (e *Bitstamp) bitcoinDepositAddress() (string, error) {
	url := "https://www.bitstamp.net/api/bitcoin_deposit_address/"

	body, err := common.Post(url, e.getAuthValues())
	if err != nil {
		return "", fmt.Errorf("cannot get bitcoin deposit address: %q", err)
	}

	// Bitstamp can return HTTP Status 200 with a JSON error
	var response errorResponse
	err = json.Unmarshal(body, &response)
	if err == nil {
		return "", fmt.Errorf("cannot get bitcoin deposit address: %q", errors.New(response.Error))
	}

	var address string
	err = json.Unmarshal(body, &address)
	if err != nil {
		return "", fmt.Errorf("cannot get bitcoin deposit address: %q", err)
	}

	return address, nil
}

func (e *Bitstamp) bitcoinWithdrawal(destination, amount string) (string, error) {
	urlString := "https://www.bitstamp.net/api/bitcoin_withdrawal/"

	values := e.getAuthValues()
	values.Add("amount", amount)
	values.Add("destination", destination)
	values.Add("instant", "0")

	body, err := common.Post(urlString, values)
	if err != nil {
		return "", err
	}

	fmt.Println("Body", string(body))

	return "", nil
}
