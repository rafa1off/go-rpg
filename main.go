package main

import (
	"fmt"
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
		panic(fmt.Errorf("error initializing logger: %w", err))
	}
	defer setup.Logger.Sync()

	sqldb, err := db.Postgres()
	if err != nil {
		setup.Logger.Sugar().Panicf("error initializing db: ", err)
	}

	charCore := app.CharCore(sqldb)

	charService := service.Character(charCore)

	srv := server.Grpc(charService.Register)

	setup.Logger.Sugar().Fatal(srv.Run(grpcPort))
}
