package main

import (
	"log"
	"os"

	"github.com/orbitaljin/ocha/app"
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

    if err := app.New(db).Run(os.Args); err != nil {
      log.Fatal(err)
  }
}
	