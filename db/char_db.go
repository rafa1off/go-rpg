package db

import "go-rpg/app"

type CharDB []*app.Character

func NewCharDB() *CharDB {
	return &CharDB{}
}

func (db *CharDB) Save(char *app.Character) {
	char.ID = uint(len(*db) + 1)
	*db = append(*db, char)
}
