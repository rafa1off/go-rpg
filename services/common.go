package services

import "go-rpg/db"

type ServerOpts struct {
	db db.Database
}

type ServerOptsFn func(*ServerOpts)

func defaultOpts() ServerOpts {
	return ServerOpts{
		db: nil,
	}
}

func SetDB(db db.Database) func(*ServerOpts) {
	return func(so *ServerOpts) {
		so.db = db
	}
}
