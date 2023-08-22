package app

import "go-rpg/proto"

type Core interface {
	New(char *proto.Character) (*character, error)
	All() ([]character, error)
}
