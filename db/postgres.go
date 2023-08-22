package db

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Postgres(models ...interface{}) (*gorm.DB, error) {
	dsn := os.Getenv("POSTGRES_CONN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
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
