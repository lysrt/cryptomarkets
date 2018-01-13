package okcoin

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/url"
	"strings"
)

func (e *Okcoin) getSignedValues(values url.Values) url.Values {
	values.Add("api_key", e.ApiKey)

	// TODO: sort url values by key

	signInput := fmt.Sprintf("%s&secret_key=%s", values.Encode(), e.Secret)

	hash := md5.New()
	hash.Write([]byte(signInput))
	bytes := hash.Sum(nil)

	signature := strings.ToUpper(hex.EncodeToString(bytes))

	values.Add("sign", signature)

	return values
}
