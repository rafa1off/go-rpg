package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InMemory(models ...interface{}) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		PrepareStmt:                              true,
	})
	if err != nil {
		return nil, err
	}

	if err = db.AutoMigrate(models...); err != nil {
		return nil, err
	}

	return db, nil
}
