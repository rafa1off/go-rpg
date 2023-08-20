package main

import (
	"go-rpg/setup"
	"log"
)

func main() {
	if err := setup.InitLogger(); err != nil {
		log.Println("error initializing logger: " + err.Error())
	}
	defer setup.Logger.Sync()

	server := setup.NewGrpcServer()

	setup.Logger.Sugar().Fatal(server.Run())
}
