package bittrex

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// TODO Is it correct?
func (e *Bittrex) getAuthValues() url.Values {
	nonce := strconv.FormatInt(time.Now().Unix(), 10)

	message := nonce + e.ApiKey

	mac := hmac.New(sha256.New, []byte(e.Secret))
	mac.Write([]byte(message))
	macSum := mac.Sum(nil)
	sig := strings.ToUpper(hex.EncodeToString(macSum))

	form := url.Values{
		"key":       {e.ApiKey},
		"signature": {sig},
		"nonce":     {nonce},
	}
	return form
}
