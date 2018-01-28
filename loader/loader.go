package loader

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/lysrt/cryptomarkets"
	"github.com/lysrt/cryptomarkets/exchange/binance"
	"github.com/lysrt/cryptomarkets/exchange/bitstamp"
	"github.com/lysrt/cryptomarkets/exchange/bittrex"
	"github.com/lysrt/cryptomarkets/exchange/gdax"
	"github.com/lysrt/cryptomarkets/exchange/okex"
	"github.com/lysrt/cryptomarkets/exchange/quoinex"
)

type exchangeConfig struct {
	Name       string `json:"name"`
	ApiKey     string `json:"apiKey"`
	Secret     string `json:"secret"`
	CustomerID string `json:"customerId"`
}

type Loader map[string]exchangeConfig

func New(fileName string) (Loader, error) {
	configs, err := readConfigFile(fileName)
	if err != nil {
		return nil, err
	}

	loader := make(Loader)

	for _, config := range configs {
		if _, ok := loader[config.Name]; ok {
			return nil, fmt.Errorf("duplicated exchange '%s' in config file %s", config.Name, fileName)
		}

		loader[config.Name] = exchangeConfig{
			Name:       config.Name,
			ApiKey:     config.ApiKey,
			Secret:     config.Secret,
			CustomerID: config.CustomerID,
		}
	}
	return loader, nil
}

func readConfigFile(fileName string) ([]exchangeConfig, error) {
	var c []exchangeConfig
	f, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("cannot find configuration file: %q", err)
	}
	if err = json.NewDecoder(f).Decode(&c); err != nil {
		return nil, fmt.Errorf("cannot parse configuration file: %q", err)
	}
	f.Close()
	return c, nil
}

func (loader Loader) GetExchange(name string) (cryptomarkets.Exchange, error) {
	c, ok := loader[name]
	if !ok {
		return nil, fmt.Errorf("no config entry for exchange %s", name)
	}

	switch name {
	case "bitstamp":
		return cryptomarkets.Exchange(bitstamp.New(c.ApiKey, c.Secret, c.CustomerID)), nil
	case "bittrex":
		return cryptomarkets.Exchange(bittrex.New(c.ApiKey, c.Secret)), nil
	case "binance":
		return cryptomarkets.Exchange(binance.New(c.ApiKey, c.Secret)), nil
	case "quoinex":
		return cryptomarkets.Exchange(quoinex.New(c.ApiKey, c.Secret, c.CustomerID)), nil
	case "gdax":
		return cryptomarkets.Exchange(gdax.New(c.ApiKey, c.Secret)), nil
	case "okex":
		return cryptomarkets.Exchange(okex.New(c.ApiKey, c.Secret, c.CustomerID)), nil
	default:
		return nil, fmt.Errorf("exchange %s is not implemented in the loader", name)
	}
}
