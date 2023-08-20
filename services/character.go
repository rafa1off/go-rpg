package services

import (
	"context"
	"go-rpg/app"
	"go-rpg/proto"
)

type CharServer struct {
	proto.UnimplementedCharactersServer
	ServerOpts
}

func NewCharServer(opts ...ServerOptsFn) *CharServer {
	opt := defaultOpts()
	for _, fn := range opts {
		fn(&opt)
	}
	return &CharServer{
		ServerOpts: opt,
	}
}

func (s *CharServer) Create(ctx context.Context, req *proto.CharCreateReq) (*proto.CharCreateRes, error) {
	char := app.NewChar(req.Char)
	s.db.Save(&char)
	return &proto.CharCreateRes{
		ID:   int32(char.ID),
		Char: char.Character,
	}, nil
}
