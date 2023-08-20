package db

import "go-rpg/app"

type Database interface {
	Save(char *app.Character)
}

type MemDB []*app.Character

func NewMemDB() *MemDB {
	return &MemDB{}
}

func (db *MemDB) Save(char *app.Character) {
	char.Id = int32(len(*db) + 1)
	*db = append(*db, char)
}
