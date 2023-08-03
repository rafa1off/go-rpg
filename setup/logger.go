package setup

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var Logger *zap.Logger

func InitLogger() error {
	config := zap.NewDevelopmentEncoderConfig()
	consoleEncoder := zapcore.NewConsoleEncoder(config)
	logFile, err := os.OpenFile("logs/app.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0222)
	if err != nil {
		return err
	}
	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(logFile), zapcore.InfoLevel),
		zapcore.NewCore(consoleEncoder, os.Stdout, zapcore.InfoLevel),
		zapcore.NewCore(consoleEncoder, os.Stderr, zapcore.ErrorLevel),
	)

	Logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))

	return nil
}
