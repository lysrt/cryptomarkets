package okcoin

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/lysrt/cryptomarkets/common"
	"github.com/lysrt/cryptomarkets/entity"
)

type okcoinBalanceResponse struct {
	Result    bool          `json:"result"`
	Info      okcoinBalance `json:"info"`
	ErrorCode int           `json:"error_code"`
}

type okcoinBalance struct {
	Funds struct {
		Borrow okCoinBalanceEntry `json:"borrow"`
		Asset  struct {
			Total float64 `json:"total,string"`
			Net   float64 `json:"net,string"`
		} `json:"asset"`
		Free    okCoinBalanceEntry `json:"free"`
		Freezed okCoinBalanceEntry `json:"freezed"`
	} `json:"funds"`
}

type okCoinBalanceEntry struct {
	BTC float64 `json:"btc,string"`
	BCC float64 `json:"bcc,string"`
	ETC float64 `json:"etc,string"`
	BCH float64 `json:"bch,string"`
	USD float64 `json:"usd,string"`
	ETH float64 `json:"eth,string"`
	LTC float64 `json:"ltc,string"`
}

func (e *Okcoin) GetBalance() (*entity.Balance, error) {
	urlString := "https://www.okcoin.com/api/v1/userinfo.do"

	body, err := common.Post(urlString, e.getSignedValues(url.Values{}))
	if err != nil {
		return nil, fmt.Errorf("bad HTTP response: %q", err.Error())
	}

	var resp okcoinBalanceResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}

	if resp.ErrorCode != 0 {
		return nil, fmt.Errorf("okcoin API error code: %d", resp.ErrorCode)
	}

	balances := entity.Balance{
		"BTC": resp.Info.Funds.Free.BTC,
		"BCC": resp.Info.Funds.Free.BCC,
		"BCH": resp.Info.Funds.Free.BCH,
		"ETC": resp.Info.Funds.Free.ETC,
		"ETH": resp.Info.Funds.Free.ETH,
		"LTC": resp.Info.Funds.Free.LTC,
		"USD": resp.Info.Funds.Free.USD,
	}

	return &balances, nil
}
