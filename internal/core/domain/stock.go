package domain

type StockPrice struct {
	CurrentPrice       float64
	PriceChange        float64
	PercentChange      float64
	HighPriceOfDay     float64
	LowPriceOfDay      float64
	OpenPriceOfDay     float64
	PreviousClosePrice float64
}
