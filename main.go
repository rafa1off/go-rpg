package main

import (
	"go-rpg/app"
	"go-rpg/db"
	"go-rpg/server"
	"go-rpg/service"
	"go-rpg/setup"
)

const (
	grpcPort string = "8000"
	restPort string = "8080"
)

func main() {
	if err := setup.InitLogger(); err != nil {
		panic("error initializing logger: " + err.Error())
	}
	defer setup.Logger.Sync()

	db, err := db.InMemory()
	if err != nil {
		setup.Logger.Sugar().Panic("error initializing db: " + err.Error())
	}

	charCore := app.CharCore(db)

	charService := service.Character(charCore)

	srv := server.Grpc(charService.Register)

	setup.Logger.Sugar().Fatal(srv.Run(grpcPort))
}
