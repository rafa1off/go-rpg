package character

import (
	"errors"
)

type database interface {
	Save(c *Character) *char
	Update(c *char) (*Character, error)
	Delete(id int32) error
	Get(id int32) (*Character, error)
	GetAll() []*char
}

type memDB []*char

func NewMemDB() *memDB {
	return &memDB{}
}

type char struct {
	id   int32
	char *Character
}

func (d *memDB) Save(c *Character) *char {
	newChar := char{
		char: c,
	}
	newChar.id = int32(len(*d) + 1)
	*d = append(*d, &newChar)
	return &newChar
}

func (d *memDB) Update(c *char) (*Character, error) {
	for _, i := range *d {
		if i.id == c.id {
			if i.char.Name != c.char.Name && i.char.Name != "" {
				i.char.Name = c.char.Name
			}
			if i.char.Race != c.char.Race {
				i.char.Race = c.char.Race
			}
			if i.char.Class != c.char.Class {
				i.char.Class = c.char.Class
			}
			if i.char.Health != c.char.Health {
				i.char.Health = c.char.Health
			}

			return c.char, nil
		}
	}
	return nil, errors.New("not found")
}

func (d *memDB) Delete(id int32) error {
	list := d.GetAll()
	for i, char := range list {
		if char.id == id {
			list = append(list[:i], list[i+1:]...)
			*d = list
			return nil
		}
	}
	return errors.New("not found")
}

func (d *memDB) Get(id int32) (*Character, error) {
	for _, item := range *d {
		if item.id == id {
			return item.char, nil
		}
	}
	return nil, errors.New("not found")
}

func (d *memDB) GetAll() []*char {
	return *d
}
