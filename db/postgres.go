package db

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type postgresDB struct {
	engine *gorm.DB
}

func NewPostgres(models ...interface{}) (*postgresDB, error) {
	dsn := os.Getenv("POSTGRES_CONN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(models...)

	return &postgresDB{
		engine: db,
	}, nil
}

func (db *postgresDB) Create(value interface{}) error {
	err := db.engine.Create(value).Error
	if err != nil {
		return err
	}
	return nil
}

func (db *postgresDB) Find(dest interface{}, conds ...interface{}) error {
	err := db.engine.Find(dest).Error
	if err != nil {
		return err
	}
	return nil
}
