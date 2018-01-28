package quoinex

type Quoinex struct {
	ApiKey     string
	Secret     string
	CustomerID string
	private    bool
}

func New(apiKey, secret, customerID string) *Quoinex {
	private := apiKey != "" && secret != "" && customerID != ""
	return &Quoinex{apiKey, secret, customerID, private}
}
