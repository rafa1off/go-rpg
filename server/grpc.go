package server

import (
	"go-rpg/proto"
	"go-rpg/service"
	"go-rpg/setup"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type grpcServer struct {
	engine   *grpc.Server
	services []interface{}
}

func Grpc(services ...interface{}) *grpcServer {
	return &grpcServer{
		engine:   grpc.NewServer(),
		services: services,
	}
}

func loadServices(s *grpcServer) {
	for _, i := range s.services {
		switch i {
		case i.(*service.CharService):
			proto.RegisterCharactersServer(s.engine, i.(*service.CharService))
		}
	}
}

func (s *grpcServer) Run(port string) error {
	lis := make(chan net.Listener)
	go initListener(lis, port)

	loadServices(s)

	reflection.Register(s.engine)

	go setup.Logger.Info("server listening on: http://localhost:" + port)
	return s.engine.Serve(<-lis)
}
