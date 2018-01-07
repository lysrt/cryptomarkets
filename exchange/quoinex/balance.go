package quoinex

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type balance struct {
	Currency string      `json:"currency"`
	Balance  json.Number `json:"balance"`
}

func (e *Quoinex) GetBalance(currency string) float64 {
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

	// fmt.Println(resp)
	// for k, v := range resp.Header {
	// 	fmt.Println(k, v)
	// }

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// fmt.Println("Body:", string(body))

	var balances []balance

	json.Unmarshal(body, &balances)
	for _, b := range balances {
		fmt.Println(b.Currency, b.Balance)
	}

	return 0.0
}
