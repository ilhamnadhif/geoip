package geoip

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Data struct {
	IP                 string  `json:"ip"`
	Network            string  `json:"network"`
	Version            string  `json:"version"`
	City               string  `json:"city"`
	Region             string  `json:"region"`
	RegionCode         string  `json:"region_code"`
	Country            string  `json:"country"`
	CountryName        string  `json:"country_name"`
	CountryCode        string  `json:"country_code"`
	CountryCapital     string  `json:"country_capital"`
	PostalCode         string  `json:"postal"`
	Latitude           float64 `json:"latitude"`
	Longitude          float64 `json:"longitude"`
	Timezone           string  `json:"timezone"`
	CountryCallingCode string  `json:"country_calling_code"`
	Currency           string  `json:"currency"`
	CurrencyName       string  `json:"currency_name"`
	Languages          string  `json:"languages"`
}

func GeoIP(ip string) (Data, error) {
	var data Data
	resp, err := http.Get(fmt.Sprintf("https://ipapi.co/%s/json", strings.TrimSpace(ip)))
	if err != nil {
		return data, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return data, err
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, err
	}

	return data, nil
}
