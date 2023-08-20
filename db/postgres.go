package db

import (
	"go-rpg/app"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresDB struct {
	engine *gorm.DB
}

var dsn = os.Getenv("POSTGRES_CONN")

func NewPostgres() (*PostgresDB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&app.Character{})

	return &PostgresDB{
		engine: db,
	}, nil
}

func (db *PostgresDB) Save(char *app.Character) {
	db.engine.Create(char)
}
