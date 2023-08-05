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

	lis := make(chan net.Listener)
	go initListener(lis)

	server := grpc.NewServer()

	character.RegisterCharactersServer(server, &character.Server{})
	raid.RegisterRaidsServer(server, &raid.Server{})

	reflection.Register(server)

	go setup.Logger.Info("server listening on: http://localhost:" + strconv.Itoa(defaultPort))
	if err := server.Serve(<-lis); err != nil {
		go setup.Logger.Error("error listening network",
			zap.String("error", err.Error()),
		)
	}
}

func initListener(c chan<- net.Listener) {
	lis, err := net.Listen("tcp", "0.0.0.0:"+strconv.Itoa(defaultPort))
	if err != nil {
		go setup.Logger.Error("error listening network",
			zap.String("details", err.Error()))
		c <- nil
		return
	}
	c <- lis
}
