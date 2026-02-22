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

type CompanyProfile struct {
	Country              string
	Currency             string
	Exchange             string
	IPO                  string
	MarketCapitalization float64
	Name                 string
	Phone                string
	ShareOutstanding     float64
	Ticker               string
	WebURL               string
	Logo                 string
	FinnhubIndustry      string
}

type StockSummary struct {
	// Stock price fields
	CurrentPrice       float64
	PriceChange        float64
	PercentChange      float64
	HighPriceOfDay     float64
	LowPriceOfDay      float64
	OpenPriceOfDay     float64
	PreviousClosePrice float64

	// Company profile fields
	Name     string
	Exchange string
	Ticker   string
	Currency string
}
