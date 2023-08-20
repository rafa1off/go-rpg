package app

import "go-rpg/proto"

type Character struct {
	Id int32
	*proto.Character
}

func NewChar(char *proto.Character) Character {
	c := Character{
		Character: char,
	}

	c.Character.Health = 100

	return c
}
