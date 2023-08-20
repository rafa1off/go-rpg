package app

import (
	"go-rpg/proto"

	"gorm.io/gorm"
)

type Character struct {
	gorm.Model
	*proto.Character
}

func NewChar(char *proto.Character) Character {
	c := Character{
		Character: char,
	}

	c.Character.Health = 100
	return c
}
