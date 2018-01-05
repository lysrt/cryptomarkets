package bitstamp

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// https://www.bitstamp.net/api/v2/withdrawal-requests/

func (e *Bitstamp) BitcoinWithdrawal(destination, value string) string {
	urlString := "https://www.bitstamp.net/api/bitcoin_withdrawal/"

	values := e.getAuthValues()
	values.Add("amount", value)
	values.Add("address", destination)
	values.Add("instant", "0")

	resp, err := http.PostForm(urlString, values)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// TODO Add better error handling to all requests
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Body", string(body))
	return ""
}

func (e *Bitstamp) BitcoinDepositAddress() string {
	urlString := "https://www.bitstamp.net/api/bitcoin_deposit_address/"

	resp, err := http.PostForm(urlString, e.getAuthValues())
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var address string

	d := json.NewDecoder(resp.Body)
	err = d.Decode(&address)
	if err != nil {
		panic(err)
	}

	return address
}
