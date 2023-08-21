package characters

import (
	"context"
	"go-rpg/db"
	"go-rpg/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CharServer struct {
	proto.UnimplementedCharactersServer
	ServerOpts
}

type ServerOpts struct {
	db db.Database
}

type ServerOptsFn func(*ServerOpts)

func defaultOpts() ServerOpts {
	return ServerOpts{
		db: nil,
	}
}

func SetDB(db db.Database) func(*ServerOpts) {
	return func(so *ServerOpts) {
		so.db = db
	}
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
	char := new(req.Char)
	s.db.Create(&char)
	return &proto.CharCreateRes{
		ID:   int32(char.ID),
		Char: char.Character,
	}, nil
}

func (s *CharServer) GetAll(req *proto.GetAllReq, stream proto.Characters_GetAllServer) error {
	var chars []*Character
	err := s.db.Find(&chars)
	if err != nil {
		return status.Error(codes.Internal, "error returning all characters")
	}
	for _, i := range chars {
		stream.Send(&proto.GetAllRes{
			ID:   int32(i.ID),
			Char: i.Character,
		})
	}
	return nil
}
