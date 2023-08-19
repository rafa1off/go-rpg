package setup

import (
	"errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var Logger *zap.Logger

func InitLogger() error {
	logFile := make(chan *os.File)

	go func() {
		if _, err := os.ReadDir("logs"); err != nil {
			os.Mkdir("logs", 0222)
		}
		file, err := os.OpenFile("logs/app.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0222)
		if err != nil {
			logFile <- nil
			return
		}
		logFile <- file
	}()

	config := zap.NewDevelopmentEncoderConfig()
	consoleEncoder := zapcore.NewConsoleEncoder(config)

	file := <-logFile
	if file == nil {
		return errors.New("err opening file")
	}

	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(file), zapcore.InfoLevel),
		zapcore.NewCore(consoleEncoder, os.Stdout, zapcore.InfoLevel),
		zapcore.NewCore(consoleEncoder, os.Stderr, zapcore.ErrorLevel),
	)

	Logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))

	return nil
}
