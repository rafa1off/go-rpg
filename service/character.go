package service

import (
	"context"
	"go-rpg/app"
	"go-rpg/proto"
	"go-rpg/setup"

	"google.golang.org/grpc"
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

func (s *charService) Register(srv *grpc.Server) {
	proto.RegisterCharactersServer(srv, s)
}

func (s *charService) Create(ctx context.Context, req *proto.CharCreateReq) (*proto.CharCreateRes, error) {
	char, err := s.core.New(req.Char)
	if err != nil {
		go setup.Logger.Error("error inserting value into db: " + err.Error())
		return nil, status.Error(codes.Internal, "internal server error")
	}

	return &proto.CharCreateRes{
		Id:   int32(char.ID),
		Char: char.Character,
	}, nil
}

func (s *charService) GetAll(req *proto.GetAllReq, stream proto.Characters_GetAllServer) error {
	limit, skip := pagination(req.Page, req.Limit)

	chars, err := s.core.All(skip, limit)
	if err != nil {
		go setup.Logger.Error("error returning all characters: " + err.Error())
		return status.Error(codes.Internal, "internal server error")
	}

	for _, i := range chars {
		stream.Send(&proto.GetAllRes{
			Id:   int32(i.ID),
			Char: i.Character,
		})
	}

	return nil
}

func (s *charService) Delete(ctx context.Context, req *proto.CharDeleteReq) (*proto.CharDeleteRes, error) {
	char, err := s.core.Get(int(req.Id))

	switch {
	case err != nil:
		return nil, status.Error(codes.Internal, "internal server error")
	case char.ID == 0:
		return nil, status.Error(codes.NotFound, "id not found")
	}

	s.core.Delete(char)

	return &proto.CharDeleteRes{
		Details: "character deleted",
		Char:    char.Character,
	}, nil
}

func (s *charService) Get(ctx context.Context, req *proto.CharGetReq) (*proto.CharGetRes, error) {
	char, err := s.core.Get(int(req.Id))

	switch {
	case err != nil:
		return nil, status.Error(codes.Internal, "internal server error")
	case char.ID == 0:
		return nil, status.Error(codes.NotFound, "id not found")
	}

	return &proto.CharGetRes{
		Id:   int32(char.ID),
		Char: char.Character,
	}, nil
}

func (s *charService) Update(ctx context.Context, req *proto.CharUpdateReq) (*proto.CharUpdateRes, error) {
	char, err := s.core.Update(int(req.Id), req.Char)

	switch {
	case err != nil:
		return nil, status.Error(codes.Internal, "internal server error")
	case char.ID == 0:
		return nil, status.Error(codes.NotFound, "id not found")
	}

	return &proto.CharUpdateRes{
		Id:   int32(char.ID),
		Char: char.Character,
	}, nil
}
