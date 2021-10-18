package commodity

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	businessCommodity "monorepo/business/commodity"
)

type HTTPRequest struct{}

type CommodityResponse struct {
	UUID         string `json:"uuid"`
	Komoditas    string `json:"komoditas"`
	AreaProvinsi string `json:"area_provinsi"`
	AreaKota     string `json:"area_kota"`
	Size         string `json:"size"`
	Price        string `json:"price"`
	TglParsed    string `json:"tgl_parsed"`
	Timestamp    string `json:"timestamp"`
}
type CurrencyResponse struct {
	USD float64 `json:"USD_IDR"`
}

func NewHTTPRequestRepository() *HTTPRequest {
	return &HTTPRequest{}
}

func (repo *HTTPRequest) FetchCommodity() ([]businessCommodity.Commodity, error) {
	var (
		commodities  []CommodityResponse
		commodity    businessCommodity.Commodity
		allCommodity []businessCommodity.Commodity
	)

	res, err := http.Get("https://stein.efishery.com/v1/storages/5e1edf521073e315924ceab4/list/")
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(data, &commodities)

	for _, data := range commodities {
		commodity.UUID = data.UUID
		commodity.Komoditas = data.Komoditas
		commodity.AreaProvinsi = data.AreaProvinsi
		commodity.AreaKota = data.AreaKota
		commodity.Size = data.Size
		commodity.PriceIDR = data.Price
		commodity.TglParsed = data.TglParsed
		commodity.Timestamp = data.Timestamp

		allCommodity = append(allCommodity, commodity)
	}

	return allCommodity, nil
}

func (repo *HTTPRequest) FetchCurrencyUSDtoIDR() (float64, error) {
	res, err := http.Get("https://free.currconv.com/api/v7/convert?q=USD_IDR&compact=ultra&apiKey=acd774afe8ed441e9393")
	if err != nil {
		return 0, err
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return 0, err
	}

	var currency CurrencyResponse
	json.Unmarshal(data, &currency)

	return currency.USD, nil
}
