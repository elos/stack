package stack

import (
	"log"

	"github.com/elos/data"
	"github.com/elos/models"
	"github.com/elos/models/action"
	"github.com/elos/models/event"
	"github.com/elos/models/fixture"
	"github.com/elos/models/routine"
	"github.com/elos/models/schedule"
	"github.com/elos/models/set"
	"github.com/elos/models/task"
	"github.com/elos/models/user"
	"github.com/elos/mongo"
)

func SetupStore(addr string) data.Store {

	db := mongo.NewDB()
	if err := db.Connect(addr); err != nil {
		log.Fatal(err)
	}
	db.SetName("test")
	db.RegisterKind(models.UserKind, "users")
	db.RegisterKind(models.EventKind, "events")
	db.RegisterKind(models.TaskKind, "tasks")
	db.RegisterKind(models.RoutineKind, "routines")
	db.RegisterKind(models.ActionKind, "actions")
	db.RegisterKind(models.SetKind, "sets")
	db.RegisterKind(models.FixtureKind, "fixtures")
	db.RegisterKind(models.ScheduleKind, "schedules")

	log.Print("Database connection established")

	s := data.NewStore(db, models.Schema)

	s.Register(models.UserKind, user.NewM)
	s.Register(models.EventKind, event.NewM)
	s.Register(models.TaskKind, task.NewM)
	s.Register(models.RoutineKind, routine.NewM)
	s.Register(models.ActionKind, action.NewM)
	s.Register(models.SetKind, set.NewM)
	s.Register(models.FixtureKind, fixture.NewM)
	s.Register(models.ScheduleKind, schedule.NewM)

	return s
}
