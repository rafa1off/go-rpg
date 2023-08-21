package characters

import (
	"go-rpg/proto"

	"gorm.io/gorm"
)

type Character struct {
	gorm.Model
	*proto.Character
}

func new(char *proto.Character) Character {
	c := Character{
		Character: char,
	}

	c.Character.Health = 100
	return c
}
