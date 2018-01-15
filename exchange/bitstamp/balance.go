package bitstamp

import (
	"encoding/json"

	"github.com/lysrt/cryptomarkets/common"
	"github.com/lysrt/cryptomarkets/entity"
)

type balance struct {
	BchAvailable float64 `json:"bch_available,string"`
	BchBalance   float64 `json:"bch_balance,string"`
	BchReserved  float64 `json:"bch_reserved,string"`
	BchbtcFee    float64 `json:"bchbtc_fee,string"`
	BcheurFee    float64 `json:"bcheur_fee,string"`
	BchusdFee    float64 `json:"bchusd_fee,string"`
	BtcAvailable float64 `json:"btc_available,string"`
	BtcBalance   float64 `json:"btc_balance,string"`
	BtcReserved  float64 `json:"btc_reserved,string"`
	BtceurFee    float64 `json:"btceur_fee,string"`
	BtcusdFee    float64 `json:"btcusd_fee,string"`
	EthAvailable float64 `json:"eth_available,string"`
	EthBalance   float64 `json:"eth_balance,string"`
	EthReserved  float64 `json:"eth_reserved,string"`
	EthbtcFee    float64 `json:"ethbtc_fee,string"`
	EtheurFee    float64 `json:"etheur_fee,string"`
	EthusdFee    float64 `json:"ethusd_fee,string"`
	EurAvailable float64 `json:"eur_available,string"`
	EurBalance   float64 `json:"eur_balance,string"`
	EurReserved  float64 `json:"eur_reserved,string"`
	EurusdFee    float64 `json:"eurusd_fee,string"`
	LtcAvailable float64 `json:"ltc_available,string"`
	LtcBalance   float64 `json:"ltc_balance,string"`
	LtcReserved  float64 `json:"ltc_reserved,string"`
	LtcbtcFee    float64 `json:"ltcbtc_fee,string"`
	LtceurFee    float64 `json:"ltceur_fee,string"`
	LtcusdFee    float64 `json:"ltcusd_fee,string"`
	UsdAvailable float64 `json:"usd_available,string"`
	UsdBalance   float64 `json:"usd_balance,string"`
	UsdReserved  float64 `json:"usd_reserved,string"`
	XrpAvailable float64 `json:"xrp_available,string"`
	XrpBalance   float64 `json:"xrp_balance,string"`
	XrpReserved  float64 `json:"xrp_reserved,string"`
	XrpbtcFee    float64 `json:"xrpbtc_fee,string"`
	XrpeurFee    float64 `json:"xrpeur_fee,string"`
	XrpusdFee    float64 `json:"xrpusd_fee,string"`
}

var urlString = "https://www.bitstamp.net/api/v2/balance/"

func (e *Bitstamp) GetBalance() (*entity.Balance, error) {
	body, err := common.Post(urlString, e.getAuthValues())
	if err != nil {
		return nil, err
	}

	var b balance
	err = json.Unmarshal(body, &b)
	if err != nil {
		return nil, err
	}

	balances := entity.Balance{
		"BTC": b.BtcAvailable,
		"ETH": b.EthAvailable,
		"USD": b.UsdAvailable,
		"LTC": b.LtcAvailable,
		"XRP": b.XrpAvailable,
	}

	return &balances, nil
}
