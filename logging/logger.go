package logging

import (
	"go.uber.org/zap"
)

var Logger *zap.Logger

func init() {
	// config := zap.NewDevelopmentConfig() // you may need your own config
	Logger, _ = zap.NewDevelopment()
}
