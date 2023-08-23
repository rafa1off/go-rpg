package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InMemory() (*gorm.DB, error) {
	return gorm.Open(sqlite.Open(":memory:"), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		PrepareStmt:                              true,
	})
}
