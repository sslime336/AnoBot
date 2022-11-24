package weather

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/sslime336/awbot/config"
	"github.com/sslime336/awbot/utils"
)

type WeatherInfo struct {
	Status    string     `json:"status"`
	Count     string     `json:"count"`
	Info      string     `json:"info"`
	Infocode  string     `json:"infocode"`
	Forecasts []Forecast `json:"forecasts"`
}

type Forecast struct {
	City       string `json:"city"`
	Adcode     string `json:"adcode"`
	Province   string `json:"province"`
	Reporttime string `json:"reporttime"`
	Casts      []Cast `json:"casts"`
}

type Cast struct {
	Date         string `json:"date"`
	Week         string `json:"week"`
	Dayweather   string `json:"dayweather"`
	Nightweather string `json:"nightweather"`
	Daytemp      string `json:"daytemp"`
	Nighttemp    string `json:"nighttemp"`
	Daywind      string `json:"daywind"`
	Nightwind    string `json:"nightwind"`
	Daypower     string `json:"daypower"`
	Nightpower   string `json:"nightpower"`
}

const _requestUrl = `https://restapi.amap.com/v3/weather/weatherInfo`

func GetWeatherInfo() *WeatherInfo {
	payload := getQueryPayload()
	request, err := http.NewRequest(http.MethodGet, _requestUrl+payload, nil)
	utils.Check(err)

	{
		request.Header.Add("User-Agent", "apifox/1.0.0 (https://www.apifox.cn)")
		request.Header.Add("Accept", "*/*")
		request.Header.Add("Host", "restapi.amap.com")
		request.Header.Add("Accept-Encoding", "gzip, deflate, br")
		request.Header.Add("Connection", "keep-alive")
	}

	// send request
	c := new(http.Client)
	resp, err := c.Do(request)
	utils.Check(err)
	defer resp.Body.Close()

	// unzip
	ungzipedbody, err := gzip.NewReader(resp.Body)
	utils.Check(err)
	defer resp.Body.Close()
	data, err := io.ReadAll(ungzipedbody)
	utils.Check(err)

	weatherInfo := new(WeatherInfo)
	err = json.Unmarshal(data, weatherInfo)
	utils.Check(err)

	return weatherInfo
}

func getQueryPayload() string {
	return fmt.Sprintf(
		"?city=%s&extensions=all&key=%s&output=JSON",
		config.CityCode, config.ApiKey)
}
