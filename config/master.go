package config

import "github.com/spf13/viper"

// Master
type MasterEntity struct {
	Uin          int64
	Salutation   string // Salutation 称呼
	GreetingText string
}

// 主人捏
var Master *MasterEntity

func setupMasterConfig() {
	Master = &MasterEntity{
		Uin:          viper.GetInt64("master.uin"),
		Salutation:   viper.GetString("master.salutation"),
		GreetingText: viper.GetString("master.greeting_text"),
	}
}
