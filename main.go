package main

import (
	"go-rpg/app"
	"go-rpg/db"
	"go-rpg/server"
	"go-rpg/service"
	"go-rpg/setup"
	"os"
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

	sqldb, err := db.Empty()

	switch os.Getenv("ENVIRONMENT") {
	case "docker":
		sqldb, err = db.Postgres()
		if err != nil {
			setup.Logger.Sugar().Panic("error initializing db: " + err.Error())
		}
	default:
		sqldb, err = db.InMemory()
		if err != nil {
			setup.Logger.Sugar().Panic("error initializing db: " + err.Error())
		}
	}

	charCore := app.CharCore(sqldb)

	charService := service.Character(charCore)

	srv := server.Grpc(charService.Register)

	setup.Logger.Sugar().Fatal(srv.Run(grpcPort))
}
