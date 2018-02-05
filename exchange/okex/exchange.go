package okex

type Okex struct {
	ApiKey     string
	Secret     string
	CustomerID string
	private    bool
}

func New(apiKey, secret, customerID string) *Okex {
	private := apiKey != "" && secret != "" && customerID != ""

	return &Okex{apiKey, secret, customerID, private}
}
