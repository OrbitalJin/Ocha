package main

import (
	"log"
	"os"

	"github.com/orbitaljin/ocha/app"
	"github.com/orbitaljin/ocha/internal/store"
)


func setupDB(path string) *store.DB {
	db, err := store.NewDB(path + "/db.db")
	if err != nil {
			panic("Failed to connect to the database")
	}	
	err = db.Init()
	if err != nil {
		panic("Failed to init db")
	}
	return db
}

func setupCli(db *store.DB) {
	if err := app.New(db).Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func getPath() string {
	path := os.Getenv("HOME") + "/.ocha-cli"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, 0755)
		
		if err != nil {
			panic(err)
		}
	}
	return path
}

func main() {
	path := getPath()
	db := setupDB(path)
	setupCli(db)
}