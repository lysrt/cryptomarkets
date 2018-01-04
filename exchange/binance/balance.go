package binance

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

var urlString = "https://api.binance.com/api/v3/account"

type account struct {
	MakerCommission  int  `json:"makerCommission"`
	TakerCommission  int  `json:"takerCommission"`
	BuyerCommission  int  `json:"buyerCommission"`
	SellerCommission int  `json:"sellerCommission"`
	CanTrade         bool `json:"canTrade"`
	CanWithdraw      bool `json:"canWithdraw"`
	CanDeposit       bool `json:"canDeposit"`
	Balances         []struct {
		Asset  string  `json:"asset"`
		Free   float64 `json:"free,string"`
		Locked float64 `json:"locked,string"`
	} `json:"balances"`
}

func (e *Binance) GetBalance(currency string) float64 {
	// timestamp is mandatory when signing requests
	timestamp := fmt.Sprintf("%d000", time.Now().Unix()) // Need milliseconds
	values := url.Values{
		"timestamp": {timestamp},
	}

	req := e.getSignedRequest(values)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var a account
	err = json.NewDecoder(resp.Body).Decode(&a)
	if err != nil {
		panic(err)
	}

	// TODO Make a dict of all balances

	// BTC by default
	var btcBalance float64
	for _, b := range a.Balances {
		if b.Asset == "BTC" {
			btcBalance = b.Free
		}
	}

	return btcBalance
}
