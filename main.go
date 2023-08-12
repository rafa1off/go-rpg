package main

import (
	"go-rpg/character"
	"go-rpg/raid"
	"go-rpg/setup"
	"log"
	"net"
	"strconv"

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

	server := initServer()

	go setup.Logger.Info("server listening on: http://localhost:" + strconv.Itoa(defaultPort))
	setup.Logger.Sugar().Fatal(server.Serve(<-lis))
}

func initListener(c chan<- net.Listener) {
	lis, err := net.Listen("tcp", "0.0.0.0:"+strconv.Itoa(defaultPort))
	if err != nil {
		c <- nil
		return
	}
	c <- lis
}

func initServer() *grpc.Server {
	srv := grpc.NewServer()

	character.RegisterCharactersServer(srv, &character.Server{})
	raid.RegisterRaidsServer(srv, &raid.Server{})

	reflection.Register(srv)

	return srv
}
