package store

import (
	"github.com/orbitaljin/ocha/internal/store/schema"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DB struct {
	path string
	db   *gorm.DB
}

func NewDB(path string) (*DB, error) {
	gormDB, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &DB{
		path: path,
		db:   gormDB,
	}, nil
}

func (db *DB) Init() error {
	// Perform any database initialization here.
	// For example, you can create tables using AutoMigrate:
	err := db.db.AutoMigrate(&schema.Note{})
	if err != nil {
		return err
	}

	return nil
}

func (db *DB) DB() *gorm.DB {
	return db.db
}

// Add other methods/functions related to your store package.
