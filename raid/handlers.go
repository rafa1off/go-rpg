package raid

import (
	"context"
	"go-rpg/setup"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	UnimplementedRaidsServer
}

func (s *Server) Create(_ context.Context, r *Raid) (*Info, error) {
	go setup.Logger.Error("method Create not implemented")
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}

func (s *Server) Get(_ context.Context, id *RaidIn) (*RaidOut, error) {
	go setup.Logger.Error("method Get not implemented")
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}

func (s *Server) Enter(entry Raids_EnterServer) error {
	go setup.Logger.Error("method Enter not implemented")
	return status.Errorf(codes.Unimplemented, "method Enter not implemented")
}

func (s *Server) Leave(_ context.Context, id *RaidIn) (*Info, error) {
	go setup.Logger.Error("method Leave not implemented")
	return nil, status.Errorf(codes.Unimplemented, "method Leave not implemented")
}
