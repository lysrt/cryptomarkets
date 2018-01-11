# cryptomarkets

### Go client library for interacting with multiple Cryptocurrency exchange market APIs

*Not production ready*

Convenient for building multi markets monitoring tools or trading bots.  
`cryptomarkets` library focuses on simplicity, and offers for each exchange the same interface:

```go
type Exchange interface {
    GetTicker(from, to string) (*entity.Ticker, error)
    GetBalance() (*entity.Balance, error)
    DepositAddress(currency string) (string, error)
    
    // Soon
    Withdrawal(currency, destination string, amount float64) error
}
```

### Features

HTTP Rest implementations for all exchanges, providing
* Tickers
* (Order books)
* (Transactions)
* (Available currency pairs info)

Authentication functionnalities
* Get balances
* Crypto deposits
* Crypto withdrawals
* (Order management)

### Usage

Installation

```
go get -u github.com/lysrt/cryptomarkets
```

Example using a basic exchanges store

```go
package main

import (
	"log"

	"github.com/lysrt/cryptomarkets/entity"
	"github.com/lysrt/cryptomarkets/exchange/bitstamp"
	"github.com/lysrt/cryptomarkets/exchange/bittrex"
)

type pricer interface {
	GetTicker(from, to string) (*entity.Ticker, error)
}

func main() {
	exchanges := map[string]pricer{
		"bitstamp": &bitstamp.Bitstamp{},
		"bittrex":  &bittrex.Bittrex{},
	}

	ticker1, err := exchanges["bitstamp"].GetTicker("BTC", "USD")
	if err != nil {
		log.Fatal(err)
	}

	ticker2, err := exchanges["bittrex"].GetTicker("USDT", "BTC")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Bitstamp:", ticker1.LastPrice)
	log.Println("Bittrex: ", ticker2.LastPrice)
}

```

Authenticated example

```go
// Using a config.json to store APIKeys
```

For more examples, see example/main.go
Soon: linking to a github client project
