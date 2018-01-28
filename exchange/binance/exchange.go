package binance

type Binance struct {
	ApiKey  string
	Secret  string
	private bool
}

func New(apiKey, secret string) *Binance {
	private := apiKey != "" && secret != ""
	return &Binance{apiKey, secret, private}
}
