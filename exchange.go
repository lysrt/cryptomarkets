package cryptomarkets

type Exchange interface {
	LastPrice(from, to string) float64
	GetBalance(currency string) float64
	// BitcoinDepositAddress() string
	// BitcoinWithdrawal(destination, value string) string
}

type ExchangeConfig struct {
	Name       string
	ApiKey     string
	Secret     string
	CustomerID string
}
