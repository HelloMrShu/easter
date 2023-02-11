package global

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func InitLogger() {
	writer := getLogWriter()
	encoder := getEncoder()

	core := zapcore.NewCore(encoder, writer, zapcore.InfoLevel)
	logger := zap.New(core, zap.AddCaller())
	Logger = logger
}

func getEncoder() zapcore.Encoder {
	return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
}

func getLogWriter() zapcore.WriteSyncer {
	file, _ := os.Create("logs/app.log")
	return zapcore.AddSync(file)
}
