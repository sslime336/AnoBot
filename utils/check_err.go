package utils

import (
	"runtime"
	"strconv"

	"github.com/sslime336/awbot/logging"
	"go.uber.org/zap"
)

func Check(err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		position := file + ":" + strconv.Itoa(line)
		logging.Logger.Error(err.Error(), zap.String("position", position))
	}
}
