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

	pgdb, err := db.NewPostgres(&characters.Character{})
	if err != nil {
		setup.Logger.Info("err initing postgres: " + err.Error())
	}

	charServiceDB := characters.SetDB(pgdb)
	charSevice := characters.NewCharServer(charServiceDB)

	srv := server.NewGrpc(grpcPort)

	setup.Logger.Sugar().Fatal(srv.Run(charSevice))
}
