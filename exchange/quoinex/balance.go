package quoinex

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/lysrt/cryptomarkets"
)

type balance struct {
	Currency string      `json:"currency"`
	Balance  json.Number `json:"balance"`
}

func (e *Quoinex) GetBalance() (*cryptomarkets.Balance, error) {
	url := "https://api.quoine.com"
	path := "/accounts/balance"
	// path := "/crypto_accounts" // More details
	url += path

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		panic(err)
	}

	signature := e.getSignature(path)
	req.Header.Add("X-Quoine-Auth", signature)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Quoine-API-Version", "2")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad http request: %s. %s", resp.Status, string(body))
	}

	var balances []balance

	json.Unmarshal(body, &balances)

	bmap := make(cryptomarkets.Balance)
	for _, b := range balances {
		v, err := b.Balance.Float64()
		if err != nil {
			continue
		}
		bmap[cryptomarkets.NewCurrency(b.Currency)] = v
	}

	return &bmap, nil
}
