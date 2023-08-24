package app

import (
	"go-rpg/proto"

	"gorm.io/gorm"
)

type character struct {
	gorm.Model
	*proto.Character
}
