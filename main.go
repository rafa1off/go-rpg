package main

import (
	"go-rpg/character"
	"go-rpg/raid"
	"go-rpg/setup"
	"log"
	"net"
	"strconv"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const defaultPort int = 8000

func main() {
	if err := setup.InitLogger(); err != nil {
		log.Println("error initializing logger: " + err.Error())
	}

	lis := initListener()

	server := grpc.NewServer()

	character.RegisterCharactersServer(server, &character.Server{})
	raid.RegisterRaidsServer(server, &raid.Server{})

	reflection.Register(server)

	setup.Logger.Info("server listening on: http://localhost:" + strconv.Itoa(defaultPort))
	if err := server.Serve(lis); err != nil {
		go setup.Logger.Error("error listening network",
			zap.String("error", err.Error()),
		)
	}
}

func initListener() net.Listener {
	lis, err := net.Listen("tcp", "0.0.0.0:"+strconv.Itoa(defaultPort))
	if err != nil {
		go setup.Logger.Error("error listening network",
			zap.String("details", err.Error()))
	}
	return lis
}
