package entity

import "strings"

type Currency string

func NewCurrency(currency string) Currency {
	return Currency(strings.ToUpper(currency))
}

func (c Currency) Lower() string {
	return strings.ToLower(string(c))
}

func (c Currency) Upper() string {
	return string(c)
}

type Pair struct {
	First, Second Currency
}

func (p Pair) Lower(sep string) string {
	return p.First.Lower() + sep + p.Second.Lower()
}

func (p Pair) Upper(sep string) string {
	return p.First.Upper() + sep + p.Second.Upper()
}
