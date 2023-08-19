package app

import (
	"context"
	"go-rpg/proto"
)

type Server struct {
	proto.UnimplementedCharactersServer
}

func NewCharServer() *Server {
	return &Server{}
}

func (s *Server) Create(ctx context.Context, req *proto.CharCreateReq) (*proto.CharCreateRes, error) {
	return &proto.CharCreateRes{}, nil
}
