package app

import "go-rpg/proto"

type Core interface {
	New(char *proto.Character) (*character, error)
	All(skip, limit int) ([]*character, error)
	Delete(char *character) error
	Get(id int) (*character, error)
	Update(id int, char *proto.Character) (*character, error)
}
