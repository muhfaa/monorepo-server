package response

import (
	businessCommodity "monorepo/business/commodity"
)

type CommodityResponse struct {
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

type CommodityReportResponse struct {
	AreaProvinsi string
	Min          float64
	Max          float64
	Median       int
	Average      float64
}

func NewCommodityDataResponse(commodities []businessCommodity.Commodity) []CommodityResponse {

	var (
		response      []CommodityResponse
		commodityData CommodityResponse
	)
	for _, value := range commodities {
		commodityData.UUID = value.UUID
		commodityData.Komoditas = value.Komoditas
		commodityData.AreaProvinsi = value.AreaProvinsi
		commodityData.AreaKota = value.AreaKota
		commodityData.Size = value.Size
		commodityData.PriceIDR = value.PriceIDR
		commodityData.PriceUSD = value.PriceUSD
		commodityData.TglParsed = value.TglParsed
		commodityData.Timestamp = value.Timestamp

		response = append(response, commodityData)
	}

	return response
}

func NewCommodityDataReportResponse(commoditis []businessCommodity.CommodityReport) []CommodityReportResponse {

	var (
		response      []CommodityReportResponse
		commodityData CommodityReportResponse
	)

	for _, value := range commoditis {
		commodityData.AreaProvinsi = value.AreaProvinsi
		commodityData.Min = value.Min
		commodityData.Max = value.Max
		commodityData.Median = value.Median
		commodityData.Average = value.Average

		response = append(response, commodityData)
	}

	return response
}
