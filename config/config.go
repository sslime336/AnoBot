package config

import (
	"github.com/sslime336/awbot/logging"

	c "github.com/spf13/viper"
)

func Init() {
	c.SetConfigName("bot")
	c.SetConfigType("yaml")
	c.AddConfigPath(".")
	err := c.ReadInConfig()
	if err != nil {
		logging.Logger.Fatal(err.Error())
	}

	setupAll()
}

func setupAll() {
	setupMasterConfig()
	setupWeatherApiConfig()
	setupResourcesConfig()
}
