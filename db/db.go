package db

import "go-rpg/app"

type Database interface {
	Save(char *app.Character)
}
