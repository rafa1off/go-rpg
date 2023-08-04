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
}

func (s *Server) Create(ctx context.Context, c *Character) (*Info, error) {
	char := Character{
		Id:     c.Id,
		Name:   c.Name,
		Class:  c.Class,
		Race:   c.Race,
		Health: c.Health,
	}
	char.Save()

	info := Info{
		Details: "Character: " + char.Name + " created",
	}

	go setup.Logger.Info("Character: " + char.Name + " created")
	return &info, nil
}

func (s *Server) Get(ctx context.Context, id *Identifier) (*Character, error) {
	char, err := Get(id.Id)
	if err != nil {
		go setup.Logger.Error("not found",
			zap.String("details", status.Errorf(codes.NotFound, "id not found").Error()),
		)
		return nil, status.Errorf(codes.NotFound, "id not found")
	}
	go setup.Logger.Info("Character: " + char.Name + " returned")
	return char, nil
}

func (s *Server) Delete(ctx context.Context, id *Identifier) (*Info, error) {
	if err := Delete(id.Id); err != nil {
		go setup.Logger.Error("not found",
			zap.String("details", status.Errorf(codes.NotFound, "id not found").Error()),
		)
		return nil, status.Errorf(codes.NotFound, "id not found")
	}

	info := Info{
		Details: "Char deleted",
	}
	go setup.Logger.Info("Character deleted")
	return &info, nil
}

func (s *Server) Update(ctx context.Context, char *CharEdit) (*Character, error) {
	err := char.Data.Update(char.Id.Id)
	if err != nil {
		go setup.Logger.Error("not found",
			zap.String("details", status.Errorf(codes.NotFound, "id not found").Error()),
		)
		return nil, status.Errorf(codes.NotFound, "id not found")
	}
	go setup.Logger.Info("Character: " + char.Data.Name + " updated")
	return char.Data, nil
}

func (s *Server) GetAll(ctx context.Context, db *DB) (*CharList, error) {
	list := CharList{
		Char: FakeDB,
	}

	go setup.Logger.Info("Characters returned")
	return &list, nil
}
