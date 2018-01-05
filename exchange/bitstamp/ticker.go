package bitstamp

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ticker struct {
	High      float64 `json:"high,string"`
	Last      float64 `json:"last,string"`
	Timestamp int     `json:"timestamp,string"`
	Bid       float64 `json:"bid,string"`
	VWAP      float64 `json:"vwap,string"`
	Volume    float64 `json:"volume,string"`
	Low       float64 `json:"low,string"`
	Ask       float64 `json:"ask,string"`
	Open      float64 `json:"open,string"`
}

/*
	Supported currency pairs (20180105):
	(https://www.bitstamp.net/api/v2/trading-pairs-info/)
	btcusd, btceur,
	eurusd,
	xrpusd, xrpeur, xrpbtc,
	ltcusd, ltceur, ltcbtc,
	ethusd, etheur, ethbtc,
	bchusd, bcheur, bchbtc
*/

// https://www.bitstamp.net/api/v2/ticker_hour/{currency_pair}/

func (e *Bitstamp) LastPrice(from, to string) float64 {
	// https://www.bitstamp.net/api/v2/ticker/{currency_pair}/
	url := "https://www.bitstamp.net/api/v2/ticker/btcusd/"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))

	var ticker ticker

	err = json.Unmarshal(body, &ticker)
	// err = json.NewDecoder(resp.Body).Decode(&ticker)
	if err != nil {
		panic(err)
	}

	// return btcusd by default
	return ticker.Last
}
