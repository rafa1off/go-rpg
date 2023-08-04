package character

import (
	"errors"
)

var FakeDB []*Character

func (t *Character) Save() {
	FakeDB = append(FakeDB, t)
}

func (t *Character) Update(id int64) error {
	for _, char := range FakeDB {
		if char.Id == id {
			if char.Name != t.Name && char.Name != "" {
				char.Name = t.Name
			}
			if char.Race != t.Race {
				char.Race = t.Race
			}
			if char.Class != t.Class {
				char.Class = t.Class
			}
			return nil
		}
	}
	return errors.New("not found")
}

func Delete(id int64) error {
	for i, char := range FakeDB {
		if char.Id == id {
			FakeDB = append(FakeDB[:i], FakeDB[i+1:]...)
			return nil
		}
	}
	return errors.New("not found")
}

func Get(id int64) (*Character, error) {
	for _, char := range FakeDB {
		if char.Id == id {
			return char, nil
		}
	}
	return nil, errors.New("not found")
}
