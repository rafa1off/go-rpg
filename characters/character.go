package characters

import (
	"go-rpg/proto"

	"gorm.io/gorm"
)

type character struct {
	gorm.Model
	*proto.Character
}

func Model() *character {
	return &character{}
}

func new(char *proto.Character) character {
	c := character{
		Character: char,
	}

	c.Character.Health = 100
	return c
}
