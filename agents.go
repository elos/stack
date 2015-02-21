package stack

import (
	"log"

	"github.com/elos/agents"
	"github.com/elos/data"
	"github.com/elos/models"
	"github.com/elos/models/user"
)

func SetupAgents(s data.Store) {
	iter, err := s.NewQuery(models.UserKind).Execute()

	if err != nil {
		log.Print("shit: err: %s", err.Error())
	}

	u, _ := user.New(s)

	for iter.Next(u) {
		access := data.NewAccess(u, s)
		a := agents.NewActionAgent(access, u)
		go a.Start()
	}

	if err := iter.Close(); err != nil {
		log.Print("oh no, err: %s", err.Error())
	}
}
