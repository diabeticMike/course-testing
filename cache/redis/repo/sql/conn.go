package sql

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const dbURL = "postgres://postgres:secret@localhost:5432/db?sslmode=disable"

func NewConn() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	return db, err
}
