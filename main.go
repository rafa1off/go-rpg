package main

import (
	"go-rpg/character"
	"go-rpg/raid"
	"net"
	"os"
	"strconv"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const defaultPort int = 8000

func main() {
	logger := initLogger()

	character.Logger(logger)
	raid.Logger(logger)

	lis := initListener(logger)

	server := grpc.NewServer()

	character.RegisterCharactersServer(server, &character.Server{})
	raid.RegisterRaidsServer(server, &raid.Server{})

	reflection.Register(server)

	logger.Info("server listening on: http://localhost:" + strconv.Itoa(defaultPort))
	if err := server.Serve(lis); err != nil {
		go logger.Error("error listening network",
			zap.String("error", err.Error()),
		)
	}
}

func initLogger() (logger *zap.Logger) {
	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	consoleDebugging := zapcore.Lock(os.Stdout)
	consoleErrors := zapcore.Lock(os.Stderr)
	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, consoleDebugging, zapcore.InfoLevel),
		zapcore.NewCore(consoleEncoder, consoleErrors, zapcore.ErrorLevel),
	)
	logger = zap.New(core)
	return
}

func initListener(l *zap.Logger) net.Listener {
	lis, err := net.Listen("tcp", "0.0.0.0:"+strconv.Itoa(defaultPort))
	if err != nil {
		go l.Error("error listening network",
			zap.String("error", err.Error()),
		)
	}
	return lis
}
