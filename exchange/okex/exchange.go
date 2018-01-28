package okex

type Okex struct {
	CustomerID string
	ApiKey     string
	Secret     string
	private    bool
}

func New(apiKey, secret, customerID string) *Okex {
	private := apiKey != "" && secret != "" && customerID != ""

	return &Okex{apiKey, secret, customerID, private}
}
