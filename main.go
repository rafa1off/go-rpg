package main

import (
	"go-rpg/characters"
	"go-rpg/db"
	"go-rpg/server"
	"go-rpg/setup"
	"log"
)

const (
	grpcPort = "8000"
	restPort = "8080"
)

func main() {
	if err := setup.InitLogger(); err != nil {
		log.Println("error initializing logger: " + err.Error())
	}
	defer setup.Logger.Sync()

	pgdb, err := db.Postgres(characters.Model())
	if err != nil {
		go setup.Logger.Error("err initializing postgres: " + err.Error())
	}

	charService := characters.Server(characters.SetDB(pgdb))

	srv := server.Grpc(charService)

	setup.Logger.Sugar().Fatal(srv.Run(grpcPort))
}
