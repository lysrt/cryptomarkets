package okex

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/lysrt/cryptomarkets/internal"
)

func (e *Okex) DepositAddress(currency string) (string, error) {
	return "", errors.New("unavailable in okex, check deposit address in your account")
}

type targetType string

type okexWithdrawalResponse struct {
	Result     bool `json:"result"`
	ErrorCode  int  `json:"error_code"`
	WithdrawID int  `json:"withdraw_id"`
}

func (e *Okex) Withdrawal(currency, destination string, amount, fee float64) (int, error) {
	urlString := "https://www.okex.com/api/v1/withdraw.do"
	/*
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
			return 0, fmt.Errorf("okex cannot withdraw %s, only: btc, ltc, eth, etc, bch", currency)
		}
	*/

	symbol := strings.ToLower(currency + "_usd")

	values := url.Values{
		"symbol":           {symbol},
		"chargefee":        {strconv.FormatFloat(fee, 'f', -1, 64)},
		"trade_pwd":        {e.CustomerID},
		"withdraw_address": {destination},
		"withdraw_amount":  {strconv.FormatFloat(amount, 'f', -1, 64)},
		// "target":           {"address"}, // Not mandatory
	}

	body, err := internal.Post(urlString, e.getSignedValues(values))
	if err != nil {
		return 0, fmt.Errorf("okex request error: %q", err)
	}

	var resp okexWithdrawalResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return 0, err
	}

	if resp.Result == false {
		return 0, fmt.Errorf("okex API error (%d): %s", resp.ErrorCode, errorCodes[resp.ErrorCode])
	}

	return resp.WithdrawID, nil
}
