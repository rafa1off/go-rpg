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

func (s *Server) Create(_ context.Context, c *Character) (*Info, error) {
	char := CharData{
		Char: c,
	}
	char.Save()

	info := Info{
		Details: "Character: " + c.Name + " created",
	}

	go setup.Logger.Info("Character: " + c.Name + " created")
	return &info, nil
}

func (s *Server) Get(_ context.Context, data *CharData) (*CharData, error) {
	char, err := Get(data.Id)
	if err != nil {
		go setup.Logger.Error("not found",
			zap.String("details", status.Errorf(codes.NotFound, "id not found").Error()),
		)
		return nil, status.Errorf(codes.NotFound, "id not found")
	}

	res := CharData{
		Id:   data.Id,
		Char: char,
	}

	go setup.Logger.Info("Character: " + char.Name + " returned")
	return &res, nil
}

func (s *Server) Delete(_ context.Context, data *CharData) (*Info, error) {
	if err := Delete(data.Id); err != nil {
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

func (s *Server) Update(_ context.Context, data *CharData) (*Character, error) {
	item := CharData{
		Char: data.Char,
	}

	err := item.Update(data.Id)
	if err != nil {
		go setup.Logger.Error("not found",
			zap.String("details", status.Errorf(codes.NotFound, "id not found").Error()),
		)
		return nil, status.Errorf(codes.NotFound, "id not found")
	}

	go setup.Logger.Info("Character: " + item.Char.Name + " updated")
	return item.Char, nil
}

func (s *Server) GetAll(db *DB, stream Characters_GetAllServer) error {
	for _, i := range FakeDB {
		if err := stream.Send(i); err != nil {
			go setup.Logger.Error("error sending stream",
				zap.String("details", err.Error()),
			)
			return err
		}
	}

	go setup.Logger.Info("Characters returned")
	return nil
}
