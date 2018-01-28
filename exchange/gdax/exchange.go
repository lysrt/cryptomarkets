package gdax

type Gdax struct {
	ApiKey  string
	Secret  string
	private bool
}

func New(apiKey, secret string) *Gdax {
	private := apiKey != "" && secret != ""
	return &Gdax{apiKey, secret, private}
}
