package initialize

import (
	. "code.sohuno.com/sky/robin/global"
	"go.uber.org/zap"
)

func InitLogger() {
	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)
	Logger = logger
}
