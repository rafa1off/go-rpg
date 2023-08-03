package character

import (
	"context"
	"go-rpg/setup"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	UnimplementedCharactersServer
}

func (s *Server) Create(ctx context.Context, c *Character) (*Info, error) {
	newChar := Character{
		Name:   c.Name,
		Class:  c.Class,
		Race:   c.Race,
		Health: c.Health,
	}
	info := Info{
		Details: "Character: " + newChar.Name + " created",
	}

	go setup.Logger.Info("Character: " + newChar.Name + " created")

	return &info, nil
}

func (s *Server) Get(ctx context.Context, id *Identifier) (*Character, error) {
	go setup.Logger.Error("method Get not implemented")
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}

func (s *Server) Delete(ctx context.Context, id *Identifier) (*Info, error) {
	go setup.Logger.Error("method Delete not implemented")
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func (s *Server) Update(ctx context.Context, id *Identifier) (*Character, error) {
	go setup.Logger.Error("method Update not implemented")
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
