package server

import (
	"go-rpg/proto"
	"go-rpg/setup"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type grpcServer struct {
	engine    *grpc.Server
	character proto.CharactersServer
}

func Grpc(charService proto.CharactersServer) *grpcServer {
	return &grpcServer{
		engine:    grpc.NewServer(),
		character: charService,
	}
}

func loadServices(s *grpcServer) {
	proto.RegisterCharactersServer(s.engine, s.character)
}

func (s *grpcServer) Run(port string) error {
	lis := make(chan net.Listener)
	go initListener(lis, port)

	loadServices(s)

	reflection.Register(s.engine)

	go setup.Logger.Info("server listening on: http://localhost:" + port)
	return s.engine.Serve(<-lis)
}
