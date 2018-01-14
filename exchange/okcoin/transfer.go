package okcoin

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/lysrt/cryptomarkets/common"
)

func (e *Okcoin) DepositAddress(currency string) (string, error) {
	return "", errors.New("unavailable in okcoin, check deposit address in your account")
}

type targetType string

const (
	okcoinCN     targetType = "okcn"
	okcoinCOM               = "okcom"
	okesCOM                 = "okex"
	outerAddress            = "address"
)

type okcoinWithdrawalResponse struct {
	Result    bool `json:"result"`
	ErrorCode int  `json:"error_code"`
}

func (e *Okcoin) Withdrawal(currency, destination string, amount float64) error {
	urlString := "https://www.okcoin.com/api/v1/withdraw.do"

	var symbol string
	switch strings.ToLower(currency) {
	case "btc":
		symbol = "btc_usd"
	case "ltc":
		symbol = "ltc_usd"
	case "eth":
		symbol = "eth_usd"
	case "etc":
		symbol = "etc_usd"
	case "bch":
		symbol = "bch_usd"
	default:
		return fmt.Errorf("okcoin cannot withdraw %s, only: btc, ltc, eth, etc, bch", currency)
	}

	values := url.Values{
		"symbol":           {symbol},
		"chargefee":        {"0.002"},
		"trade_pwd":        {e.CustomerID},
		"withdraw_address": {destination},
		"withdraw_amount":  {strconv.FormatFloat(amount, 'f', -1, 64)},
		// "target":           {outerAddress}, // Not mandatory
	}

	body, err := common.Post(urlString, e.getSignedValues(values))
	if err != nil {
		return err
	}

	var resp okcoinWithdrawalResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return err
	}

	if resp.Result == false {
		return fmt.Errorf("okcoin API error (%d): %s", resp.ErrorCode, errorCodes[resp.ErrorCode])
	}

	fmt.Println(string(body))
	fmt.Println(resp)

	return nil
}
