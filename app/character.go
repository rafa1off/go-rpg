package app

import (
	"go-rpg/proto"

	"gorm.io/gorm"
)

type charCore struct {
	db *gorm.DB
}

func CharCore(db *gorm.DB) *charCore {
	return &charCore{
		db: db,
	}
}

type character struct {
	gorm.Model
	*proto.Character
}

func CharModel() *character {
	return &character{}
}

func (c *charCore) New(char *proto.Character) (*character, error) {
	entry := character{
		Character: char,
	}
	entry.Character.Health = 100

	err := c.db.Create(&entry).Error
	if err != nil {
		return nil, err
	}

	return &entry, nil
}

func (c *charCore) All() ([]character, error) {
	var chars []character
	err := c.db.Find(&chars).Error
	if err != nil {
		return nil, err
	}
	return chars, nil
}
