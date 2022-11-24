package groups

import (
	. "github.com/sslime336/awbot/bot/export"
	"github.com/sslime336/awbot/modules/libs/weather"
	"github.com/sslime336/awbot/res"
	"github.com/sslime336/awbot/utils"
)

func DailyWeatherReport() {
	Bot.SendGroupMessage(
		res.M.Group["placeholder_group"],
		utils.NewText(weather.GetFormattedWeatherInfo()))
}
