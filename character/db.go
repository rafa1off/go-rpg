package character

import (
	"errors"
)

var FakeDB []*CharData

func (t *CharData) Save() {
	t.Id = int32(len(FakeDB) + 1)
	FakeDB = append(FakeDB, t)
}

func (t *CharData) Update(id int32) error {
	for _, i := range FakeDB {
		if i.Id == id {
			if i.Char.Name != t.Char.Name && i.Char.Name != "" {
				i.Char.Name = t.Char.Name
			}
			if i.Char.Race != t.Char.Race {
				i.Char.Race = t.Char.Race
			}
			if i.Char.Class != t.Char.Class {
				i.Char.Class = t.Char.Class
			}
			return nil
		}
	}
	return errors.New("not found")
}

func Delete(id int32) error {
	for i, char := range FakeDB {
		if char.Id == id {
			FakeDB = append(FakeDB[:i], FakeDB[i+1:]...)
			return nil
		}
	}
	return errors.New("not found")
}

func Get(id int32) (*Character, error) {
	for _, i := range FakeDB {
		if i.Id == id {
			return i.Char, nil
		}
	}
	return nil, errors.New("not found")
}
