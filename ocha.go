package main

import (
	"github.com/orbitaljin/ocha/internal/store"
)

func main() {
		db, err := store.NewDB("./data/db.db")
		if err != nil {
				panic("Failed to connect to the database")
		}	
		err = db.Init()
		if err != nil {
			panic("Failed to init db")
		}
}
	