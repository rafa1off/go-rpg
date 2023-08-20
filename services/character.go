package services

import (
	"context"
	"go-rpg/app"
	"go-rpg/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func (s *CharServer) GetAll(req *proto.GetAllReq, stream proto.Characters_GetAllServer) error {
	res, err := s.db.Find()
	if err != nil {
		return status.Error(codes.Internal, "error returning all characters")
	}
	for _, i := range res {
		stream.Send(&proto.GetAllRes{
			ID:   int32(i.ID),
			Char: i.Character,
		})
	}
	return nil
}
