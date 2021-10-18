package commodity

import (
	"fmt"
	"strconv"

	golinq "github.com/ahmetb/go-linq/v3"
)

var currencyUSD float64

type Service interface {
	FindAllCommodity() ([]Commodity, error)
	FindAllCommodityByProvinsi() ([]CommodityReport, error)
}

type service struct {
	commodityRepository Repository
}

func NewService(commodityRepository Repository) Service {

	return &service{
		commodityRepository,
	}
}

func (s *service) FindAllCommodity() ([]Commodity, error) {
	var (
		commodity    Commodity
		allCommodity []Commodity
	)

	commodities, err := s.commodityRepository.FetchCommodity()
	if err != nil {
		return nil, err
	}

	if currencyUSD == 0 {
		currencyUSD, err = s.commodityRepository.FetchCurrencyUSDtoIDR()
		if err != nil {
			return nil, err
		}
	}

	for _, data := range commodities {
		if data.UUID != "" && data.Komoditas != "" {

			priceIDRFloat64, _ := strconv.ParseFloat(data.PriceIDR, 64)
			commodity.UUID = data.UUID
			commodity.Komoditas = data.Komoditas
			commodity.AreaProvinsi = data.AreaProvinsi
			commodity.AreaKota = data.AreaKota
			commodity.Size = data.Size
			commodity.PriceIDR = data.PriceIDR
			commodity.PriceUSD = fmt.Sprintf("%2f", (priceIDRFloat64 / currencyUSD))
			commodity.TglParsed = data.TglParsed
			commodity.Timestamp = data.Timestamp

			allCommodity = append(allCommodity, commodity)
		}
	}

	return allCommodity, nil
}

func (s *service) FindAllCommodityByProvinsi() ([]CommodityReport, error) {
	var (
		commodity    Commodity
		allCommodity []Commodity
	)

	commodities, err := s.commodityRepository.FetchCommodity()
	if err != nil {
		return nil, err
	}

	if currencyUSD == 0 {
		currencyUSD, err = s.commodityRepository.FetchCurrencyUSDtoIDR()
		if err != nil {
			return nil, err
		}
	}

	fmt.Println(currencyUSD)

	for _, data := range commodities {
		if data.UUID != "" && data.Komoditas != "" {

			priceIDRFloat64, _ := strconv.ParseFloat(data.PriceIDR, 64)
			commodity.UUID = data.UUID
			commodity.Komoditas = data.Komoditas
			commodity.AreaProvinsi = data.AreaProvinsi
			commodity.AreaKota = data.AreaKota
			commodity.Size = data.Size
			commodity.PriceIDR = data.PriceIDR
			commodity.PriceUSD = fmt.Sprintf("%2f", (priceIDRFloat64 / currencyUSD))
			commodity.TglParsed = data.TglParsed
			commodity.Timestamp = data.Timestamp

			allCommodity = append(allCommodity, commodity)
		}
	}

	// golinq
	var query []golinq.Group
	golinq.From(allCommodity).GroupByT(
		func(c Commodity) string { return c.AreaProvinsi },
		func(c Commodity) string { return c.Size },
	).OrderByT(
		func(g golinq.Group) string { return g.Key.(string) },
	).ToSlice(&query)

	var sizeArrayFloat, sizeArraySorted []float64
	var reports []CommodityReport
	var report CommodityReport

	// parse query to report data
	for _, comGroup := range query {
		// sorted size
		for _, val := range comGroup.Group {
			sizeString := val.(string)
			sizeFloat, _ := strconv.ParseFloat(sizeString, 64)
			sizeArrayFloat = append(sizeArrayFloat, sizeFloat)
		}

		golinq.From(sizeArrayFloat).Sort(func(i interface{}, j interface{}) bool {
			return i.(float64) < j.(float64)
		}).ToSlice(&sizeArraySorted)

		max := golinq.From(sizeArraySorted).Max()
		maxInt, _ := max.(float64)

		min := golinq.From(sizeArraySorted).Min()
		minInt, _ := min.(float64)

		average := golinq.From(sizeArraySorted).Average()

		median := len(sizeArraySorted) / 2

		report.AreaProvinsi = fmt.Sprintf("%v", comGroup.Key)
		report.Max = maxInt
		report.Min = minInt
		report.Average = average
		report.Median = median

		reports = append(reports, report)
	}

	return reports, nil
}
