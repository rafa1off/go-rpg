package raid

import (
	"context"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var logger *zap.Logger

type Server struct {
	UnimplementedRaidsServer
}

func Logger(l *zap.Logger) {
	logger = l
}

func (s *Server) Create(ctx context.Context, r *Raid) (*Info, error) {
	go func() {
		logger.Error("method Create not implemented")
	}()
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}

func (s *Server) Get(ctx context.Context, id *Identifier) (*Raid, error) {
	go func() {
		logger.Error("method Get not implemented")
	}()
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}

func (s *Server) Enter(entry Raids_EnterServer) error {
	go func() {
		logger.Error("method Enter not implemented")
	}()
	return status.Errorf(codes.Unimplemented, "method Enter not implemented")
}

func (s *Server) Leave(ctx context.Context, id *Identifier) (*Info, error) {
	go func() {
		logger.Error("method Leave not implemented")
	}()
	return nil, status.Errorf(codes.Unimplemented, "method Leave not implemented")
}
