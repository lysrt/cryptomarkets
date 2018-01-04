package binance

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func (e *Binance) getSignedRequest(values url.Values) *http.Request {
	var params []string
	for k, v := range values {
		params = append(params, fmt.Sprintf("%s=%s", k, v[0]))
	}
	totalParams := strings.Join(params, "&")

	mac := hmac.New(sha256.New, []byte(e.Secret))
	mac.Write([]byte(totalParams))
	macSum := mac.Sum(nil)
	sig := hex.EncodeToString(macSum)

	s := fmt.Sprintf("signature=%s", sig)
	// On GET requests, all parameters must be sent as QueryStrings
	qs := fmt.Sprintf("?%s&%s", totalParams, s)
	req, err := http.NewRequest("GET", urlString+qs, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("X-MBX-APIKEY", e.ApiKey)

	return req
}
