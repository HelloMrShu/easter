package global

import (
	"fmt"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitLogger() {
	writer := getLogWriter()
	encoder := getEncoder()

	core := zapcore.NewCore(encoder, writer, zapcore.InfoLevel)
	Logger = zap.New(core, zap.AddCaller())
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	filename := fmt.Sprintf("%s/%s", ServerConfig.Log.Path, ServerConfig.Log.File)
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename, //日志文件名称
		MaxSize:    50,       //文件大小限制，默认MB
		MaxBackups: 5,        //备份文件数量
		MaxAge:     30,       //最大天数
		Compress:   false,    //是否压缩
	}
	return zapcore.AddSync(lumberJackLogger)
}
