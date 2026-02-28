package finnhub_api

type Quote struct {
	CurrentPrice        float64 `json:"c"`  // Current price
	PriceChange         float64 `json:"d"`  // Change
	PercentChange       float64 `json:"dp"` // Percent change
	HighPriceOfDay      float64 `json:"h"`  // High price of the day
	LowPriceOfDay       float64 `json:"l"`  // Low price of the day
	OpenPriceOfDay      float64 `json:"o"`  // Open price of the day
	PreviousClosePrice  float64 `json:"pc"` // Previous close price
	TimestampUnixSecond int64   `json:"t"`  // Timestamp
}

type CompanyProfile struct {
	Country              string  `json:"country"`
	Currency             string  `json:"currency"`
	Exchange             string  `json:"exchange"`
	IPO                  string  `json:"ipo"`
	MarketCapitalization float64 `json:"marketCapitalization"`
	Name                 string  `json:"name"`
	Phone                string  `json:"phone"`
	ShareOutstanding     float64 `json:"shareOutstanding"`
	Ticker               string  `json:"ticker"`
	WebURL               string  `json:"weburl"`
	Logo                 string  `json:"logo"`
	FinnhubIndustry      string  `json:"finnhubIndustry"`
	EstimateCurrency     string  `json:"estimateCurrency"`
}
