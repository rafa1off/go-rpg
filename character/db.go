package character

import (
	"errors"
)

type database []*char

type char struct {
	id   int32
	char *Character
}

func (d *database) Save(c *Character) *char {
	newChar := char{
		char: c,
	}
	newChar.id = int32(len(*d) + 1)
	*d = append(*d, &newChar)
	return &newChar
}

func (d *database) Update(c *char) (*Character, error) {
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

func (d *database) Delete(id int32) error {
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

func (d *database) Get(id int32) (*Character, error) {
	for _, item := range *d {
		if item.id == id {
			return item.char, nil
		}
	}
	return nil, errors.New("not found")
}

func (d *database) GetAll() []*char {
	return *d
}
