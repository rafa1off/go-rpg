package character

import (
	"context"
	"go-rpg/setup"

	"go.uber.org/zap"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	UnimplementedCharactersServer
	db database
}

func (s *Server) Create(_ context.Context, req *CharCreateReq) (*CharCreateRes, error) {
	item := s.db.Save(req.Char)

	res := CharCreateRes{
		Id:   item.id,
		Char: item.char,
	}

	return &res, nil
}

func (s *Server) Get(_ context.Context, req *CharGetReq) (*CharGetRes, error) {
	char, err := s.db.Get(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "id not found")
	}

	res := CharGetRes{
		Id:   req.Id,
		Char: char,
	}

	return &res, nil
}

func (s *Server) Delete(_ context.Context, req *CharDeleteReq) (*CharDeleteRes, error) {
	if err := s.db.Delete(req.Id); err != nil {
		return nil, status.Errorf(codes.NotFound, "id not found")
	}

	info := CharDeleteRes{
		Details: "Char deleted",
	}

	return &info, nil
}

func (s *Server) Update(_ context.Context, data *CharUpdateReq) (*CharUpdateRes, error) {
	item := char{
		id:   data.Id,
		char: data.Char,
	}

	char, err := s.db.Update(&item)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "id not found")
	}

	updatedChar := CharUpdateRes{
		Id:   item.id,
		Char: char,
	}

	return &updatedChar, nil
}

func (s *Server) GetAll(_ *GetAllReq, stream Characters_GetAllServer) error {
	for _, i := range s.db {
		res := GetAllRes{
			Id:   i.id,
			Char: i.char,
		}
		if err := stream.Send(&res); err != nil {
			go setup.Logger.Error("error sending stream",
				zap.String("err: ", err.Error()),
			)
			return err
		}
	}

	return nil
}
