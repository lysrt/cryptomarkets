package entity

type Ticker struct {
	Timestamp    int64
	LastPrice    float64
	LastQuantity float64
	High         float64
	Low          float64
	Open         float64
	Close        float64
	Bid          float64
	Ask          float64
	VWAP         float64
	Volume       float64
	Pair         Pair
}
