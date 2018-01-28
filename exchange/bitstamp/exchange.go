package bitstamp

type Bitstamp struct {
	ApiKey     string
	Secret     string
	CustomerID string
	private    bool
}

func New(apiKey, secret, customerID string) *Bitstamp {
	private := apiKey != "" && secret != "" && customerID != ""

	return &Bitstamp{apiKey, secret, customerID, private}
}

type errorResponse struct {
	Error string `json:"error"`
}
