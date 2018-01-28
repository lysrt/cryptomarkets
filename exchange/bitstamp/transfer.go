package bitstamp

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/lysrt/cryptomarkets"

	"github.com/lysrt/cryptomarkets/internal"
)

// https://www.bitstamp.net/api/v2/withdrawal-requests/

func (e *Bitstamp) DepositAddress(currency string) (string, error) {
	ccy := cryptomarkets.NewCurrency(currency)
	switch ccy.Upper() {
	case "BTC":
		return e.bitcoinDepositAddress()
	case "ETH":
		return e.ethereumDepositAddress()
	case "LTC":
		return e.litecoinDepositAddress()
	default:
		return "", fmt.Errorf("%s deposit address unimplemented in bitstamp", currency)
	}
}

func (e *Bitstamp) bitcoinDepositAddress() (string, error) {
	url := "https://www.bitstamp.net/api/bitcoin_deposit_address/"

	body, err := internal.Post(url, e.getAuthValues())
	if err != nil {
		return "", fmt.Errorf("cannot get Bitstamp Bitcoin deposit address: %q", err)
	}

	// Bitstamp can return HTTP Status 200 with a JSON error
	var response errorResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return "", fmt.Errorf("cannot get Bitstamp Bitcoin deposit address: %q", errors.New(response.Error))
	}

	var address string
	err = json.Unmarshal(body, &address)
	if err != nil {
		return "", fmt.Errorf("cannot get Bitstamp Bitcoin deposit address: %q", err)
	}

	return address, nil
}

func (e *Bitstamp) ethereumDepositAddress() (string, error) {
	url := "https://www.bitstamp.net/api/v2/eth_address/"

	body, err := internal.Post(url, e.getAuthValues())
	if err != nil {
		return "", fmt.Errorf("cannot get Bitstamp Ethereum deposit address: %q", err)
	}

	// Bitstamp can return HTTP Status 200 with a JSON error
	var response errorResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return "", fmt.Errorf("cannot get Bitstamp Ethereum deposit address: %q", errors.New(response.Error))
	}

	var address struct {
		Address string `json:"address"`
	}
	err = json.Unmarshal(body, &address)
	if err != nil {
		return "", fmt.Errorf("cannot get Bitstamp Ethereum deposit address: %q", err)
	}

	return address.Address, nil
}

func (e *Bitstamp) litecoinDepositAddress() (string, error) {
	url := "https://www.bitstamp.net/api/v2/ltc_address/"

	body, err := internal.Post(url, e.getAuthValues())
	if err != nil {
		return "", fmt.Errorf("cannot get Bitstamp Litecoin deposit address: %q", err)
	}

	// Bitstamp can return HTTP Status 200 with a JSON error
	var response errorResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return "", fmt.Errorf("cannot get Bitstamp Litecoin deposit address: %q", errors.New(response.Error))
	}

	var address struct {
		Address string `json:"address"`
	}
	err = json.Unmarshal(body, &address)
	if err != nil {
		return "", fmt.Errorf("cannot get Bitstamp Litecoin deposit address: %q", err)
	}

	return address.Address, nil
}

func (e *Bitstamp) Withdrawal(currency, destination string, amount float64) (int, error) {
	ccy := cryptomarkets.NewCurrency(currency)
	switch ccy.Upper() {
	case "BTC":
		return e.bitcoinWithdrawal(destination, strconv.FormatFloat(amount, 'f', -1, 64))
	case "LTC":
		return e.litecoinWithdrawal(destination, strconv.FormatFloat(amount, 'f', -1, 64))
	default:
		return 0, fmt.Errorf("%s withdrawal unimplemented in bitstamp", currency)
	}
}

func (e *Bitstamp) bitcoinWithdrawal(destination, amount string) (int, error) {
	urlString := "https://www.bitstamp.net/api/bitcoin_withdrawal/"

	values := e.getAuthValues()
	values.Add("amount", amount)
	values.Add("address", destination)
	values.Add("instant", "0")

	body, err := internal.Post(urlString, values)
	if err != nil {
		return 0, err
	}

	var resp struct {
		WithdrawalID int    `json:"id"`
		Status       string `json:"status"`
	}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return 0, err
	}

	if resp.Status == "error" {
		return 0, errors.New(string(body)) // TODO Make this better
	}

	return resp.WithdrawalID, nil
}

func (e *Bitstamp) litecoinWithdrawal(destination, amount string) (int, error) {
	urlString := "https://www.bitstamp.net/api/v2/ltc_withdrawal/"

	values := e.getAuthValues()
	values.Add("amount", amount)
	values.Add("address", destination)
	values.Add("instant", "0")

	body, err := internal.Post(urlString, values)
	if err != nil {
		return 0, err
	}

	var resp struct {
		WithdrawalID int    `json:"id"`
		Status       string `json:"status"`
	}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return 0, err
	}

	if resp.Status == "error" {
		return 0, errors.New(string(body)) // TODO Make this better
	}

	return resp.WithdrawalID, nil
}
