package config

import "github.com/spf13/viper"

var (
	ApiKey   string // 高德 key
	CityCode string // 城市 adcode
)

func setupWeatherApiConfig() {
	ApiKey = viper.GetString("weather_api.key")
	CityCode = viper.GetString("weather_api.city_code")
}
