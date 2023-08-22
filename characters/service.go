package characters

import (
	"context"
	"go-rpg/db"
	"go-rpg/proto"
	"go-rpg/setup"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CharServer struct {
	proto.UnimplementedCharactersServer
	serverOpts
}

type serverOpts struct {
	db db.Database
}

type ServerOptsFn func(*serverOpts)

func defaultOpts() serverOpts {
	return serverOpts{
		db: nil,
	}
}

func SetDB(db db.Database) func(*serverOpts) {
	return func(so *serverOpts) {
		so.db = db
	}
}

func NewCharServer(opts ...ServerOptsFn) *CharServer {
	opt := defaultOpts()
	for _, fn := range opts {
		fn(&opt)
	}
	return &CharServer{
		serverOpts: opt,
	}
}

func (s *CharServer) Create(ctx context.Context, req *proto.CharCreateReq) (*proto.CharCreateRes, error) {
	char := new(req.Char)
	err := s.db.Create(&char)
	if err != nil {
		go setup.Logger.Error("error inserting value into db: " + err.Error())
		return nil, status.Error(codes.Internal, "internal server error")
	}
	return &proto.CharCreateRes{
		ID:   int32(char.ID),
		Char: char.Character,
	}, nil
}

func (s *CharServer) GetAll(req *proto.GetAllReq, stream proto.Characters_GetAllServer) error {
	var chars []*character
	err := s.db.Find(&chars)
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
