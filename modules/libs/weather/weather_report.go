package weather

import (
	"bytes"
	"text/template"

	"github.com/sslime336/awbot/res"
	"github.com/sslime336/awbot/utils"
)

func BuildWeatherReportMessage(info *WeatherInfo) string {
	if info == nil || info.Infocode != "10000" {
		return ""
	}

	tmpl, err := template.ParseFiles(res.M.String["weather_report_template"])
	utils.Check(err)

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, info)
	utils.Check(err)

	return buf.String()
}

func GetFormattedWeatherInfo() string {
	return BuildWeatherReportMessage(GetWeatherInfo())
}
