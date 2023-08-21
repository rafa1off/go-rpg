package db

type Database interface {
	Create(value interface{}) error
	Find(dest interface{}, conds ...interface{}) error
}
