package characters

import (
	"context"
	"go-rpg/proto"
	"go-rpg/setup"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type CharServer struct {
	proto.UnimplementedCharactersServer
	serverOpts
}

type serverOpts struct {
	db *gorm.DB
}

type ServerOptsFn func(*serverOpts)

func defaultOpts() serverOpts {
	return serverOpts{
		db: nil,
	}
}

func SetDB(db *gorm.DB) func(*serverOpts) {
	return func(opt *serverOpts) {
		opt.db = db
	}
}

func Server(opts ...ServerOptsFn) *CharServer {
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
	err := s.db.Create(&char).Error
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
	var chars []character
	err := s.db.Find(&chars).Error
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
