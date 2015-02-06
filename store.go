package stack

import (
	"log"

	"github.com/elos/data"
	"github.com/elos/data/mongo"
	"github.com/elos/models"
	"github.com/elos/models/event"
	"github.com/elos/models/task"
	"github.com/elos/models/user"
)

func SetupStore(addr string) data.Store {
	mongo.RegisterKind(models.UserKind, "users")
	mongo.RegisterKind(models.EventKind, "events")
	mongo.RegisterKind(models.TaskKind, "tasks")

	db, err := mongo.NewDB(addr)

	if err != nil {
		log.Fatal(err)
	}

	log.Print("Database connection established")

	s := data.NewStore(db, models.Schema)

	s.Register(models.UserKind, user.NewM)
	s.Register(models.EventKind, event.NewM)
	s.Register(models.TaskKind, task.NewM)

	return s
}
