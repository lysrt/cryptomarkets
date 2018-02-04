package bitstamp

import (
	"encoding/json"
	"errors"
)

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

func (e *Bitstamp) checkResponse(body []byte) error {
	// Bitstamp can return HTTP Status 200 with a JSON error
	var response errorResponse
	err := json.Unmarshal(body, &response)
	if err == nil {
		return errors.New(response.Error)
	}

	return nil
}
