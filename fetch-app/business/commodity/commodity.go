package commodity

type Commodity struct {
	UUID         string
	Komoditas    string
	AreaProvinsi string
	AreaKota     string
	Size         string
	PriceIDR     string
	PriceUSD     string
	TglParsed    string
	Timestamp    string
}

type CommodityReport struct {
	AreaProvinsi string
	Min          float64
	Max          float64
	Median       int
	Average      float64
}
