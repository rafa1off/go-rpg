package db

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Postgres() (*gorm.DB, error) {
	dsn := os.Getenv("POSTGRES_CONN")
	return gorm.Open(postgres.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
}
