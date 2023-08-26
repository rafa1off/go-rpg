package app

import (
	"go-rpg/proto"

	"gorm.io/gorm"
)

type charCore struct {
	db *gorm.DB
}

func CharCore(db *gorm.DB) *charCore {
	db.AutoMigrate(&character{})

	return &charCore{
		db: db,
	}
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

func (c *charCore) All(limit, skip int) ([]*character, error) {
	var chars []*character

	err := c.db.Limit(limit).Offset(skip).Find(&chars).Error
	if err != nil {
		return nil, err
	}

	return chars, nil
}

func (c *charCore) Delete(char *character) error {
	return c.db.Delete(char).Error
}

func (c *charCore) Get(id int) (*character, error) {
	var entry character

	err := c.db.Find(&entry, id).Error
	if err != nil {
		return nil, err
	}

	return &entry, nil
}

func (c *charCore) Update(id int, char *proto.Character) (*character, error) {
	entry := character{
		Model: gorm.Model{
			ID: uint(id),
		},
		Character: char,
	}
	entry.Character.Health = 100

	err := c.db.Save(&entry).Error
	if err != nil {
		return nil, err
	}

	return &entry, nil
}
