package services

import (
	"context"
	"go-rpg/proto"
)

type CharServer struct {
	proto.UnimplementedCharactersServer
}

func NewCharServer() *CharServer {
	return &CharServer{}
}

func (s *CharServer) Create(ctx context.Context, req *proto.CharCreateReq) (*proto.CharCreateRes, error) {
	return &proto.CharCreateRes{}, nil
}
