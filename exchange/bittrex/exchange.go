package bittrex

type Bittrex struct {
	ApiKey  string
	Secret  string
	private bool
}

func New(apiKey, secret string) *Bittrex {
	private := apiKey != "" && secret != ""
	return &Bittrex{apiKey, secret, private}
}
