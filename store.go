package stack

import (
	"log"

	"github.com/elos/data"
	"github.com/elos/models/persistence"
)

func SetupStore(addr string) data.Store {
	db, err := persistence.MongoDB(addr)
	if err != nil {
		log.Fatal(err)
	}

	log.Print("Database connection established")
	s := persistence.Store(db)

	return s
}
