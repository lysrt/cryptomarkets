# cryptomarkets

### Go client library for interacting with multiple Cryptocurrency exchange market APIs

*Not production ready*

Convenient for building multi markets monitoring tools or trading bots.  
`cryptomarkets` library focuses on simplicity, and offers for each exchange a common interface:

```go
type Exchange interface {
	GetTicker(from, to string) (*Ticker, error)
	OrderBook(from, to string) (*OrderBook, error)
	GetBalance() (*Balance, error)

	DepositAddress(currency string) (string, error)
	Withdrawal(currency, destination string, amount float64) (int, error)
}
```

### Features

HTTP Rest implementations for all exchanges, providing
* Tickers
* Order books
* (Transactions)
* (Available currency pairs info)

Authentication functionnalities
* Get balances
* Crypto deposits
* Crypto withdrawals
* Order management (buy/sell, limit and market orders)

### Usage

#### Installation

```
go get -u github.com/lysrt/cryptomarkets
```

#### Example using a basic exchanges store

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

#### Authenticated example

```go
var (
	apiKey     = "abc...xyz"
	secret     = "secret_key"
	customerId = "123456"
)
exchange := bitstamp.New(apiKey, secret, customerId)

balance, _ := exchange.GetBalance()

fmt.Println(balance["BTC"])
```

#### Using Loader and a JSON config file to store API keys

`config.json`

```json
[
	{
        "name": "bitstamp",
        "apiKey" : "",
        "secret" : "",
        "customerId" : ""
    }
]
```

`main.go`

```go
loader, err := loader.New("config.json")
if err != nil {
	log.Fatal(err)
}

exchange, err := loader.GetExchange("bitstamp")
if err != nil {
	log.Fatal(err)
}

balance, _ := exchange.GetBalance()

log.Println(balance)
```

For a full example, see [example/main.go]

### Code organization

`exchange/`

Each package inside `exchange` has one main public struct, like `exchange/bitstamp.Bitstamp`. The methods of this struct are exploded across several files. Although this is often not a good practice in Go, it keeps files short and ensure each package has the same structure.

`exchange/*/exchange.go`

*exchange.go* contains the struct definition and any types used by several files of the package.

`exchange/*/ticker.go`

*ticker.go* holds methods to query the public API of an exchange, to get prices, order book, and general information from an exchange.

`exchange/*/auth.go`

*auth.go* defines authentication specific logic for an exchange. It is used by all private API functions.

`exchange/*/balance.go`

*balance.go* is used to get the user account balances.

`exchange/*/order.go`

*order.go* contains all the methods used to trade on the exchange: buy and sell, for limit and market orders.

`exchange/*/transfer.go`

*transfer.go* is used to transfer funds to and from an exchange.