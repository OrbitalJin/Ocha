package store

import (
	"github.com/orbitaljin/ocha/internal/store/schema"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// DB is a wrapper around gorm.DB
type DB struct {
	path string
	db   *gorm.DB
}

// Contructor for DB
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

// Initialize the database
func (db *DB) Init() error {
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
