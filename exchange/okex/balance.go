package okex

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/lysrt/cryptomarkets/common"
	"github.com/lysrt/cryptomarkets/entity"
)

type okexBalanceResponse struct {
	Result    bool        `json:"result"`
	Info      okexBalance `json:"info"`
	ErrorCode int         `json:"error_code"`
}

type okexBalance struct {
	Funds struct {
		Borrow okexBalanceEntry `json:"borrow"`
		Asset  struct {
			Total float64 `json:"total,string"`
			Net   float64 `json:"net,string"`
		} `json:"asset"`
		Free    okexBalanceEntry `json:"free"`
		Freezed okexBalanceEntry `json:"freezed"`
	} `json:"funds"`
}

type okexBalanceEntry struct {
	// OKcoin
	BTC float64 `json:"btc,string"`
	BCC float64 `json:"bcc,string"`
	ETC float64 `json:"etc,string"`
	BCH float64 `json:"bch,string"`
	USD float64 `json:"usd,string"`
	ETH float64 `json:"eth,string"`
	LTC float64 `json:"ltc,string"`
	// OKEx
	SSC    float64 `json:"ssc,string"`
	MOF    float64 `json:"mof,string"`
	XUC    float64 `json:"xuc,string"`
	EOS    float64 `json:"eos,string"`
	FAIR   float64 `json:"fair,string"`
	KCASH  float64 `json:"kcash,string"`
	THETA  float64 `json:"theta,string"`
	VIB    float64 `json:"vib,string"`
	UGC    float64 `json:"ugc,string"`
	OST    float64 `json:"ost,string"`
	MOT    float64 `json:"mot,string"`
	BRD    float64 `json:"brd,string"`
	DNA    float64 `json:"dna,string"`
	XMR    float64 `json:"xmr,string"`
	LEV    float64 `json:"lev,string"`
	IPC    float64 `json:"ipc,string"`
	BT1    float64 `json:"bt1,string"`
	CTR    float64 `json:"ctr,string"`
	BT2    float64 `json:"bt2,string"`
	XEM    float64 `json:"xem,string"`
	NAS    float64 `json:"nas,string"`
	IOTA   float64 `json:"iota,string"`
	VIU    float64 `json:"viu,string"`
	STC    float64 `json:"stc,string"`
	WTC    float64 `json:"wtc,string"`
	TNB    float64 `json:"tnb,string"`
	DNT    float64 `json:"dnt,string"`
	LIGHT  float64 `json:"light,string"`
	IOST   float64 `json:"iost,string"`
	DGB    float64 `json:"dgb,string"`
	DGD    float64 `json:"dgd,string"`
	ZRX    float64 `json:"zrx,string"`
	SUB    float64 `json:"sub,string"`
	BCD    float64 `json:"bcd,string"`
	AAC    float64 `json:"aac,string"`
	POE    float64 `json:"poe,string"`
	OMG    float64 `json:"omg,string"`
	CMT    float64 `json:"cmt,string"`
	HOT    float64 `json:"hot,string"`
	CVC    float64 `json:"cvc,string"`
	FirST  float64 `json:"1st,string"`
	MAG    float64 `json:"mag,string"`
	BCS    float64 `json:"bcs,string"`
	BTG    float64 `json:"btg,string"`
	BCX    float64 `json:"bcx,string"`
	BTM    float64 `json:"btm,string"`
	ARK    float64 `json:"ark,string"`
	SMT    float64 `json:"smt,string"`
	RCN    float64 `json:"rcn,string"`
	KEY    float64 `json:"key,string"`
	KNC    float64 `json:"knc,string"`
	RCT    float64 `json:"rct,string"`
	SALT   float64 `json:"salt,string"`
	SNC    float64 `json:"snc,string"`
	STORJ  float64 `json:"storj,string"`
	GNT    float64 `json:"gnt,string"`
	DPY    float64 `json:"dpy,string"`
	GNX    float64 `json:"gnx,string"`
	SNM    float64 `json:"snm,string"`
	MANA   float64 `json:"mana,string"`
	PPT    float64 `json:"ppt,string"`
	LA     float64 `json:"la,string"`
	SNT    float64 `json:"snt,string"`
	SNGLS  float64 `json:"sngls,string"`
	RDN    float64 `json:"rdn,string"`
	FUN    float64 `json:"fun,string"`
	ACE    float64 `json:"ace,string"`
	AST    float64 `json:"ast,string"`
	PYN    float64 `json:"pyn,string"`
	UBTC   float64 `json:"ubtc,string"`
	UKG    float64 `json:"ukg,string"`
	REF    float64 `json:"ref,string"`
	ACT    float64 `json:"act,string"`
	YOYO   float64 `json:"yoyo,string"`
	ICN    float64 `json:"icn,string"`
	MKR    float64 `json:"mkr,string"`
	DAT    float64 `json:"dat,string"`
	ETF    float64 `json:"etf,string"`
	VEE    float64 `json:"vee,string"`
	USDT   float64 `json:"usdt,string"`
	MCO    float64 `json:"mco,string"`
	AIDOC  float64 `json:"aidoc,string"`
	TOPC   float64 `json:"topc,string"`
	ATL    float64 `json:"atl,string"`
	ZEC    float64 `json:"zec,string"`
	NEO    float64 `json:"neo,string"`
	ITC    float64 `json:"itc,string"`
	TIO    float64 `json:"tio,string"`
	LRC    float64 `json:"lrc,string"`
	ELF    float64 `json:"elf,string"`
	PRA    float64 `json:"pra,string"`
	REQ    float64 `json:"req,string"`
	ICX    float64 `json:"icx,string"`
	MTH    float64 `json:"mth,string"`
	READ   float64 `json:"read,string"`
	MTL    float64 `json:"mtl,string"`
	SPF    float64 `json:"spf,string"`
	PAY    float64 `json:"pay,string"`
	BNT    float64 `json:"bnt,string"`
	MDA    float64 `json:"mda,string"`
	F4SBTC float64 `json:"f4sbtc,string"`
	UTK    float64 `json:"utk,string"`
	EDO    float64 `json:"edo,string"`
	XRP    float64 `json:"xrp,string"`
	TRUE   float64 `json:"true,string"`
	RNT    float64 `json:"rnt,string"`
	TRX    float64 `json:"trx,string"`
	DASH   float64 `json:"dash,string"`
	MDT    float64 `json:"mdt,string"`
	NULS   float64 `json:"nuls,string"`
	AMM    float64 `json:"amm,string"`
	HSR    float64 `json:"hsr,string"`
	LINK   float64 `json:"link,string"`
	CAG    float64 `json:"cag,string"`
	SHOW   float64 `json:"show,string"`
	SBTC   float64 `json:"sbtc,string"`
	NGC    float64 `json:"ngc,string"`
	QUN    float64 `json:"qun,string"`
	QTUM   float64 `json:"qtum,string"`
	PST    float64 `json:"pst,string"`
	CAN    float64 `json:"can,string"`
	OF     float64 `json:"of,string"`
	GAS    float64 `json:"gas,string"`
	YEE    float64 `json:"yee,string"`
	LEND   float64 `json:"lend,string"`
	AVT    float64 `json:"avt,string"`
	ENG    float64 `json:"eng,string"`
	SAN    float64 `json:"san,string"`
	TCT    float64 `json:"tct,string"`
	EVX    float64 `json:"evx,string"`
	OAX    float64 `json:"oax,string"`
	WRC    float64 `json:"wrc,string"`
	QVT    float64 `json:"qvt,string"`
	INT    float64 `json:"int,string"`
	INS    float64 `json:"ins,string"`
	XLM    float64 `json:"xlm,string"`
	SWFTC  float64 `json:"swftc,string"`
}

func (e *Okex) GetBalance() (*entity.Balance, error) {
	urlString := "https://www.okex.com/api/v1/userinfo.do"
	// urlString := "https://www.okcoin.com/api/v1/userinfo.do"

	body, err := common.Post(urlString, e.getSignedValues(url.Values{}))
	if err != nil {
		return nil, fmt.Errorf("bad HTTP response: %q", err.Error())
	}

	var resp okexBalanceResponse

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

		"SSC":    resp.Info.Funds.Free.SSC,
		"MOF":    resp.Info.Funds.Free.MOF,
		"XUC":    resp.Info.Funds.Free.XUC,
		"EOS":    resp.Info.Funds.Free.EOS,
		"FAIR":   resp.Info.Funds.Free.FAIR,
		"KCASH":  resp.Info.Funds.Free.KCASH,
		"THETA":  resp.Info.Funds.Free.THETA,
		"VIB":    resp.Info.Funds.Free.VIB,
		"UGC":    resp.Info.Funds.Free.UGC,
		"OST":    resp.Info.Funds.Free.OST,
		"MOT":    resp.Info.Funds.Free.MOT,
		"BRD":    resp.Info.Funds.Free.BRD,
		"DNA":    resp.Info.Funds.Free.DNA,
		"XMR":    resp.Info.Funds.Free.XMR,
		"LEV":    resp.Info.Funds.Free.LEV,
		"IPC":    resp.Info.Funds.Free.IPC,
		"BT1":    resp.Info.Funds.Free.BT1,
		"CTR":    resp.Info.Funds.Free.CTR,
		"BT2":    resp.Info.Funds.Free.BT2,
		"XEM":    resp.Info.Funds.Free.XEM,
		"NAS":    resp.Info.Funds.Free.NAS,
		"IOTA":   resp.Info.Funds.Free.IOTA,
		"VIU":    resp.Info.Funds.Free.VIU,
		"STC":    resp.Info.Funds.Free.STC,
		"WTC":    resp.Info.Funds.Free.WTC,
		"TNB":    resp.Info.Funds.Free.TNB,
		"DNT":    resp.Info.Funds.Free.DNT,
		"LIGHT":  resp.Info.Funds.Free.LIGHT,
		"IOST":   resp.Info.Funds.Free.IOST,
		"DGB":    resp.Info.Funds.Free.DGB,
		"DGD":    resp.Info.Funds.Free.DGD,
		"ZRX":    resp.Info.Funds.Free.ZRX,
		"SUB":    resp.Info.Funds.Free.SUB,
		"BCD":    resp.Info.Funds.Free.BCD,
		"AAC":    resp.Info.Funds.Free.AAC,
		"POE":    resp.Info.Funds.Free.POE,
		"OMG":    resp.Info.Funds.Free.OMG,
		"CMT":    resp.Info.Funds.Free.CMT,
		"HOT":    resp.Info.Funds.Free.HOT,
		"CVC":    resp.Info.Funds.Free.CVC,
		"FirST":  resp.Info.Funds.Free.FirST,
		"MAG":    resp.Info.Funds.Free.MAG,
		"BCS":    resp.Info.Funds.Free.BCS,
		"BTG":    resp.Info.Funds.Free.BTG,
		"BCX":    resp.Info.Funds.Free.BCX,
		"BTM":    resp.Info.Funds.Free.BTM,
		"ARK":    resp.Info.Funds.Free.ARK,
		"SMT":    resp.Info.Funds.Free.SMT,
		"RCN":    resp.Info.Funds.Free.RCN,
		"KEY":    resp.Info.Funds.Free.KEY,
		"KNC":    resp.Info.Funds.Free.KNC,
		"RCT":    resp.Info.Funds.Free.RCT,
		"SALT":   resp.Info.Funds.Free.SALT,
		"SNC":    resp.Info.Funds.Free.SNC,
		"STORJ":  resp.Info.Funds.Free.STORJ,
		"GNT":    resp.Info.Funds.Free.GNT,
		"DPY":    resp.Info.Funds.Free.DPY,
		"GNX":    resp.Info.Funds.Free.GNX,
		"SNM":    resp.Info.Funds.Free.SNM,
		"MANA":   resp.Info.Funds.Free.MANA,
		"PPT":    resp.Info.Funds.Free.PPT,
		"LA":     resp.Info.Funds.Free.LA,
		"SNT":    resp.Info.Funds.Free.SNT,
		"SNGLS":  resp.Info.Funds.Free.SNGLS,
		"RDN":    resp.Info.Funds.Free.RDN,
		"FUN":    resp.Info.Funds.Free.FUN,
		"ACE":    resp.Info.Funds.Free.ACE,
		"AST":    resp.Info.Funds.Free.AST,
		"PYN":    resp.Info.Funds.Free.PYN,
		"UBTC":   resp.Info.Funds.Free.UBTC,
		"UKG":    resp.Info.Funds.Free.UKG,
		"REF":    resp.Info.Funds.Free.REF,
		"ACT":    resp.Info.Funds.Free.ACT,
		"YOYO":   resp.Info.Funds.Free.YOYO,
		"ICN":    resp.Info.Funds.Free.ICN,
		"MKR":    resp.Info.Funds.Free.MKR,
		"DAT":    resp.Info.Funds.Free.DAT,
		"ETF":    resp.Info.Funds.Free.ETF,
		"VEE":    resp.Info.Funds.Free.VEE,
		"USDT":   resp.Info.Funds.Free.USDT,
		"MCO":    resp.Info.Funds.Free.MCO,
		"AIDOC":  resp.Info.Funds.Free.AIDOC,
		"TOPC":   resp.Info.Funds.Free.TOPC,
		"ATL":    resp.Info.Funds.Free.ATL,
		"ZEC":    resp.Info.Funds.Free.ZEC,
		"NEO":    resp.Info.Funds.Free.NEO,
		"ITC":    resp.Info.Funds.Free.ITC,
		"TIO":    resp.Info.Funds.Free.TIO,
		"LRC":    resp.Info.Funds.Free.LRC,
		"ELF":    resp.Info.Funds.Free.ELF,
		"PRA":    resp.Info.Funds.Free.PRA,
		"REQ":    resp.Info.Funds.Free.REQ,
		"ICX":    resp.Info.Funds.Free.ICX,
		"MTH":    resp.Info.Funds.Free.MTH,
		"READ":   resp.Info.Funds.Free.READ,
		"MTL":    resp.Info.Funds.Free.MTL,
		"SPF":    resp.Info.Funds.Free.SPF,
		"PAY":    resp.Info.Funds.Free.PAY,
		"BNT":    resp.Info.Funds.Free.BNT,
		"MDA":    resp.Info.Funds.Free.MDA,
		"F4SBTC": resp.Info.Funds.Free.F4SBTC,
		"UTK":    resp.Info.Funds.Free.UTK,
		"EDO":    resp.Info.Funds.Free.EDO,
		"XRP":    resp.Info.Funds.Free.XRP,
		"TRUE":   resp.Info.Funds.Free.TRUE,
		"RNT":    resp.Info.Funds.Free.RNT,
		"TRX":    resp.Info.Funds.Free.TRX,
		"DASH":   resp.Info.Funds.Free.DASH,
		"MDT":    resp.Info.Funds.Free.MDT,
		"NULS":   resp.Info.Funds.Free.NULS,
		"AMM":    resp.Info.Funds.Free.AMM,
		"HSR":    resp.Info.Funds.Free.HSR,
		"LINK":   resp.Info.Funds.Free.LINK,
		"CAG":    resp.Info.Funds.Free.CAG,
		"SHOW":   resp.Info.Funds.Free.SHOW,
		"SBTC":   resp.Info.Funds.Free.SBTC,
		"NGC":    resp.Info.Funds.Free.NGC,
		"QUN":    resp.Info.Funds.Free.QUN,
		"QTUM":   resp.Info.Funds.Free.QTUM,
		"PST":    resp.Info.Funds.Free.PST,
		"CAN":    resp.Info.Funds.Free.CAN,
		"OF":     resp.Info.Funds.Free.OF,
		"GAS":    resp.Info.Funds.Free.GAS,
		"YEE":    resp.Info.Funds.Free.YEE,
		"LEND":   resp.Info.Funds.Free.LEND,
		"AVT":    resp.Info.Funds.Free.AVT,
		"ENG":    resp.Info.Funds.Free.ENG,
		"SAN":    resp.Info.Funds.Free.SAN,
		"TCT":    resp.Info.Funds.Free.TCT,
		"EVX":    resp.Info.Funds.Free.EVX,
		"OAX":    resp.Info.Funds.Free.OAX,
		"WRC":    resp.Info.Funds.Free.WRC,
		"QVT":    resp.Info.Funds.Free.QVT,
		"INT":    resp.Info.Funds.Free.INT,
		"INS":    resp.Info.Funds.Free.INS,
		"XLM":    resp.Info.Funds.Free.XLM,
		"SWFTC":  resp.Info.Funds.Free.SWFTC,
	}

	return &balances, nil
}
