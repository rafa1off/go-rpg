package server

import (
	"go-rpg/setup"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type grpcServer struct {
	engine   *grpc.Server
	services []register
}

type register func(srv *grpc.Server)

func Grpc(services ...register) *grpcServer {
	return &grpcServer{
		engine:   grpc.NewServer(),
		services: services,
	}
}

func (s *grpcServer) loadServices() {
	for _, service := range s.services {
		service(s.engine)
	}
}

func (s *grpcServer) Run(port string) error {
	lis := make(chan net.Listener)
	go initListener(lis, port)

	s.loadServices()

	reflection.Register(s.engine)

	go setup.Logger.Info("server listening on: http://localhost:" + port)
	return s.engine.Serve(<-lis)
}
