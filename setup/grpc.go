package setup

import (
	"go-rpg/db"
	"go-rpg/proto"
	"go-rpg/services"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GrpcServer struct {
	engine *grpc.Server
	port   string
}

func NewGrpcServer() *GrpcServer {
	return &GrpcServer{
		port: grpcPort,
	}
}

func (s *GrpcServer) Run() error {
	lis := make(chan net.Listener)
	go initListener(lis, s.port)

	srv := grpc.NewServer()

	postgresdb, err := db.NewPostgres()
	if err != nil {
		return err
	}

	charServerDb := services.SetDB(postgresdb)
	proto.RegisterCharactersServer(srv, services.NewCharServer(charServerDb))
	proto.RegisterRaidsServer(srv, nil)

	reflection.Register(srv)

	s.engine = srv

	go Logger.Info("server listening on: http://localhost:" + s.port)
	s.engine.Serve(<-lis)
	return nil
}
