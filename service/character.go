package service

import (
	"context"
	"go-rpg/app"
	"go-rpg/proto"
	"go-rpg/setup"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type charService struct {
	proto.UnimplementedCharactersServer
	core app.Core
}

func Character(core app.Core) *charService {
	return &charService{
		core: core,
	}
}

func (s *charService) Create(ctx context.Context, req *proto.CharCreateReq) (*proto.CharCreateRes, error) {
	char, err := s.core.New(req.Char)
	if err != nil {
		go setup.Logger.Error("error inserting value into db: " + err.Error())
		return nil, status.Error(codes.Internal, "internal server error")
	}
	return &proto.CharCreateRes{
		ID:   int32(char.ID),
		Char: char.Character,
	}, nil
}

func (s *charService) GetAll(req *proto.GetAllReq, stream proto.Characters_GetAllServer) error {
	chars, err := s.core.All()
	if err != nil {
		go setup.Logger.Error("error returning all characters: " + err.Error())
		return status.Error(codes.Internal, "internal server error")
	}
	for _, i := range chars {
		stream.Send(&proto.GetAllRes{
			ID:   int32(i.ID),
			Char: i.Character,
		})
	}
	return nil
}
