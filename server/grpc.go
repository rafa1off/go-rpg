package server

import (
	"go-rpg/characters"
	"go-rpg/proto"
	"go-rpg/setup"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GrpcServer struct {
	engine *grpc.Server
	port   string
}

func NewGrpc(port string) *GrpcServer {
	return &GrpcServer{
		engine: grpc.NewServer(),
		port:   port,
	}
}

func (s *GrpcServer) Run(charService *characters.CharServer) error {
	lis := make(chan net.Listener)
	go initListener(lis, s.port)

	proto.RegisterCharactersServer(s.engine, charService)

	reflection.Register(s.engine)

	go setup.Logger.Info("server listening on: http://localhost:" + s.port)
	return s.engine.Serve(<-lis)
}
