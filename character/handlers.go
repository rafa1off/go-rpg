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
	ServerOpts
}

type ServerOpts struct {
	DB database
}

func NewServer(opts ...ServerOpts) *Server {
	return &Server{
		ServerOpts: opts[0],
	}
}

func (s *Server) Create(_ context.Context, req *CharCreateReq) (*CharCreateRes, error) {
	item := s.DB.Save(req.Char)

	return &CharCreateRes{
		Id:   item.id,
		Char: item.char,
	}, nil
}

func (s *Server) Get(_ context.Context, req *CharGetReq) (*CharGetRes, error) {
	char, err := s.DB.Get(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "id not found")
	}

	return &CharGetRes{
		Id:   req.Id,
		Char: char,
	}, nil
}

func (s *Server) Delete(_ context.Context, req *CharDeleteReq) (*CharDeleteRes, error) {
	if err := s.DB.Delete(req.Id); err != nil {
		return nil, status.Errorf(codes.NotFound, "id not found")
	}

	return &CharDeleteRes{
		Details: "Char deleted",
	}, nil
}

func (s *Server) Update(_ context.Context, data *CharUpdateReq) (*CharUpdateRes, error) {
	item := char{
		id:   data.Id,
		char: data.Char,
	}

	char, err := s.DB.Update(&item)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "id not found")
	}

	return &CharUpdateRes{
		Id:   item.id,
		Char: char,
	}, nil
}

func (s *Server) GetAll(_ *GetAllReq, stream Characters_GetAllServer) error {
	for _, i := range s.DB.GetAll() {
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
