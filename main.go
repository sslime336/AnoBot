package main

import (
	"github.com/sslime336/awbot/bot"
	"github.com/sslime336/awbot/config"
	"github.com/sslime336/awbot/logging"
	"github.com/sslime336/awbot/modules"
	"github.com/sslime336/awbot/res"
)

func init() {
	logging.Logger.Info("开始读取配置...")
	config.Init()
	logging.Logger.Info("配置读取完成")
}

func main() {
	logging.Logger.Info("开始配置 Bot...")
	bot.Setup()
	logging.Logger.Info("登录成功!")

	res.Load()
	logging.Logger.Info("资源文件已加载")

	modules.Load()
	logging.Logger.Info("模块装载完成")

	bot.StartScheduleMissions()
	logging.Logger.Info("定时任务已开启")

	logging.Logger.Info("Bot 已上线")
	bot.KeepRunning()
}
