package commodity

type Repository interface {
	// Fetch all commodity
	FetchCommodity() ([]Commodity, error)
	// Fetch currency USD to IDR
	FetchCurrencyUSDtoIDR() (float64, error)
}
