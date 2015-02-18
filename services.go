package stack

import (
	"github.com/elos/agents"
	"github.com/elos/data"
	"github.com/elos/models"
	"github.com/elos/models/user"
)

var Outfitter *agents.Outfitter

func SetupServices(s data.Store) {
	Outfitter = agents.NewOutfitter()
	go Outfitter.Run()

	iter, err := s.NewQuery(models.UserKind).Execute()
	if err != nil {
	}

	u, _ := user.New(s)

	for iter.Next(u) {
		agents.OutfitUser(Outfitter, s, u)
		access := data.NewAccess(u, s)
		a := agents.NewActionAgent(access, u)
		go a.Start()
	}

	if err := iter.Close(); err != nil {
	}

}
