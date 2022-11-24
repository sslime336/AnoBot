package bot

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/sslime336/awbot/logging"
	"go.uber.org/zap"
)

func KeepRunning() {
	sig := make(chan os.Signal, 1)

	signal.Notify(sig, syscall.SIGINT|syscall.SIGKILL)

	logging.Logger.Info("Bot 已退出", zap.String("reason", (<-sig).String()))

	closeAll()
}

func closeAll() {
	// 关闭定时任务
	cronManager.Stop()

	// 断开 Bot 连接
	Bot.Disconnect()
	Bot.Release()
}
