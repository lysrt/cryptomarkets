package quoinex

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func (e *Quoinex) GetBalance(currency string) float64 {
	url := "https://api.quoine.com"
	path := "/accounts/balance"
	url += path

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		panic(err)
	}
	signature := e.getSignature()
	req.Header.Add("X-Quoine-Auth", signature)
	req.Header.Add("Content-Type", "application/json")
	// req.Header.Add("X-Quoine-API-Version", "2")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
	panic("Unimplemented")
}
